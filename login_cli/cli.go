package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/brotherlogic/goserver/utils"

	pb "github.com/brotherlogic/login/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	ctx, cancel := utils.BuildContext("login-cli", "login")
	defer cancel()

	conn, err := utils.LFDialServer(ctx, "login")
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoginServiceClient(conn)

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
