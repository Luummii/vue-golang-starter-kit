package server

import (
	"log"
	"net/http"
	"os"
	"projects/vue-golang-starter-kit/backend/db"
	"time"

	"projects/vue-golang-starter-kit/backend/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// SrvData -
type SrvData struct {
	Port         string
	DataBaseName string
	URL          string
	User         string
	Password     string
	DataBase     db.DataBase
	router       *mux.Router
}

// SrvInterface -
type SrvInterface interface {
	Routes()
	Run()
}

// Create -
func Create(port, dataBasebName, url, user, password string) (SrvInterface, error) {
	db := db.DataBase{}
	return &SrvData{Port: ":" + port, DataBaseName: dataBasebName, URL: url, User: user, Password: password, DataBase: db}, nil
}

// ConnectDB -
func (srv *SrvData) ConnectDB() *db.Store {
	s, err := srv.DataBase.Connect(srv.DataBaseName, srv.User, srv.URL, srv.Password)
	if err != nil {
		log.Printf("SERVER:-->::ConnectDB:::ERROR::NO CONNECTED IN DB: %s\n", err)
		panic(err)
	}
	return s
}

// Routes -
func (srv *SrvData) Routes() {
	store := srv.ConnectDB()
	handler := &handler.Store{DB: store}

	r := mux.NewRouter().StrictSlash(true)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/add", handler.AddUsers).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/").HandlerFunc(indexHandler("./index.html")).Methods("GET")
	srv.router = r
}

func indexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}
	return http.HandlerFunc(fn)
}

// Run -
func (srv *SrvData) Run() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(originsOk, headersOk, methodsOk)(srv.router))
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	server := &http.Server{
		Handler:      handler,
		Addr:         srv.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server listening on port%s", srv.Port)
	log.Fatal(server.ListenAndServe())
}
