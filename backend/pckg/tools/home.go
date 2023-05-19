package tools

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 HOMEPAGE NOT FOUND...", 404)
		return
	}

	return

}
