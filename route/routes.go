package route

import (
	"io/ioutil"
	"net/http"

	"github.com/angadthandi/gocommerce/dbconnect/test"
	"github.com/angadthandi/gocommerce/gosocket"
	"github.com/angadthandi/gocommerce/registry"

	"github.com/angadthandi/gocommerce/api/rest"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"

	log "github.com/angadthandi/gocommerce/log"
)

func Handle(
	dbRef *mongo.Database,
	hub *gosocket.Hub,
	reg *registry.Registry,
) {
	r := mux.NewRouter().StrictSlash(true)

	// r.HandleFunc("/", rest.Home)
	r.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {

		b := postDataHandler(r)

		rest.API(w, r, dbRef, reg, b)
	})

	r.HandleFunc("/ws/{token}", func(w http.ResponseWriter, r *http.Request) {

		varsMap := mux.Vars(r)
		token, ok := varsMap["token"]
		if !ok {
			log.Error("Invalid token!")
			return
		}

		gosocket.ServeWs(
			hub,
			w,
			r,
			dbRef,
			reg,
			token,
		)
	})

	// Test Routes --------------------------------
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		test.TestHandler(w, r, dbRef)
	})
	// Test Routes --------------------------------

	// // static files
	// r.HandleFunc("/vendor/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.StripPrefix("/vendor/",
	// 		http.FileServer(http.Dir("./public")))
	// })
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)
}

func postDataHandler(
	r *http.Request,
) []byte {
	// w.Header().Set("Content-Type", "application/json")

	// var m Member
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("error reading post data: %v", err)
		return nil
	}
	//   json.Unmarshal(b, &m)

	//   members = append(members, m)

	//   j, _ := json.Marshal(m)
	//   w.Write(j)

	return b
}
