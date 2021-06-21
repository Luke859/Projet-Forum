package main

import (
	"github.com/google/uuid"
	"fmt"
)

func New() uuid.UUID {
	test := uuid.New()
	return test
}

func main(){
	New()
	fmt.Println(test)
}