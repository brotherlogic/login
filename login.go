package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbg "github.com/brotherlogic/goserver/proto"
	pb "github.com/brotherlogic/login/proto"
)

const (
	//CONFIG where we store the config
	CONFIG = "github.com/brotherlogic/login/config"

	//USERS where we store users
	USERS = "github.com/brotherlogic/login/users"
)

// Server main server type
type Server struct {
	*goserver.GoServer
}

// Init builds the server
func Init() *Server {
	s := &Server{
		GoServer: &goserver.GoServer{},
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterLoginServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{
		&pbg.State{Key: "no", Value: int64(233)},
	}
}

func (s *Server) verifyFirebaseToken(ctx context.Context, tokenStr string) (string, error) {
	return "", nil
}

func (s *Server) verifyToken(ctx context.Context, token string, firebaseToken string) (string, error) {
	return "", nil
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	var init = flag.Bool("init", false, "Prep server")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.PrepServer("login")
	server.Register = server

	err := server.RegisterServerV2(false)
	if err != nil {
		return
	}

	if *init {
		log.Fatalf("INIT")
		//		ctx, cancel := utils.BuildContext("login", "login")
		//		defer cancel()
		//
		//		err := server.KSclient.Save(ctx, QUEUE, &pb.Queue{ProcessedRecords: 1})
		//		fmt.Printf("Initialised: %v\n", err)
		//		return
	}

	server.Serve()
}
