package main

import (
	"testing"

	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/login/proto"
)

func InitTest() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient("./testing")
	s.GoServer.KSclient.Save(context.Background(), CONFIG, &pb.Config{})
	return s
}

func TestRandString(t *testing.T) {
	str := randString(10)

	if len(str) != 10 {
		t.Errorf("Bad string: %v (%v)", str, len(str))
	}
}

func TestGetToken(t *testing.T) {
	s := InitTestServer()
	token1, err := s.getToken(context.Background(), "brotherlogic@gmail.com")
	if err != nil {
		t.Errorf("Bad get: %v", err)
	}

	token2, err := s.getToken(context.Background(), "brotherlogic@gmail.com")
	if err != nil {
		t.Errorf("Bad get: %v", err)
	}

	if token1 != token2 {
		t.Errorf("%v should match %v", token1, token2)
	}
}

func TestBadGetToken(t *testing.T) {
	s := InitTestServer()
	token, err := s.getToken(context.Background(), "bruderlogic@gmail.com")
	if err == nil {
		t.Errorf("Got token for unauth user: %v", token)
	}
}

func TestBadReadToken(t *testing.T) {
	s := InitTestServer()
	s.GoServer.KSclient.Fail = true
	token, err := s.getToken(context.Background(), "bruderlogic@gmail.com")
	if err == nil {
		t.Errorf("Got token for unauth user: %v", token)
	}
}
