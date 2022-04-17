package multiplemiddleware

import (
	"fmt"
	"github.com/guozhe001/brwswg/chapter3/cityapi"
	"log"
	"net/http"
	"strconv"
	"time"
)

func FilterContentType(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("currently in the check content type middleware")
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			fmt.Fprintln(w, "415 - Unsupported Media Type, Please send JSON")
			return
		}
		handlerFunc.ServeHTTP(w, r)
	})
}

func SetServerTimeCookie(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerFunc.ServeHTTP(w, r)
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("currently in set server time middleware")
	})
}

func Main() {
	mainHandlerFunc := http.HandlerFunc(cityapi.MainLogic)
	http.Handle("/", FilterContentType(FilterContentType(mainHandlerFunc)))
	http.ListenAndServe(":8088", nil)
}
