package main

import (
	"github.com/Tasma-110110/mid1-prj"
	"github.com/Tasma-110110/mid1-prj/package/handler"
	"github.com/Tasma-110110/mid1-prj/package/repository"
	"github.com/Tasma-110110/mid1-prj/package/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		Password: "mid1",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := Config(); err != nil {
		logrus.Fatalf("error init..")
	}

	rep := repository.NewRepository(db)
	servic := service.NewService(rep)
	handlers := handler.NewHandler(servic)

	srv := new(mid1.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error: %s", err.Error())
	}

}
func Config() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
