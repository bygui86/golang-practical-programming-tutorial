package main

import (
	"fmt"
)

type authorInterface interface {
	getFirstName() string
	getLastName() string
	getBio() string
	fullName() string
}

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a *author) getFirstName() string {
	return a.firstName
}

func (a *author) getLastName() string {
	return a.lastName
}

func (a *author) getBio() string {
	return a.bio
}

func (a *author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	authorInterface
}

func (p *post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Bio: ", p.getBio())
	fmt.Println("First name: ", p.authorInterface.getFirstName())
	fmt.Println("Last name: ", p.authorInterface.getLastName())
	fmt.Println("Author full name: ", p.fullName())
}

func main() {
	var authorInt authorInterface
	authorInt = &author{
		"John",
		"Doe",
		"Golang learner",
	}

	post1 := post{
		"Composition (a.k.a. Inheritance) in Go",
		"Go supports composition instead of inheritance :(",
		authorInt,
	}
	post1.details()
}
