package main

import (
	"flag"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/quanergyo/ozon-test-assingment/repository"
	"github.com/quanergyo/ozon-test-assingment/repository/postgres"
	"github.com/spf13/viper"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/quanergyo/ozon-test-assingment/graph"
)

const defaultPort = "8080"

type Mode struct {
	memory bool
}

func main() {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		slog.Error("Error: loading env variables", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	memoryFlag := Mode{}
	memoryFlag.initFlag()

	var db *sqlx.DB = nil

	if memoryFlag.memory == false {
		dbCpy, err := postgres.NewDB(postgres.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		})

		if err != nil {
			slog.Error("Error: failde to init db connection", err)
			os.Exit(1)
		}
		db = dbCpy
	}

	repos := repository.NewRepository(db)
	resolver := graph.NewResolver(repos)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			slog.Error("Error: failed to start server on port:", viper.GetString("port"), err.Error())
			os.Exit(1)
		}
	}()

	slog.Info("Start server")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if memoryFlag.memory == false {
		if err := db.Close(); err != nil {
			slog.Error("error occured on close db connection:", err)
			os.Exit(1)
		}
	}
	slog.Info("Server shutting down")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func (m *Mode) initFlag() {
	flag.BoolVar(&m.memory, "memory", false, "memory mode")
	flag.Parse()
}
