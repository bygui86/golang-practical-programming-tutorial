package main

import (
	"fmt"
)

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a *author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p *post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Bio: ", p.bio)
	fmt.Println("First name: ", p.author.firstName)
	fmt.Println("Last name: ", p.author.lastName)
	fmt.Println("Author full name: ", p.fullName())
}

func main() {
	author := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	post := post{
		"Composition (a.k.a. Inheritance) in Go",
		"Go supports composition instead of inheritance :(",
		author,
	}

	post.details()
}
