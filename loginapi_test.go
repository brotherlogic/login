package main

import (
	"context"
	"testing"

	"github.com/brotherlogic/keystore/client"

	pb "github.com/brotherlogic/login/proto"
)

//InitTestServer gets a test version of the server
func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient(".test")
	return s
}

func TestAllFail(t *testing.T) {
	s := InitTestServer()
	_, err := s.Login(context.Background(), &pb.LoginRequest{})
	if err == nil {
		t.Errorf("Oh oh")
	}
	_, err = s.Authenticate(context.Background(), &pb.AuthenticateRequest{})
	if err == nil {
		t.Errorf("Oh oh")
	}

}
