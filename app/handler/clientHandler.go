package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/eslami200117/clientCli/app/repository"
	"github.com/eslami200117/clientCli/config"
	"github.com/eslami200117/clientCli/database"
)

type auth struct {
	Token string
}

func LoginHandler(username string, password string) {

	requestURL := "http://localhost:5000/loginuser"

	loginData := map[string]string{"username": username, "password": password}
	jsonData, err := json.Marshal(loginData)
	bodyReader := bytes.NewReader(jsonData)
	if err != nil {
		fmt.Printf("Failed to marshal login data: %v\n", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Println("Failed to create request", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var token auth

	err = json.Unmarshal(resBody, &token)
	if err != nil {
		fmt.Println("error in unmarshal", err)
	}
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepo(db)

	repo.InsertAuth(username, token.Token)

}
