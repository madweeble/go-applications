package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	ID int `json:"id"`
	Username string `json:"username"`
	About string `json:"about"`
	UpdatedAt string `json:"updated_at"`
	SubmissionCount int `json:"submission_count"`
	CommentCount int `json:"comment_count"`
	CreatedAt int32 `json:"created_at"`
}

type Response struct {
	Page string `json:"page"`
	PerPage int `json:"per_page"`
	Total int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data []Data `json:"data"`
}

func main() {
	url := "https://jsonmock.hackerrank.com/api/article_users?page=1"
	client := http.Client{
		Timeout: time.Second * 4,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.TotalPages)
}