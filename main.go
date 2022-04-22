package main

import (
	"fmt"
	"log"
	"net/http"
	"social-app/src/router"
	"social-app/src/utils"
)

func init() {
	utils.LoadTemplates()
}

func main() {
	appRouter := router.Generate()

	Port := 3000

	fmt.Printf("App is running at port: %d\n", Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), appRouter))
}
