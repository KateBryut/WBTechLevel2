package main

import (
	"context"
	"dev11/internal/config"
	"dev11/internal/server"
	"dev11/internal/server/handler"
	"dev11/internal/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	repo, err := storage.NewMemoryStorage()
	if err != nil {
		log.Fatalf("error initializing storage: %s", err.Error())
	}
	handler := handler.NewHandler(repo)
	srv := server.NewServer()
	go func() {
		if err := srv.Run(cfg, handler.InitRoutes()); err != nil {
			log.Fatalf("error running server: %s", err.Error())
		}
	}()
	log.Print("server is successfully started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("server is shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}
}
