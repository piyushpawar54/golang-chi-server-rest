package main

import (
	"context"
	"fmt"

	"github.com/piyushpawar54/golang-chi-server-rest/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
