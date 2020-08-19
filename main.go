package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-chi/chi/middleware"
	//"../go-restapi/vendor/wrapper"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router .Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	//router.Route("/v1", func(r chi.Router) {
	//	r.Mount("/api/wrapper", wrapper.Routes())
	//})

	return router
}

func main() {
	router := Routes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
//package wrapper
//
//import (
//"net/http"
//
//"github.com/go-chi/chi"
//"github.com/go-chi/render"
//)
//
//type Wrapper struct {
//	Mobile  string `json:"mobile"`
//	Identity string `json:"identity"`
//	Api  string `json:"api"`
//}
//
//func Routes() *chi.Mux {
//	router := chi.NewRouter()
//	router.Post("/query-data", QueryData)
//	return router
//}
//
//func QueryData()  {
//	mobile := chi.URLParam(r, "mobile")
//	identity := chi.URLParam(r, "identity")
//	api := chi.URLParam(r, "api")
//	wrapper := Wrapper{
//		Mobile: mobile,
//		Identity: identity,
//		Api: api,
//	}
//	render.JSON(w, r, wrapper) // A chi router helper for serializing and returning json
//}