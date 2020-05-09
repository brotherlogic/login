package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	oauth2 "google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
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

//Server main server type
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
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if err != nil {
		s.Log(fmt.Sprintf("Err %v", err))
		return "", err
	}
	conf := data.(*pb.Config)

	opt := option.WithCredentialsJSON([]byte(conf.GetAuthToken()))
	config := &firebase.Config{ProjectID: "androidgetter-d1c08"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		s.Log(fmt.Sprintf("ERR : %v", err))
		return "", err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		s.Log(fmt.Sprintf("ERR : %v", err))
		return "", err
	}

	token, err := client.VerifyIDToken(ctx, tokenStr)
	if err != nil {
		s.Log(fmt.Sprintf("ERR : %v", err))
		return "", err
	}

	tok, err := s.getToken(ctx, token.Claims["email"].(string))
	if err != nil {
		s.Log(fmt.Sprintf("BAD TOKEN: %v", err))
		return "", err
	}

	return tok, nil
}

func (s *Server) verifyToken(ctx context.Context, token string, firebaseToken string) (string, error) {
	if len(firebaseToken) > 0 {
		return s.verifyFirebaseToken(ctx, firebaseToken)
	}

	svc, err := oauth2.New(http.DefaultClient)
	ti, err := svc.Tokeninfo().IdToken(token).Context(ctx).Do()
	if err != nil {
		return fmt.Sprintf("Err for token %v: %v", token, err), err
	}
	if ti.VerifiedEmail {
		return ti.Email, nil
	}
	return fmt.Sprintf("%+v", ti), nil
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
	server.PrepServer()
	server.Register = server

	err := server.RegisterServerV2("login", false, true)
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

	fmt.Printf("%v", server.Serve())
}
