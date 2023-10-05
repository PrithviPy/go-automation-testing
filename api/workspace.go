package api

import (
	"encoding/json"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/types"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func AllWorkspcaeHandlers(router *httprouter.Router) *httprouter.Router {
	router.POST("/workspace/create", utils.JWTMiddleware(CreateWorkspace))
	router.POST("/workspace/get-all-for-user", utils.JWTMiddleware(GetAllWorkspaceForUser))
	router.POST("/workspace/delete", utils.JWTMiddleware(DeleteWorkspace))
	router.GET("/workspace/update", utils.JWTMiddleware(UpdateWorkspace))
	return router
}

func CreateWorkspace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var workspace *types.GBWorkspace = new(types.GBWorkspace)
	utils.DecodeRequestBody(r, &workspace)
	workspace.GOBID = utils.GetUid()
	storage.InsertOne("workspace", workspace)
	resonse, _ := json.Marshal(workspace)
	w.Write(resonse)
}

func GetAllWorkspaceForUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	workspace := new([]types.GBWorkspace)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.FindAll("workspace", user, workspace)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}

func DeleteWorkspace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var filter *types.GBWorkspace = new(types.GBWorkspace)
	utils.DecodeRequestBody(r, &filter)
	filter.GOBID = utils.GetUid()
	storage.DeleteOne("workspace", filter)
	resonse, _ := json.Marshal(filter)
	w.Write(resonse)
}

func UpdateWorkspace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	var updateDoc *types.GBWorkspace = new(types.GBWorkspace)
	var filter *types.GBWorkspace = new(types.GBWorkspace)
	utils.DecodeRequestBody(r, &updateDoc)
	filter.GOBID = utils.GetUid()
	storage.UpdateOne("workspace", filter, updateDoc)
	resonse, _ := json.Marshal(updateDoc)
	w.Write(resonse)
}
