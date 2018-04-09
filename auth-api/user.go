package main

import (
	"context"
	"fmt"
	"net/http"
)

type (
	User struct {
		Username  string `json:"username"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	HTTPDoer interface {
		Do(req *http.Request) (*http.Response, error)
	}

	UserService struct {
		Client HTTPDoer
	}
)

func (h *UserService) Login(ctx context.Context, username string, password string) (User, error) {
	user := User{}
	return user, nil
}

func (h *UserService) getUser(ctx context.Context, username string) (User, error) {
	var user User

	token, err := h.getUserAPIToken(username)
	if err != nil {
		return user, err
	}

	url := fmt.Sprintf("%s/users/%s", h.getUserAPIToken, username)
}

func (h *UserService) getUserAPIToken(username string) (string, error) {
	return "", nil
}
