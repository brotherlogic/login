package main

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/login/proto"
)

const (
	// USERS - store of all users
	USERS = "/github.com/brotherlogic/login/users"
)

//Login logs us un
func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.Log(fmt.Sprintf("Processing  login request: %v", s.verifyToken(ctx, req.GetToken())))
	return nil, fmt.Errorf("Not implemented yet")
}

//Authenticate attempts to authenticate
func (s *Server) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	return nil, fmt.Errorf("Not implemented yet")
}
