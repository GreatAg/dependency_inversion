package main

import (
	"dependency_inversion/handler"
)

func init_routes() {
	sample := app.Group("/sample")
	sample.Get("/dependency_inversion", handler.GetUser)
}
