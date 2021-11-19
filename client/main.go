package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Country string `json:"country"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/api/v1/get-data")

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Error, status code:", resp.StatusCode)

		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	var response []Person

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(response)
}