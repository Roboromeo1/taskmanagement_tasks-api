package application

import (
	"github.com/Roboromeo1/taskmanagement_tasks-api/src/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.TasksController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.TasksController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.TasksController.Search).Methods(http.MethodPost)
}
