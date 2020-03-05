package routes

import (
	"Beq/api/genaral"
	"Beq/auth"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers function Used for arrange routes
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/Register/Student", genaral.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/Student/Login", genaral.Login).Methods("POST", "OPTIONS")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	// use middleware
	s.Use(auth.JwtVerify)
	s.HandleFunc("/verify", genaral.Verify).Methods("GET", "OPTIONS")
	// s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	// s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	// s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding,X-Access-Token")
		next.ServeHTTP(w, r)
	})
}
