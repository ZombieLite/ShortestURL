package script

import (
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/formatlink", CoreLink)
	http.ListenAndServe(":8080", nil)
}
