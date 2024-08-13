package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/eslami200117/clientCli/app/models"
	"github.com/eslami200117/clientCli/app/repository"
	"github.com/eslami200117/clientCli/config"
	"github.com/eslami200117/clientCli/database"
)

type auth struct {
	Token string
}

const requestURL = "http://localhost:5000"
const receiveTime = time.Second

func LoginHandler(username string, password string) {

	loginData := map[string]string{"username": username, "password": password}
	jsonData, err := json.Marshal(loginData)
	bodyReader := bytes.NewReader(jsonData)
	if err != nil {
		fmt.Printf("Failed to marshal login data: %v\n", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, requestURL+"/loginuser", bodyReader)
	if err != nil {
		fmt.Println("Failed to create request", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return
	}

	var token auth
	err = json.Unmarshal(resBody, &token)
	if err != nil {
		fmt.Println("error in unmarshal", err)
		return
	}
	if token.Token == "" {
		fmt.Println("error: Unauthorized")
		return
	}
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepo(db)
	repo.InsertAuth(username, token.Token)

}

func ListHandler(username string) {
	jsonData, err := json.Marshal(username)
	bodyReader := bytes.NewReader(jsonData)
	if err != nil {
		fmt.Println("Failed to create marshal username", err)
		return
	}
	token := getTokenByUser(username)
	req, err := http.NewRequest(http.MethodGet, requestURL+"/test/list", bodyReader)
	if err != nil {
		fmt.Println("Failed to create request", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("client: error making http request:", err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("client: could not read response body:", err)
		return
	}
	value := struct {
		Nodes map[string]bool
	}{}

	err = json.Unmarshal(resBody, &value)
	if err != nil {
		fmt.Println("line 103 error in unmarshal", err)
	}
	fmt.Println(value)

}

func NodeHandler(username string, nodename string) {
	jsonData, err := json.Marshal(username)
	bodyReader := bytes.NewReader(jsonData)
	if err != nil {
		fmt.Println("Failed to create marshal username", err)
		return
	}
	token := getTokenByUser(username)
	req, err := http.NewRequest(http.MethodGet, requestURL+"/test/node?node="+nodename, bodyReader)
	if err != nil {
		fmt.Println("Failed to create request", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	for {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("client: error making http request:", err)
			return
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("client: could not read response body:", err)
			return
		}
		value := struct {
			Data model.Weather
		}{}

		err = json.Unmarshal(resBody, &value)
		if err != nil {
			fmt.Println("line 103 error in unmarshal", err)
		}
		fmt.Println(value)
		time.Sleep(receiveTime)
	}
}

func getTokenByUser(username string) string {
	config := config.GetConfig()
	db := database.NewPostgresDatabase(config)
	repo := repository.NewRepo(db)
	var token string = repo.GetToken(username)
	return token

}
