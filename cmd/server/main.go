// package server runs the main golang api server.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alextanhongpin/go-ddd/pkg/usecase/usercrud"
)

func main() {
	// Init usecase.
	uc := usercrud.New(nil)
	u, err := uc.Service.FindOne(context.Background(), "1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("running server", u)
}
