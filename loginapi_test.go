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
	s.GoServer.KSclient.Save(context.Background(), CONFIG, &pb.Config{})
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

func TestSetToken(t *testing.T) {
	s := InitTestServer()

	_, err := s.SetToken(context.Background(), &pb.SetTokenRequest{Token: "aha"})
	if err != nil {
		t.Fatalf("Add Record failed: %v", err)
	}
}

func TestAddRequestFail(t *testing.T) {
	s := InitTestServer()
	s.GoServer.KSclient.Fail = true

	val, err := s.SetToken(context.Background(), &pb.SetTokenRequest{Token: "aha"})
	if err == nil {
		t.Errorf("Set Token passed: %v", val)
	}
}

func TestAuthenticate(t *testing.T) {
	s := InitTestServer()
	tok, err := s.getToken(context.Background(), "brotherlogic@gmail.com")

	_, err = s.Authenticate(context.Background(), &pb.AuthenticateRequest{Token: tok})
	if err != nil {
		t.Fatalf("Could not authenticate: %v", err)
	}
}

func TestAuthenticateBad(t *testing.T) {
	s := InitTestServer()

	_, err := s.Authenticate(context.Background(), &pb.AuthenticateRequest{Token: "aha"})
	if err == nil {
		t.Fatalf("Authenticated with bad token")
	}
}

func TestAuthenticateBadRead(t *testing.T) {
	s := InitTestServer()
	s.GoServer.KSclient.Fail = true

	_, err := s.Authenticate(context.Background(), &pb.AuthenticateRequest{Token: "aha"})
	if err == nil {
		t.Errorf("Authenticate Passed despite bad read")
	}
}
