package table8

import (
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
	"table8/app"
	"log"
)

func init() {
	log.Println("table8_restaurant_api.init() ... starting")
	rpc_v2 := rpc.NewServer()
	rpc_v2.RegisterCodec(json2.NewCodec(), "application/json")
	app.Register(rpc_v2)

	log.Println("table8_restaurant_api.init() ... ready")
	http.Handle("/restaurant/0.1/", rpc_v2)
}
