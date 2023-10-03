package api

import (
	"encoding/json"
	"net/http"

	"github.com/PrithviPy/go-automation-testing/storage"
	"github.com/PrithviPy/go-automation-testing/types"
	"github.com/PrithviPy/go-automation-testing/utils"
	"github.com/julienschmidt/httprouter"
)

func AllHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/", WelcomeUser)
	router.POST("/login", LoginUser)
	router.POST("/create-user", CreateUser)
	router.GET("/get-user", utils.JWTMiddleware(GetUser))
	return router
}

//*************************************************User Account Handlers***************************************************************************//

func WelcomeUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcoem To Go Buddies !"))
}

func LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	utils.DecodeRequestBody(r, &user)
	_, err := storage.FindOne("user", &user)
	var commonRes *types.GBCommongResponse = new(types.GBCommongResponse)
	if err != nil {
		commonRes.Message = "User Not Found !"
		w.WriteHeader(http.StatusForbidden)
	} else {
		token, _ := utils.CreateTokenForUser(user.GOBID)
		commonRes.Token = string(token)
		w.WriteHeader(http.StatusAccepted)
	}
	response, _ := json.Marshal(commonRes)
	w.Write(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	utils.DecodeRequestBody(r, &user)
	user.GOBID = utils.GetUid()
	storage.InsertOne("user", user)
	resonse, _ := json.Marshal(user)
	w.Write(resonse)
}

func GetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	var user *types.GBUser = new(types.GBUser)
	utils.DecodeRequestBody(r, &user)
	storage.FindOne("user", &user)
	resonse, _ := json.Marshal(*user)
	w.Write(resonse)
}

//****************************************************************************************************************************************//
