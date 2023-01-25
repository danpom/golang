package data

import (
	"fmt"
	"log"
	"net/url"
	"strconv"

	"github.com/mateors/msql"
)

func AddPlayer(name, location, rating, byes string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "entrants")
	data.Set("dbtype", "sqlite3")
	data.Set("name", name)
	data.Set("location", location)
	data.Set("rating", rating)
	data.Set("byes", byes)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	fmt.Println("Successfully inserted", pid)
	return pid, nil
}

func ShowAll() []map[string]interface{} {
	qs := "SELECT * FROM entrants"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func ShowById(id string) []map[string]interface{} {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	qs := fmt.Sprintf("SELECT * FROM entrants WHERE id=%v;", idInt)
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func UpdatePlayer(name, location, rating, byes, playerId string) (bool, error) {
	idInt, err := strconv.ParseInt(playerId, 10, 64)
	if err != nil {
		log.Println(err)
	}
	qs := fmt.Sprintf("UPDATE entrants SET name = '%s', location = '%s', rating= '%s', byes='%s' WHERE id=%v;", name, location, rating, byes, idInt)
	row := msql.RawSQL(qs, db)
	return row, nil
}

func DeleteById(id string) (bool, error) {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	qs := fmt.Sprintf("DELETE FROM entrants WHERE id=%v;", idInt)
	row := msql.RawSQL(qs, db)
	return row, nil
}
