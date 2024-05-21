package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserID int
	PostID int
	Title  string
	Body   string
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/")

	if err != nil {
		fmt.Print("Error while fetching")
		return
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Print("Error while reading data from resp body")
	}

	var resArr []Post
	err = json.Unmarshal(resData, &resArr)

	if err != nil {
		fmt.Print("Error while unmarshalling")
	}

	fmt.Println("==============")
	for _, item := range resArr {
		fmt.Print(item.PostID, item.UserID, "\n")
		fmt.Println("-----------")
		fmt.Println(item.Title)
		fmt.Println("-----------")
		fmt.Println(item.Body)
		fmt.Println("==============")
	}
}
