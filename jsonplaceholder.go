package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	placeholderAPITpl = "https://jsonplaceholder.typicode.com/%s"
)

type user struct {
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func useAPI() {
	users, err := getUsers()
	if err != nil {
		log.Printf("cannot get users: %v", err)
		return
	}

	if len(users) == 0 {
		log.Printf("zero users received")
		return
	}

	log.Printf("First users: ID=%d, Name=%s, Username=%s, Email=%s",
		users[0].ID, users[0].Name, users[0].Username, users[0].Email)

	userToCreate := user{
		Name:     "Our User",
		Username: "ouruser",
		Email:    "test@example.com",
	}
	newUser, err := createUser(userToCreate)
	if err != nil {
		log.Printf("cannot create user: %v", err)
		return
	}

	log.Printf("New user created: ID=%d, Name=%s, Username=%s, Email=%s",
		newUser.ID, newUser.Name, newUser.Username, newUser.Email)
}

func getUsers() ([]user, error) {
	resp, err := http.Get(fmt.Sprintf(placeholderAPITpl, "users"))
	if err != nil {
		return nil, fmt.Errorf("cannot perform users http req: %w", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("could not close users response body: %v", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read users response body: %w", err)
	}

	var users []user
	if err = json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("cannot parse users: %w", err)
	}

	return users, nil
}

func createUser(userToCreate user) (user, error) {
	if userToCreate.ID != 0 {
		return user{}, fmt.Errorf("invalid user: id cannot be non-zero for creation")
	}
	if len(userToCreate.Name) == 0 || len(userToCreate.Username) == 0 || len(userToCreate.Email) == 0 {
		return user{}, fmt.Errorf("invalid user: some mandatory fields are empty")
	}

	userBytes, err := json.Marshal(userToCreate)
	if err != nil {
		return user{}, fmt.Errorf("cannot marshal user: %w", err)
	}

	buf := bytes.NewBuffer(userBytes)
	resp, err := http.Post(fmt.Sprintf(placeholderAPITpl, "users"),
		"application/json", buf)
	if err != nil {
		return user{}, fmt.Errorf("cannot send http post request to create user: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return user{}, fmt.Errorf("cannot read users response body: %w", err)
	}

	var newUser user
	if err = json.Unmarshal(body, &newUser); err != nil {
		return user{}, fmt.Errorf("cannot marshal created user: %w", err)
	}

	return newUser, nil
}