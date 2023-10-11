package api

import (
	"encoding/json"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/types"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func AllTestCaseModuleHandler(router *httprouter.Router) *httprouter.Router {
	router.POST("/test-case/create", utils.JWTMiddleware(createTestCaseModule))
	router.POST("/test-case/get-all", utils.JWTMiddleware(getAllTesCaseModule))
	router.POST("/test-case/get-one", utils.JWTMiddleware(getOneTestCaseModule))
	router.GET("/test-case/update-one", utils.JWTMiddleware(updateTestCaseModule))
	router.GET("/test-case/delete-one", utils.JWTMiddleware(deleteTestCaseModule))
	return router
}

func createTestCaseModule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var testCaseModule *types.GBTestCaseModule = new(types.GBTestCaseModule)
	utils.DecodeRequestBody(r, &testCaseModule)
	testCaseModule.GOBID = utils.GetUid()
	storage.InsertOne("testCaseModule", testCaseModule)
	resonse, _ := json.Marshal(testCaseModule)
	w.Write(resonse)
}

func getAllTesCaseModule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBTestCaseModule = new(types.GBTestCaseModule)
	testCaseModule := new([]types.GBTestCaseModule)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindAll("testCaseModule", user, testCaseModule)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}

func updateTestCaseModule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var updateDoc *types.GBTestCaseModule = new(types.GBTestCaseModule)
	var filter *types.GBTestCaseModule = new(types.GBTestCaseModule)
	utils.DecodeRequestBody(r, &updateDoc)
	filter.GOBID = utils.GetUid()
	storage.UpdateOne("testCaseModule", filter, updateDoc)
	resonse, _ := json.Marshal(updateDoc)
	w.Write(resonse)
}

func deleteTestCaseModule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var filter *types.GBTestCaseModule = new(types.GBTestCaseModule)
	utils.DecodeRequestBody(r, &filter)
	filter.GOBID = utils.GetUid()
	storage.DeleteOne("testCaseModule", filter)
	resonse, _ := json.Marshal(filter)
	w.Write(resonse)
}

func getOneTestCaseModule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	testCaseModule := new(types.GBTestCaseModule)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindOne("testCaseModule", user, testCaseModule)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}
