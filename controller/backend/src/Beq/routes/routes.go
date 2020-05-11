package routes

import (
	"Beq/api/genaral"
	"Beq/auth"
	"Beq/nodes/service"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers function Used for arrange routes
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	//Node information handling
	//GET
	r.HandleFunc("/Info", genaral.Information).Methods("GET", "OPTIONS")
	r.HandleFunc("/GetNodeInfo", service.GetNodeInfo).Methods("GET", "OPTIONS")
	//POST
	r.HandleFunc("/AddNodeInfo", service.AddNodeInfo).Methods("POST", "OPTIONS")

	// user registrarion
	// r.HandleFunc("/Register/Student", genaral.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/Student/Login", genaral.Login).Methods("POST", "OPTIONS")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	// use middleware
	s.Use(auth.JwtVerify)
	s.HandleFunc("/verify", genaral.Verify).Methods("GET", "OPTIONS")
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
