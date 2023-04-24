package main

import (
	"database/sql"

	"github.com/Tasma-110110/mid1-prj"
	"github.com/Tasma-110110/mid1-prj/package/handler"
	"github.com/Tasma-110110/mid1-prj/package/repository"
	"github.com/Tasma-110110/mid1-prj/package/service"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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

type Item struct {
	ID     int
	Name   string
	Price  float64
	Rating float64
}

func (i *Item) AddRating(rating float64, db *sql.DB) error {
	row := db.QueryRow("SELECT rating, num_ratings FROM items WHERE id = ?", i.ID)
	var currentRating float64
	var numR int
	err := row.Scan(&currentRating, &numR)
	if err != nil {
		return err
	}

	newRating := (currentRating*float64(numR) + rating) / float64(numR+1)
	_, err = db.Exec("UPDATE items SET rating = ?, num_ratings = ? WHERE id = ?", newRating, numR+1, i.ID)
	if err != nil {
		return err
	}

	i.Rating = newRating
	return nil
}
func GetItemsPR(minPrice, maxPrice, minRating float64, db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, name, price, rating FROM items WHERE price >= ? AND price <= ? AND rating >= ?", minPrice, maxPrice, minRating)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]Item, 0)
	for rows.Next() {
		var id int
		var name string
		var price float64
		var rating float64
		err = rows.Scan(&id, &name, &price, &rating)
		if err != nil {
			return nil, err
		}
		items = append(items, Item{ID: id, Name: name, Price: price, Rating: rating})
	}

	return items, nil
}

func (i *Item) AddReview(review string, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO reviews (item_id, review) VALUES (?, ?)", i.ID, review)
	if err != nil {
		return err
	}
	return nil
}
func Config() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
