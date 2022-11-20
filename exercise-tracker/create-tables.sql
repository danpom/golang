CREATE SCHEMA IF NOT EXISTS et;
DROP TABLE IF EXISTS et.workout;
DROP TABLE IF EXISTS et.exercise;
CREATE TABLE et.workout (
  workout_id         INT AUTO_INCREMENT NOT NULL,
  workout_name    VARCHAR(128) NOT NULL,
  program_name    VARCHAR(128) NOT NULL,
  date    DATETIME,
  PRIMARY KEY (`workout_id`)
);
CREATE TABLE et.exercise (
  workout_id        INT NOT NULL,
  exercise_name    VARCHAR(128) NOT NULL,
  sets      int NOT NULL,
  reps     int NOT NULL,
  weight    int NOT NULL,
  CONSTRAINT WI_exercise PRIMARY KEY (workout_id, exercise_name)
);


INSERT INTO et.workout
  (workout_name, program_name, date)
VALUES
  ('A', 'Stronglifts 5X5', now());

INSERT INTO et.exercise
  (workout_id, exercise_name, weight, sets, reps)
VALUES
  ( LAST_INSERT_ID(), 'Squat', 275, 5, 5),
  ( LAST_INSERT_ID(), 'Overhead Press', 125, 5, 5),
  ( LAST_INSERT_ID(), 'Deadlift', 300, 1, 5);