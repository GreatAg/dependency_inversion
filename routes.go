package main

import (
	"dependency_inversion/handler"
)

func init_routes() {
	sample := app.Group("/sample")
	sample.Get("/getUser/:id", handler.GetUser)
	sample.Post("/updateUser", handler.UpdateUser)
}
