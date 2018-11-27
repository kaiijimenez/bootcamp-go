package main

import "fmt"

type error interface {
	Error() string
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func internalError(text string) error {
	return &errorString{text}
}

func thirdError(text string) error {
	return &errorString{text}
}

func otherError(text string) error {
	return &errorString{text}
}

func defaultError(text string) error {
	return &errorString{text}
}

func main() {
	user1 := checkErrors("sintaxis")
	user2 := checkErrors("package")
	user3 := checkErrors("other")

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
}

//sintaxis error, package error, other error
func checkErrors(e string) error {
	switch e {
	case "sintaxis":
		return internalError("It occurred and internal error related with sintaxis!")
	case "package":
		return thirdError("A third error package occurred!")
	case "other":
		return otherError("Other error occurred!")
	default:
		return defaultError("Someother error occurred!")
	}
}
