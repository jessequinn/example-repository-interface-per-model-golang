package main

import (
	"fmt"
	"net/http"

	"example-repository-interface-per-model-golang/internal/controllers"
	"example-repository-interface-per-model-golang/internal/repositories"
	"example-repository-interface-per-model-golang/internal/sqldb"
)

func main() {
	db, err := sqldb.ConnectDB("mysql", "username:password@/dbname")
	if err != nil {
		panic(err.Error())
	}

	// Create repos
	userRepo := repositories.NewUserRepo(db)

	h := controllers.NewBaseHandler(userRepo)

	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())

	}
}
