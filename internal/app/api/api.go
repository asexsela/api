package api

import (
	"net/http"

	"github.com/asexsela/standart_web_server/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base Api server instance description
type API struct {
	//UNEXPORTED FIELD
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Добавление поля для работы с хранилищем
	storage *storage.Storage
}

// Api constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http searver/configure/loggers/router/database connections etc...
func (api *API) Start() error {

	//Trying to configure Logger
	err := api.configureLoggerField()

	if err != nil {
		return err
	}

	api.logger.Info("Starting api server at port:", api.config.BindAddr)

	//Конфигурируем маршрутизатор
	api.configureRouterField()

	//Конфигурируем хранилище
	err = api.configureStorageField()

	if err != nil {
		return err
	}

	//На этапе валидного завершения, стартуем http server
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
