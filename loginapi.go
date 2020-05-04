package main

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/login/proto"
)

//Login logs us un
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.Log(fmt.Sprintf("Processing  login request: %v", s.verifyToken(ctx, req.GetToken(), req.GetFirebaseToken())))
	return nil, fmt.Errorf("Not implemented yet")
}

//Authenticate attempts to authenticate
func (s *Server) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

//SetToken sets the firebase auth token
func (s *Server) SetToken(ctx context.Context, req *pb.SetTokenRequest) (*pb.SetTokenResponse, error) {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if err != nil {
		return nil, err
	}
	config := data.(*pb.Config)

	config.AuthToken = req.GetToken()

	return nil, s.KSclient.Save(ctx, CONFIG, config)
}
