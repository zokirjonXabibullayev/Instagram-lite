package main

import (
	"fmt"
	"instagram/handlers"
	"net/http"
)

func main() {
	fmt.Println("Server ishlayapti ... :8080")

	http.HandleFunc("/user", handlers.UserHendler)
	http.HandleFunc("/post", handlers.PostHendler)
	http.HandleFunc("/comment", handlers.CommentHendler)

	http.ListenAndServe(":8080", nil)
}

