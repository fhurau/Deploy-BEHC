package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	ArticleRoutes(r)
	AuthRoutes(r)
	ConsultationRoutes(r)
}
