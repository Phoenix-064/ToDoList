package main

import (
	"ToDoList/internal/email"
	"fmt"
)

func main() {
	em := email.NewEmailManager()
	err := em.ConfigureEmail("3440480965@qq.com")
	fmt.Println(err)
}
