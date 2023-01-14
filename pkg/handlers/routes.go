package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/powerslider/flight-tracker/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func registerHTTPRoutes(
	config *configs.ServerConfig, muxer *mux.Router, handler *TrackerHandler) *mux.Router {
	muxer.HandleFunc("/calculate", handler.TracePath()).Methods("POST")

	swaggerJsonURL := fmt.Sprintf("http://%s:%d/swagger/doc.json", config.Host, config.Port)

	muxer.PathPrefix("/swagger/").
		Handler(httpSwagger.Handler(httpSwagger.URL(swaggerJsonURL))) // The url pointing to API definition

	return muxer
}
