package main

import (
	"github.com/Lolodin/kodix.git/internal/route"
	serverapi "github.com/Lolodin/kodix.git/internal/server"
	"github.com/Lolodin/kodix.git/internal/store"
	"net/http"
)

func main() {
	// init store
	mstore:= store.NewMemory()



	http.HandleFunc("/product/", route.GetProduct(mstore))
	http.HandleFunc("/", route.IndexHandler())
	http.HandleFunc("/list", route.ProductList(mstore))
	http.HandleFunc("/delete/", route.DeleteProduct(mstore))
	http.HandleFunc("/put", route.UpdateProduct(mstore))
	http.HandleFunc("/new", route.CreateProduct(mstore))

	//static
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("./kodix-app/build/static/js"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./kodix-app/build/static/css"))))
	http.Handle("/icon/", http.StripPrefix("/icon/", http.FileServer(http.Dir("./kodix-app/build"))))
	http.Handle("/static/media/", http.StripPrefix("/static/media/", http.FileServer(http.Dir("./kodix-app/build/static/media"))))
	s := serverapi.NewServer(":8080")
	s.ListenAndServe()
}
