package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandlerRecive(w http.ResponseWriter, r *http.Request) {
	var reciveUser ReciveFromUserDTO
	if err := json.NewDecoder(r.Body).Decode(&reciveUser); err != nil {
		fmt.Println("err ", err)
		return
	}

	fmt.Fprintln(w, "сообщение доставлено")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(reciveUser)

}
