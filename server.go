package main

import (
	"carApp/app/config"
	"carApp/app/logging"
	"carApp/app/middleware"
	"carApp/app/repository"
	"carApp/app/service"
	tracing "carApp/app/tracing"
	"carApp/database"
	"carApp/graph"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
)

func main() {
	// load log
	appLog := logging.NewAppLogger()
	logConsole := appLog.LogConsole()
	logConsole.Info("start application")

	// load config
	cfg := config.New(appLog)

	// connect to db
	db := database.ConnectDB(cfg)
	fmt.Println(db)

	// connect to jaeger
	tracer, closer := tracing.ConnectJaeger(cfg, appLog, "carApp")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// register repository
	carRepo := repository.NewCarRepository(db)

	// register car service
	carService := service.NewCarProvider(appLog, validator.New(), carRepo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(carService),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", middleware.LogMiddleware(srv, appLog))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Config().App.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Config().App.Port, nil))
}
