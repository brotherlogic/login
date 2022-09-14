package main

import (
	"fmt"
	"math/rand"

	pb "github.com/brotherlogic/login/proto"
	"golang.org/x/net/context"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *Server) getToken(ctx context.Context, email string) (string, error) {
	data, _, err := s.KSclient.Read(ctx, CONFIG, &pb.Config{})
	if err != nil {
		s.CtxLog(ctx, fmt.Sprintf("Err %v", err))
		return "", err
	}
	conf := data.(*pb.Config)

	for _, user := range conf.GetUsers() {
		if user.GetEmail() == email {
			return user.GetToken(), nil
		}
	}

	if email != "brotherlogic@gmail.com" {
		return "", fmt.Errorf("Not authorized")
	}

	token := randString(20)
	user := &pb.User{Email: email, Token: token}
	conf.Users = append(conf.Users, user)
	return token, s.KSclient.Save(ctx, CONFIG, conf)
}
