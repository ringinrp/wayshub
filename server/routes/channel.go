package routes

import (
	"wayshub/handlers"
	"wayshub/pkg/middleware"
	"wayshub/pkg/mysql"
	"wayshub/repositories"

	"github.com/gorilla/mux"
)

func ChannelRoutes(r *mux.Router) {
	channelRepository := repositories.RepositoryChannel(mysql.DB)
	h := handlers.HandlerChannel(channelRepository)

	r.HandleFunc("/channels", middleware.Auth(h.FindChannels)).Methods("GET")
	r.HandleFunc("/channel/{id}", middleware.Auth(h.GetChannel)).Methods("GET")
	r.HandleFunc("/mychannel", middleware.Auth(h.GetMyChannel)).Methods("GET")
	r.HandleFunc("/editchannel", middleware.Auth((middleware.UploadPhoto(middleware.UploadCover(h.EditChannel))))).Methods("PATCH")
	r.HandleFunc("/channel/{id}", middleware.Auth(h.DeleteChannel)).Methods("DELETE")

	r.HandleFunc("/plusSubs/{id}", middleware.Auth(h.PlusSubscriber)).Methods("PATCH")
	r.HandleFunc("/minusSubs/{id}", middleware.Auth(h.MinusSubscriber)).Methods("PATCH")
}
