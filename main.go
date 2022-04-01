package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

//エラーメッセージの作成
var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}

func main() {
	employee, err := getInformation(1000)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Print(employee)
	}
}
