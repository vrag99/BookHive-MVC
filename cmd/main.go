package main

import (
	"BookHive/pkg/api"
	"BookHive/pkg/models"
	"fmt"
)

func main() {
	db, err := models.Connection()
	if err!= nil {
        fmt.Println(err)
    }
	defer db.Close()

    api.Run()
}