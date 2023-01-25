package main

import (
	"fmt"
	"net/http"

	"chess-tournament-register/data"
	"chess-tournament-register/web"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var (
	homeWeb     *web.Web
	addWeb      *web.Web
	updateWeb   *web.Web
	deleteWeb   *web.Web
	notFoundWeb *web.Web
	editWeb     *web.Web
)

func init() {
	data.Dbcon()
}

var port string = ":8080"

func main() {
	homeWeb = web.NewWeb("web/pages/index.gohtml")
	addWeb = web.NewWeb("web/pages/add.gohtml")
	updateWeb = web.NewWeb("web/pages/update.gohtml")
	deleteWeb = web.NewWeb("web/pages/delete.gohtml")
	notFoundWeb = web.NewWeb("web/pages/notfound.gohtml")
	editWeb = web.NewWeb("web/pages/edit.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("web/pages/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/add", add)
	r.HandleFunc("/addauth", addAuth)
	r.HandleFunc("/update", update)
	r.HandleFunc("/updateauth", updateAuth)
	r.HandleFunc("/update/{id}", updateId)
	r.HandleFunc("/delete", delete)
	r.HandleFunc("/delete/{id}", deleteId)

	fmt.Println("Listening port ", port)
	http.ListenAndServe(port, r)
}

type player struct {
	Id       string
	Name     string
	Location string
	Rating   string
	Byes     string
}

var playerId string

func updateId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	playerId = id
	p := []player{}
	allPlayer := data.ShowById(id)
	for i := 0; i < len(allPlayer); i++ {
		id := allPlayer[i]["id"].(string)
		name := allPlayer[i]["name"].(string)
		location := allPlayer[i]["location"].(string)
		rating := allPlayer[i]["rating"].(string)
		byes := allPlayer[i]["byes"].(string)
		p = append(p, player{Id: id, Name: name, Location: location, Rating: rating, Byes: byes})
	}

	err := editWeb.Template.Execute(w, p)
	FetchError(err)
}

func deleteId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	data.DeleteById(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sname := r.FormValue("sname")
	slocation := r.FormValue("slocation")
	srating := r.FormValue("srating")
	sbyes := r.FormValue("sbyes")

	data.UpdatePlayer(sname, slocation, srating, sbyes, playerId)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func home(w http.ResponseWriter, r *http.Request) {
	p := []player{}
	allPlayer := data.ShowAll()
	for i := 0; i < len(allPlayer); i++ {
		id := allPlayer[i]["id"].(string)
		name := allPlayer[i]["name"].(string)
		location := allPlayer[i]["location"].(string)
		rating := allPlayer[i]["rating"].(string)
		byes := allPlayer[i]["byes"].(string)
		p = append(p, player{Id: id, Name: name, Location: location, Rating: rating, Byes: byes})
	}
	err := homeWeb.Template.Execute(w, p)
	FetchError(err)

}

func add(w http.ResponseWriter, r *http.Request) {
	err := addWeb.Template.Execute(w, nil)
	FetchError(err)

}

func addAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("sname")
	location := r.FormValue("slocation")
	rating := r.FormValue("srating")
	byes := r.FormValue("sbyes")
	data.AddPlayer(name, location, rating, byes)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func update(w http.ResponseWriter, r *http.Request) {
	err := updateWeb.Template.Execute(w, nil)
	FetchError(err)
}

func delete(w http.ResponseWriter, r *http.Request) {
	err := deleteWeb.Template.Execute(w, nil)
	FetchError(err)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	err := notFoundWeb.Template.Execute(w, nil)
	FetchError(err)
}

func FetchError(err error) {
	if err != nil {
		panic(err)
	}
}
