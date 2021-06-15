package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
  myuuid, err := uuid.NewV4()
  if err == nil{
	fmt.Println("Your UUID is:", myuuid)
  }
}
