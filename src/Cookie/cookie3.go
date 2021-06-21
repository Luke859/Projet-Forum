package main

import (
	uuid"github.com/google/uuid"
	"fmt"
)

func NewUUID() string {
	myuuid := uuid.New()
	test := myuuid.String()
	fmt.Println(test)
	return ""
}

func main() {
	test := 1
	fmt.Println(test)
	return
}
