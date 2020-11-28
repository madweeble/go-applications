package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

const (
	url = "https://jsonmock.hackerrank.com/api/article_users?page="
	threshold = 10
)

func main() {
	fmt.Println("Hello, world")

	authorHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello, handler")
		doStuff()
	}

	http.HandleFunc("/authors", authorHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func doStuff() {
	numPages := getPageCount()

	for pageNum := 1; pageNum <= numPages; pageNum++ {
		response := getPage(strconv.Itoa(pageNum))
		data := response.Data
		for i := 0; i < len(data); i++ {
			user := data[i]
			if user.SubmissionCount > threshold {
				fmt.Println(user.Username)
			}
		}
	}
}

func getPageCount() int {
	response := getPage("1")
	return response.TotalPages
}

func getPage(pageNum string) Response {
	client := http.Client{
		Timeout: time.Second * 4,
	}

	req, err := http.NewRequest(http.MethodGet, url + pageNum, nil)
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

	return response
}