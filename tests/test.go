package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uid, err := uuid.NewUUID()
	if err!= nil{
		panic(err)
	}
	fmt.Println(uid)
}