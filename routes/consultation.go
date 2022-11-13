package routes

import (
	"backend/handlers"
	"backend/pkg/middleware"
	"backend/pkg/mysql"
	"backend/repositories"

	"github.com/gorilla/mux"
)

func ConsultationRoutes(r *mux.Router) {
	consultationRepository := repositories.RepositoryConsultation(mysql.DB)
	h := handlers.HandlerConsultation(consultationRepository)

	r.HandleFunc("/consultation", middleware.Auth(h.CreateConsultation)).Methods("POST")
	r.HandleFunc("/consultations", middleware.Auth(h.FindConsultation)).Methods("GET")
	r.HandleFunc("/consultation/{id}", h.GetConsultation).Methods("GET")
	r.HandleFunc("/consultation/{id}", middleware.Auth(h.UpdateConsultation)).Methods("PATCH")
	r.HandleFunc("/consultation/{id}", middleware.Auth(h.DeleteConsultation)).Methods("DELETE")

}
