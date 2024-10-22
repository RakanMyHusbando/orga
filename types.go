package main

import "math/rand"

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	return &User{
		Id:   rand.Intn(1000),
		Name: name,
	}
}
