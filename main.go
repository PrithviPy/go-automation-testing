package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/api"
	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func main() {
	dbUri, port := utils.GetEnvForStarup()
	logFile := utils.LoggerStartup()
	cancel, err := storage.Connect(dbUri)
	defer storage.Close(cancel)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	storage.Ping()
	fmt.Print("Application Started complete log available in application.log file !")
	router := httprouter.New()
	api.AllUserGroupHandlers(router)
	api.AllWorkspcaeHandlers(router)
	api.AllTestSuiteHandler(router)
	log.Fatal(http.ListenAndServe(port, router))
}
