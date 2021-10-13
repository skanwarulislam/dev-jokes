package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/markbates/goth/gothic"
	"html/template"
	"io/ioutil"
	"log"
	"makeajoke/models"
	"net/http"
	"strings"
)

func DoAuth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func DoCallback(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		log.Fatal(res, err)
		return
	}
	joke, err := makeJoke()
	if err != nil {
		log.Fatal(err)
	}
	provider := strings.Split(req.URL.Path, "/")[2]
	log.Println("User is logged in using :"+provider)
	data := &models.Data{
		User:        user,
		CompanyName: "InterestingJokes AS",
		Joke:        joke,
		Provider:    provider,
	}
	htmlFile := []string{
		"templates/success.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/content.html",
	}
	t, _ := template.ParseFiles(htmlFile...)
	err = t.Execute(res, data)
	if err != nil {
		log.Fatal(err)
	}
}

func DoLogout(res http.ResponseWriter, req *http.Request) {
	err := gothic.Logout(res, req)
	if err != nil {
		log.Fatal(err)
	}
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func makeJoke() (models.ChuckNorris, error) {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random?category=dev")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var chuckNorris models.ChuckNorris
	err = json.Unmarshal(body, &chuckNorris)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chuckNorris)
	return chuckNorris, nil
}

func ShowIndex(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(res, false)
	if err != nil {
		log.Fatal(err)
	}
}
