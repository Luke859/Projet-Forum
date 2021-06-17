package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

myuuid=""

func UUID() {
	myuuid, err := uuid.NewV4()
	if err == nil{
		fmt.Println("Your UUID is:", myuuid)
	}
}


document.cookie = "test1= test; SameSite=None; Secure; HttpOnly";

const cookieValue1 = document.cookie
	.split('; ')
	.find(row => row.startsWith('test1='))
	.split('=')[1];

func showCookieValue() {
	const output = document.getElementById('cookie-value')
	output.textContent = '> ' + cookieValue1
	return output
}
