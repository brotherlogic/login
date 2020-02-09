package main

import (
	"testing"

	"github.com/brotherlogic/keystore/client"
)

//InitTestServer gets a test version of the server
func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient(".test")
	return s
}

func TestNothing(t *testing.T) {
	doNothing()
}
