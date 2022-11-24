package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"git/golang/exercise-tracker/data"
	"git/golang/exercise-tracker/web"
)

func main() {
	var err error
	//load environment variables
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Capture DB connection properties.
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "et",
		ParseTime: true,
	}
	// Get a database handle.

	data.DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := data.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// //hard code tests
	// workout, err := workoutByID(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("workout found: %v\n", workout)

	// wID, err := addWorkout(Workout{
	// 	WorkoutName: "B",
	// 	ProgramName: "Stronglifts 5X5",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Workout added: %v\n", wID)

	// err = addExercise(Exercise{
	// 	WorkoutID: wID,
	// 	Name:      "Squat",
	// 	Sets:      5,
	// 	Reps:      5,
	// 	Weight:    280,
	// }, wID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Exercise added: %v\n", wID)

	// err = addExercise(Exercise{
	// 	WorkoutID: wID,
	// 	Name:      "Bench",
	// 	Sets:      5,
	// 	Reps:      5,
	// 	Weight:    180,
	// }, wID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Exercise added: %v\n", wID)

	// err = addExercise(Exercise{
	// 	WorkoutID: wID,
	// 	Name:      "Pendlay Row",
	// 	Sets:      5,
	// 	Reps:      5,
	// 	Weight:    170,
	// }, wID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Exercise added: %v\n", wID)

	//HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", web.IndexHandler)
	http.ListenAndServe(":"+port, mux)
}
