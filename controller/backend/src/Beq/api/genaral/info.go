package genaral

import (
	"Beq/nodes/db"
	"fmt"
	"net/http"
)

//Information give details about service
func Information(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	fmt.Println(db.GetDataBase())
	fmt.Fprintf(w, "Welcome home!")
}
