package main

import (
	"log"

	"github.com/gillhoang/hoc_api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
