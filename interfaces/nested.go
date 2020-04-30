package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Notifier UserNotifier
}

func (user *User) notify(message string) error {
	return user.Notifier.SendMessage(user, message)
}

type UserNotifier interface {
	SendMessage(user *User, message string) error
}

type EmailNotifier struct{}

func (notifier EmailNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending email to %s with content %s\n", user.Name, message)
	return err
}

type SmsNotifier struct{}

func (notifier SmsNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending SMS to %s with content %s\n", user.Name, message)
	return err
}

func main() {
	user1 := User{"Dirk", "dirk@email.com", EmailNotifier{}}
	user2 := User{"Justin", "bieber@email.com", SmsNotifier{}}

	user1.notify("Welcome Email user!")
	user2.notify("Welcome SMS user!")
}
