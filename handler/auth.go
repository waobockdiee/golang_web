package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type credential struct {
// 	Username string `json:"username"`
// 	Password string `json:"pasword"`
// }

func Run(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	// if decoder.Decode()

	date_time := decoder

	// if err != nil {
	// 	fmt.Fprintln(rw, "Error al traer los datos")
	// 	return
	// }

	fmt.Fprintln(rw, date_time)
}
