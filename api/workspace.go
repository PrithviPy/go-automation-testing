package api

import (
	"encoding/json"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/types"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func AllWorkspcaeHandlers() {

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
