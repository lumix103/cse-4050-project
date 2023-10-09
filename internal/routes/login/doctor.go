package login

import (
	"fmt"
	"net/http"
)

func Doctor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DOCTOR Login Page TODO")
}
