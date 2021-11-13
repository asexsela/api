package api

import (
	"net/http"

	"github.com/asexsela/standart_web_server/internal/app/middleware"
	"github.com/asexsela/standart_web_server/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods(http.MethodGet)

	// a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods(http.MethodGet)
	//With JWT
	a.router.Handle(prefix+"/articles/{id}", middleware.JwtMiddleware.Handler(
		http.HandlerFunc(a.GetArticleById),
	)).Methods(http.MethodGet)

	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods(http.MethodDelete)
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods(http.MethodPost)
	a.router.HandleFunc(prefix+"/user/register", a.PostUserRegister).Methods(http.MethodPost)
	//auth
	a.router.HandleFunc(prefix+"/user/auth", a.PostAuth).Methods(http.MethodPost)
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
