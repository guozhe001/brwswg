package cityapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}

func MainLogic(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "500 - can not Decode request body")
			return
		}
		defer r.Body.Close()
		fmt.Printf("got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "201 - Created")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "405 - Method Not Allowed")
	}
}

func Main() {
	http.HandleFunc("/", MainLogic)
	http.ListenAndServe(":8088", nil)
}
