package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/brotherlogic/login/proto"
)

//Login logs us un
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.verifyToken(ctx, req.GetToken(), req.GetFirebaseToken())
	return &pb.LoginResponse{Token: token}, err
}

//Authenticate attempts to authenticate
func (s *Server) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if status.Convert(err).Code() != codes.NotFound && status.Convert(err).Code() != codes.OK {
		return nil, err
	}

	config := &pb.Config{}
	if data != nil {
		config = data.(*pb.Config)
	}

	for _, user := range config.GetUsers() {
		if user.GetToken() == req.GetToken() {
			return &pb.AuthenticateResponse{}, nil
		}
	}

	return nil, fmt.Errorf("Not able to authenticate this token (%v)", req.GetToken())
}

//SetToken sets the firebase auth token
func (s *Server) SetToken(ctx context.Context, req *pb.SetTokenRequest) (*pb.SetTokenResponse, error) {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if status.Convert(err).Code() != codes.NotFound && status.Convert(err).Code() != codes.OK {
		return nil, err
	}

	config := &pb.Config{}
	if data != nil {
		config = data.(*pb.Config)
	}

	config.AuthToken = req.GetToken()

	return &pb.SetTokenResponse{}, s.KSclient.Save(ctx, CONFIG, config)
}
