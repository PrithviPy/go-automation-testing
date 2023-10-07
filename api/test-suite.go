package api

import (
	"encoding/json"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/types"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func AllTestSuiteHandler(router *httprouter.Router) *httprouter.Router {
	router.POST("/test-suite/create", utils.JWTMiddleware(createTestSuite))
	router.POST("/test-suite/get-all", utils.JWTMiddleware(getAllTesSuite))
	router.POST("/test-suite/get-one", utils.JWTMiddleware(getOneTestSuite))
	router.GET("/test-suite/update-one", utils.JWTMiddleware(updateTestSuite))
	router.GET("/test-suite/delete-one", utils.JWTMiddleware(deleteTestSuite))
	return router
}

func createTestSuite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var testSuite *types.GBTestSuite = new(types.GBTestSuite)
	utils.DecodeRequestBody(r, &testSuite)
	testSuite.GOBID = utils.GetUid()
	storage.InsertOne("testsuite", testSuite)
	resonse, _ := json.Marshal(testSuite)
	w.Write(resonse)
}

func getAllTesSuite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBTestSuite = new(types.GBTestSuite)
	testSuite := new([]types.GBTestSuite)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindAll("testsuite", user, testSuite)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}

func updateTestSuite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var updateDoc *types.GBTestSuite = new(types.GBTestSuite)
	var filter *types.GBTestSuite = new(types.GBTestSuite)
	utils.DecodeRequestBody(r, &updateDoc)
	filter.GOBID = utils.GetUid()
	storage.UpdateOne("testsuite", filter, updateDoc)
	resonse, _ := json.Marshal(updateDoc)
	w.Write(resonse)
}

func deleteTestSuite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var filter *types.GBTestSuite = new(types.GBTestSuite)
	utils.DecodeRequestBody(r, &filter)
	filter.GOBID = utils.GetUid()
	storage.DeleteOne("testsuite", filter)
	resonse, _ := json.Marshal(filter)
	w.Write(resonse)
}

func getOneTestSuite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	testSuite := new(types.GBTestSuite)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindOne("testsuite", user, testSuite)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}
