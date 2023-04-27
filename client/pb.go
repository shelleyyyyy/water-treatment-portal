package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pocket struct {
	Username       string
	Password       string
	Pocketbase     string
	PocketbasePort string
	Token          string
}

type authPost struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type authRes struct {
	Record struct {
		Avatar       string `json:"avatar"`
		CollectionID string `json:"collectionId"`
		Created      string `json:"created"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		Username     string `json:"username"`
	}
	Token string `json:"token"`
}

func (p Pocket) Testing() {
	fmt.Println("hello")
}

func (p Pocket) Auth() error {
	posturl := fmt.Sprintf("http://%s:%s/api/collections/users/auth-with-password", p.Pocketbase, p.PocketbasePort)

	// Create a new HTTP client
	data := authPost{
		Identity: p.Username,
		Password: p.Password,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request with a POST method and the API endpoint
	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Print the response body

	var ares authRes

	err = json.Unmarshal(body, &ares)
	if err != nil {
		fmt.Println(err)
		return err
	}

	p.Token = ares.Token

	return nil

}

func (p Pocket) GetList(collection string) ([]byte, error) {
	api := fmt.Sprintf("http://%s:%s/api/collections/%s/records", p.Pocketbase, p.PocketbasePort, collection)

	req, err := http.NewRequest("GET", api, nil)

	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return nil, err
	}

	// Set the Authorization header to include the token
	req.Header.Set("Authorization", p.Token)

	// Create a new HTTP client and send the request
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return nil, err
	}

	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return nil, err
	}

	// Print the response body
	return body, nil
}

func (p Pocket) Update(recordID string, collection string, payload []byte) error {

	httpposturl := fmt.Sprintf("http://%s:%s/api/collections/%s/records/%s", p.Pocketbase, p.PocketbasePort, collection, recordID)

	request, err := http.NewRequest("PATCH", httpposturl, bytes.NewBuffer(payload))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)

	return nil
}

func (p Pocket) Create(collection string, payload []byte) error {

	posturl := fmt.Sprintf("http://%s:%s/api/collections/%s/records", p.Pocketbase, p.PocketbasePort, collection)
	// Create a new HTTP client
	fmt.Println(posturl)
	client := &http.Client{}

	// Create a new HTTP request with a POST method and the API endpoint
	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Set the Bearer token in the Authorization header
	req.Header.Set("Content-Type", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	return nil
}
