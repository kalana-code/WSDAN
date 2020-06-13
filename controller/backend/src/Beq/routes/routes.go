package routes

import (
	"Beq/api/genaral"
	"Beq/auth"
	nodeService "Beq/nodes/service"
	ruleService "Beq/rules/service"
	settingService "Beq/settings/service"
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
	r.HandleFunc("/GetNodeInfo", nodeService.GetNodeInfo).Methods("GET", "OPTIONS")
	//POST
	r.HandleFunc("/AddNodeInfo", nodeService.AddNodeInfo).Methods("POST", "OPTIONS")

	//rules handling APIs
	r.HandleFunc("/AddRule", ruleService.AddRule).Methods("POST", "OPTIONS")
	r.HandleFunc("/RemoveFlow/{FlowID}", ruleService.RemoveRulesByFlowID).Methods("GET", "OPTIONS")
	r.HandleFunc("/RemoveRule/{RuleID}", ruleService.RemoveRuleByRuleID).Methods("GET", "OPTIONS")
	r.HandleFunc("/GetAllRules", ruleService.GetAllRules).Methods("GET", "OPTIONS")
	// r.HandleFunc("/GetRulesInFlow", ruleService.AddRule).Methods("POST", "OPTIONS")

	//System setting handling APIs
	r.HandleFunc("/StateToggle", settingService.Toggle).Methods("GET", "OPTIONS")
	r.HandleFunc("/SystemSetting", settingService.GetCurrentSetting).Methods("GET", "OPTIONS")

	// user registrarion
	// r.HandleFunc("/Register/Student", genaral.Register).Methods("POST", "OPTIONS")
	r.HandleFunc("/User/Login", genaral.Login).Methods("POST", "OPTIONS")

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
