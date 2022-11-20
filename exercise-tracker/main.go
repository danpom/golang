package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Exercise struct {
	WorkoutID int64
	Name      string
	//Volume map[int]int //k set, v reps
	Sets   int
	Reps   int
	Weight int
}

type Workout struct {
	WorkoutID   int64
	WorkoutName string
	ProgramName string
	Date        time.Time
	Exercises   []Exercise
}

func main() {

	// Capture connection properties.
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "et",
		ParseTime: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	//hard code tests
	workout, err := workoutByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("workout found: %v\n", workout)

	wID, err := addWorkout(Workout{
		WorkoutName: "B",
		ProgramName: "Stronglifts 5X5",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Workout added: %v\n", wID)

	err = addExercise(Exercise{
		WorkoutID: wID,
		Name:      "Squat",
		Sets:      5,
		Reps:      5,
		Weight:    280,
	}, wID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Exercise added: %v\n", wID)

	err = addExercise(Exercise{
		WorkoutID: wID,
		Name:      "Bench",
		Sets:      5,
		Reps:      5,
		Weight:    180,
	}, wID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Exercise added: %v\n", wID)

	err = addExercise(Exercise{
		WorkoutID: wID,
		Name:      "Pendlay Row",
		Sets:      5,
		Reps:      5,
		Weight:    170,
	}, wID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Exercise added: %v\n", wID)
}

// workoutByID queries for workouts that have the specified workout ID.
func workoutByID(id int) (Workout, error) {
	// A workout instance to hold data from returned rows.
	var output Workout

	Exercises, err := exercisesByID(id)
	if err != nil {
		return Workout{}, fmt.Errorf("workoutByID %q: %v", id, err)
	}
	row := db.QueryRow("SELECT * FROM workout WHERE workout_id = ?", id)
	if err := row.Scan(&output.WorkoutID, &output.WorkoutName, &output.ProgramName, &output.Date); err != nil {
		if err == sql.ErrNoRows {
			return output, fmt.Errorf("workoutByID %d: workout not found", id)
		}
		return output, fmt.Errorf("workoutByID %d: %v", id, err)
	}

	output.Exercises = append(output.Exercises, Exercises...)
	return output, nil
}

// exercisesByID queries for exercises that have the specified workout ID.
func exercisesByID(id int) ([]Exercise, error) {
	// An exercise instance to hold data from returned rows.
	var output []Exercise

	rows, err := db.Query("SELECT * FROM exercise WHERE workout_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("exercisesByID %q: %v", id, err)
	}
	defer rows.Close()
	// // Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var temp Exercise
		if err := rows.Scan(&temp.WorkoutID, &temp.Name, &temp.Sets, &temp.Reps, &temp.Weight); err != nil {
			return nil, fmt.Errorf("exercisesByID %q: %v", id, err)
		}
		output = append(output, temp)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("exercisesByID %q: %v", id, err)
	}
	return output, nil
}

// addWorkout adds the specified workour to the database,
// returning the workout ID of the new entry
func addWorkout(work Workout) (int64, error) {
	result, err := db.Exec("INSERT INTO workout (workout_name, program_name, date) VALUES (?, ?, now())", work.WorkoutName, work.ProgramName)
	if err != nil {
		return 0, fmt.Errorf("addWorkout: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addWorkout: %v", err)
	}
	// for _, ex := range work.Exercises {
	// 	err := addExercise(ex, id)
	// 	if err != nil {
	// 		return 0, fmt.Errorf("addWorkout: %v", err)
	// 	}

	// }

	return id, nil
}

// addExercise adds the exercise to the database,
func addExercise(ex Exercise, id int64) error {
	_, err := db.Exec("INSERT INTO exercise (workout_id, exercise_name, sets, reps, weight) VALUES (?, ?, ?, ?, ?)", id, ex.Name, ex.Sets, ex.Reps, ex.Weight)
	if err != nil {
		return fmt.Errorf("addExercise: %v", err)
	}

	return nil
}
