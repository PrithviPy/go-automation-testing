package api

import (
	"encoding/json"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/types"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func AllTestCaseHandler(router *httprouter.Router) *httprouter.Router {
	router.POST("/test-case/create", utils.JWTMiddleware(createTestcCase))
	router.POST("/test-case/get-all", utils.JWTMiddleware(getAllTesCase))
	router.POST("/test-case/get-one", utils.JWTMiddleware(getOneTestCase))
	router.GET("/test-case/update-one", utils.JWTMiddleware(updateTestCase))
	router.GET("/test-case/delete-one", utils.JWTMiddleware(deleteTestCase))
	return router
}

func createTestcCase(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var testcase *types.GBTestCase = new(types.GBTestCase)
	utils.DecodeRequestBody(r, &testcase)
	testcase.GOBID = utils.GetUid()
	storage.InsertOne("testcase", testcase)
	resonse, _ := json.Marshal(testcase)
	w.Write(resonse)
}

func getAllTesCase(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBTestCase = new(types.GBTestCase)
	testcase := new([]types.GBTestCase)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindAll("testcase", user, testcase)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}

func updateTestCase(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var updateDoc *types.GBTestCase = new(types.GBTestCase)
	var filter *types.GBTestCase = new(types.GBTestCase)
	utils.DecodeRequestBody(r, &updateDoc)
	filter.GOBID = utils.GetUid()
	storage.UpdateOne("testcase", filter, updateDoc)
	resonse, _ := json.Marshal(updateDoc)
	w.Write(resonse)
}

func deleteTestCase(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var filter *types.GBTestCase = new(types.GBTestCase)
	utils.DecodeRequestBody(r, &filter)
	filter.GOBID = utils.GetUid()
	storage.DeleteOne("testcase", filter)
	resonse, _ := json.Marshal(filter)
	w.Write(resonse)
}

func getOneTestCase(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	testcase := new(types.GBTestCase)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindOne("testcase", user, testcase)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}
