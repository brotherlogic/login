package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/login/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&utils.DiscoveryClientResolverBuilder{})
}

func main() {
	conn, err := grpc.Dial("discovery:///login", grpc.WithInsecure(), grpc.WithBalancerName("my_pick_first"))
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoginServiceClient(conn)
	ctx, cancel := utils.BuildContext("login-cli", "login")
	defer cancel()

	switch os.Args[1] {
	case "set":
		content, err := ioutil.ReadFile(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		res, err := client.SetToken(ctx, &pb.SetTokenRequest{Token: string(content)})
		if err != nil {
			log.Fatalf("Error in listing: %v", err)
		}
		fmt.Printf("Result = %v\n", res)
	}

}
