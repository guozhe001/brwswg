package alice

import (
	"github.com/guozhe001/brwswg/chapter3/cityapi"
	"github.com/guozhe001/brwswg/chapter3/multiplemiddleware"
	"github.com/justinas/alice"
	"net/http"
)

func Main() {
	mainHandlerFunc := http.HandlerFunc(cityapi.MainLogic)
	chain := alice.New(multiplemiddleware.FilterContentType, multiplemiddleware.SetServerTimeCookie).Then(mainHandlerFunc)
	http.Handle("/", chain)
	http.ListenAndServe(":8088", nil)
}
