package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	todoapp "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig();err != nil{
		logrus.Fatalf("error initializing configs: %s",err.Error())
	}

	if err := godotenv.Load("../.env");err!=nil{
		logrus.Fatalf("error loading env variables: %s",err.Error())
	}

	db ,err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("false to initialize db: %s",err.Error())
	}

	repos :=repository.NewRepository(db)
	services :=service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"),handlers.InitRoutes());err!=nil{
			logrus.Fatalf("error occured while running http server: %s",err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make (chan os.Signal,1)

	signal.Notify(quit,syscall.SIGTERM,syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s",err.Error())
	}
	db.Close()

	
	
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}