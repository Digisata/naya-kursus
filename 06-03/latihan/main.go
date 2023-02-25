package main

import "fmt"

func actionLike(ch chan string) {
	for v := range ch {
		if v != "like" {
			ch <- v
		} else {
			fmt.Printf("like = %s\n", v)
		}
	}
}

func actionComment(ch chan string) {
	for v := range ch {
		if v != "comment" {
			ch <- v
		} else {
			fmt.Printf("comment = %s\n", v)
		}
	}
}

func main() {
	action := make(chan string)
	types := []string{"like", "comment", "like", "comment", "like", "comment", "like", "comment", "like", "like", "like"}

	go actionLike(action)
	go actionComment(action)

	for _, v := range types {
		action <- v
	}
}
