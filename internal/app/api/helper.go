package api

import (
	"net/http"

	"github.com/asexsela/standart_web_server/storage"
	"github.com/sirupsen/logrus"
)

//Пытаемся отконфигурировать наш API инстанс (а конкретнее - поле Logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)

	if err != nil {
		return err
	}

	a.logger.SetLevel(log_level)

	return nil
}

//Пытаемся отконфигурировать маршрутизатор (а конкретнее router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is rest api"))
	})
}

//Пытаемся конфигурировать наше хранилище (а конкретно storage API)
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//Пытаемся установить соединение, если не возможно возвращаем ошибку
	err := storage.Open()

	if err != nil {
		return err
	}

	a.storage = storage

	return nil
}
