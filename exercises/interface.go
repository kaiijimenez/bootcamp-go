package main

import (
	"fmt"
)

//interface whith one method notify
type notifier interface {
	notify()
}

//structure for user
type user struct {
	name  string
	email string
}

//function notify for users structure
func (u *user) notify() {
	fmt.Printf("Sending user email to: %s<%s>\n", u.name, u.email)
}

//structure for admin
type admin struct {
	name  string
	email string
}

//notify method for admin
func (a *admin) notify() {
	fmt.Printf("Sending adming email to: %s<%s>\n", a.name, a.email)
}

func main() {
	u := user{"Bill", "bill@example.com"}
	sendNotification(&u)
	a := admin{"admin", "admin@example.com"}
	sendNotification(&a)
}

func sendNotification(n notifier) {
	n.notify()
}
