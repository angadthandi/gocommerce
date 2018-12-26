package route

import (
	"io/ioutil"
	"net/http"

	"github.com/angadthandi/gocommerce/gosocket"

	"github.com/angadthandi/gocommerce/api/rest"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"

	log "github.com/angadthandi/gocommerce/log"
)

func Handle(
	dbRef *mongo.Database,
	hub *gosocket.Hub,
) {
	r := mux.NewRouter().StrictSlash(true)

	// r.HandleFunc("/", rest.Home)
	r.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {

		b := postDataHandler(r)

		rest.API(w, r, hub, dbRef, b)
	})

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		gosocket.ServeWs(
			hub,
			w,
			r,
			dbRef,
		)
	})

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
