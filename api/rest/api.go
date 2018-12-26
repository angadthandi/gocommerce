package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/angadthandi/gocommerce/auth"
	"github.com/mongodb/mongo-go-driver/mongo"

	log "github.com/angadthandi/gocommerce/log"
)

type GenericAPIRecieve struct {
	Api     string          `json:"api"`
	Message json.RawMessage `json:"message"`
}

type GenericAPIResponse struct {
	Api     string      `json:"api"`
	Message interface{} `json:"message"`
}

// handler for rest/API
func API(
	w http.ResponseWriter,
	r *http.Request,
	dbRef *mongo.Database,
	jsonMsg json.RawMessage,
) {
	var (
		resp    GenericAPIResponse
		recieve GenericAPIRecieve
	)

	err := json.Unmarshal(jsonMsg, recieve)
	if err != nil {
		log.Errorf("rest/API JSON unmarshal error: %v", err)
		return
	}

	switch recieve.Api {
	case "auth":
		resp.Message = auth.Authenticate(recieve.Message)

	default:
		resp.Message = "Default JSON Message!"
	}

	// Response
	resp.Api = recieve.Api
	b, err := json.Marshal(resp)
	if err != nil {
		log.Errorf("rest/API JSON Marshal error: %v", err)
		return
	}

	log.Debugf("rest/API JSON Response: %v", string(b))
	fmt.Fprintf(w, "rest/API JSON Response: %v", string(b))
}

// handler for home
func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Home Page! %s", r.URL.Path[1:])

	log.Debug("Home Page!")
	http.FileServer(http.Dir("./public/home.html"))
	// http.FileServer(http.Dir("./vendor/home.html"))
}
