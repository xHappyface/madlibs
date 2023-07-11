package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:           "localhost:8080",
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    10 * time.Second,
		ErrorLog:       log.Default(),
	}
	fs := http.FileServer(http.Dir("./static/"))
	mux.HandleFunc("/award_ceremonies_madlib", handleAwardCeremoniesMadlib)
	mux.HandleFunc("/bowl_games_madlib", handleBowlGamesMadlib)
	mux.HandleFunc("/christmas_blizzard_madlib", handleChristmasBlizzardMadlib)
	mux.HandleFunc("/christmas_clothing_madlib", handleChristmasClothingMadlib)
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			switch r.URL.Path {
			case "/", "/index":
				http.ServeFile(w, r, "./static/index.gohtml")
			case "/award_ceremonies", "/award_ceremonies_form":
				http.ServeFile(w, r, "./static/award_ceremonies_form.gohtml")
			case "/bowl_games", "/bowl_games_form":
				http.ServeFile(w, r, "./static/bowl_games_form.gohtml")
			case "/christmas_blizzard", "/christmas_blizzard_form":
				http.ServeFile(w, r, "./static/christmas_blizzard_form.gohtml")
			case "/christmas_clothing", "/christmas_clothing_form":
				http.ServeFile(w, r, "./static/christmas_clothing_form.gohtml")
			case "/christmas_cookies", "/christmas_cookies_form":
				http.ServeFile(w, r, "./static/christmas_cookies_form.gohtml")
			case "/christmas_dinner", "/christmas_dinner_form":
				http.ServeFile(w, r, "./static/christmas_dinner_form.gohtml")
			case "/christmas_july", "/christmas_july_form":
				http.ServeFile(w, r, "./static/christmas_july_form.gohtml")
			case "/christmas_pageant", "/christmas_pageant_form":
				http.ServeFile(w, r, "./static/chistmas_pageant_form.gohtml")
			case "/christmas_shopping", "/christmas_shopping_form":
				http.ServeFile(w, r, "./static/christmas_shopping_form.gohtml")
			case "/christmas_solo", "/christmas_solo_form":
				http.ServeFile(w, r, "./static/christmas_solo_form.gohtml")
			case "/christmas_vacation", "/christmas_vacation_form":
				http.ServeFile(w, r, "./static/christmas_vacation_form.gohtml")
			case "/dancing_cool", "/dancing_cool_form":
				http.ServeFile(w, r, "./static/dancing_cool_form.gohtml")
			case "/decorating_tree", "decorating_tree_form":
				http.ServeFile(w, r, "./static/decorating_tree_form.gohtml")
			case "/figure_skating", "/figure_skating_form":
				http.ServeFile(w, r, "./static/figure_skating_form.gohtml")
			case "/going_caroling", "/going_caroling_form":
				http.ServeFile(w, r, "./static/going_caroling_form.gohtml")
			case "/happy_messages", "/happy_messages_form":
				http.ServeFile(w, r, "./static/happy_messages_form.gohtml")
			case "/holiday_traveling", "/holiday_traveling_form":
				http.ServeFile(w, r, "./static/holiday_traveling_form.gohtml")
			case "/hot_chocolate", "/hot_chocolate_form":
				http.ServeFile(w, r, "./static/hot_chocolate_form.gohtml")
			case "/igloo_facts", "/igloo_facts_form":
				http.ServeFile(w, r, "./static/igloo_facts_form.gohtml")
			case "/lodge", "/lodge_form":
				http.ServeFile(w, r, "./static/lodge_form.gohtml")
			case "/music", "/music_form":
				http.ServeFile(w, r, "./static/music_form.gohtml")
			case "/naughty_list", "/naughty_list_form.gohtml":
				http.ServeFile(w, r, "./static/naughty_list_form.gohtml")
			case "/nutcracker", "/nutcracker_form":
				http.ServeFile(w, r, "./static/nutcracker_form.gohtml")
			case "/popular_gifts", "/popular_gifts_form":
				http.ServeFile(w, r, "./static/popular_gifts_form.gohtml")
			case "/reindeer", "/reindeer_form":
				http.ServeFile(w, r, "./static/reindeer_form.gohtml")
			case "/school_party", "/school_party_form":
				http.ServeFile(w, r, "./static/school_party_form.gohtml")
			case "/selecting_tree", "/selecting_tree_form":
				http.ServeFile(w, r, "./static/selecting_tree_form.gohtml")
			case "/skating_champ", "/skating_champ_form":
				http.ServeFile(w, r, "./static/skating_champ_form.gohtml")
			case "/snow_day", "/snow_day_form":
				http.ServeFile(w, r, "./static/snow_day_form.gohtml")
			case "/snowboarding", "/snowboarding_form":
				http.ServeFile(w, r, "./static/snowboarding_form.gohtml")
			case "/snowed_in", "/snowed_in_form":
				http.ServeFile(w, r, "./static/snowed_in_form.gohtml")
			case "/snowman_building", "/snowman_buidling_form":
				http.ServeFile(w, r, "./static/snowman_building_form.gohtml")
			case "/specialty_house", "/specialty_house_form":
				http.ServeFile(w, r, "./static/specialty_house_form.gohtml")
			case "/toys_kids", "/toys_kids_form":
				http.ServeFile(w, r, "./static/toys_kids_form.gohtml")
			case "/transylvania", "/transylvania_form":
				http.ServeFile(w, r, "./static/transylvania_form.gohtml")
			case "/tropical_christmas", "/tropical_christmas_form":
				http.ServeFile(w, r, "./static/tropical_christmas_form")
			case "/weather_report", "/weather_report_form":
				http.ServeFile(w, r, "./static/weather_report_form.gohtml")
			default:
				fs.ServeHTTP(w, r)
			}
		case http.MethodPost:
			switch r.URL.Path {
			case "/award_ceremonies":
				handleAwardCeremoniesSubmit(w, r)
			case "/bowl_games":
				handleBowlGamesSubmit(w, r)
			case "/christmas_blizzard":
				handleChristmasBlizzardSubmit(w, r)
			case "/christmas_clothing":
				handleChristmasClothingSubmit(w, r)
			default:
				http.NotFound(w, r)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("ERR:", err)
	}
}

func handleAwardCeremoniesSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	queryParams := url.Values{}
	queryParams.Set("Adjective1", r.Form.Get("Adjective1"))
	queryParams.Add("Noun1", r.Form.Get("Noun1"))
	queryParams.Add("PluralNoun1", r.Form.Get("PluralNoun1"))
	queryParams.Add("PluralNoun2", r.Form.Get("PluralNoun2"))
	queryParams.Add("Noun2", r.Form.Get("Noun2"))
	queryParams.Add("PartOfTheBody", r.Form.Get("PartOfTheBody"))
	queryParams.Add("Adjective2", r.Form.Get("Adjective2"))
	queryParams.Add("Noun3", r.Form.Get("Noun3"))
	queryParams.Add("Adjective3", r.Form.Get("Adjective3"))
	queryParams.Add("PluralNoun3", r.Form.Get("PluralNoun3"))
	queryParams.Add("Adjective4", r.Form.Get("Adjective4"))
	queryParams.Add("Adjective5", r.Form.Get("Adjective5"))
	queryParams.Add("PluralNoun4", r.Form.Get("PluralNoun4"))
	queryParams.Add("PluralNoun5", r.Form.Get("PluralNoun5"))

	redirectURL := r.URL.String() + "_madlib?" + queryParams.Encode()

	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func handleAwardCeremoniesMadlib(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	args := [14]string{
		queryParams.Get("Adjective1"), queryParams.Get("Noun1"),
		queryParams.Get("PluralNoun1"), queryParams.Get("PluralNoun2"),
		queryParams.Get("Noun2"), queryParams.Get("PartOfTheBody"),
		queryParams.Get("Adjective2"), queryParams.Get("Noun3"),
		queryParams.Get("Adjective3"), queryParams.Get("PluralNoun3"),
		queryParams.Get("Adjective4"), queryParams.Get("Adjective5"),
		queryParams.Get("PluralNoun4"), queryParams.Get("PluralNoun5"),
	}
	for _, arg := range args {
		if arg == "" {
			http.Error(w, "empty arguments", http.StatusBadRequest)
			return
		}
	}
	if len(args) != 14 {
		http.Error(w, "invalid number of arguments", http.StatusInternalServerError)
		return
	}
	tpl, err := template.ParseFiles("./templates/award_ceremonies_madlib.gohtml")
	if err != nil {
		http.Error(w, "parsing error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Adjective1, Noun1, PluralNoun1, PluralNoun2, Noun2, PartOfTheBody,
		Adjective2, Noun3, Adjective3, PluralNoun3, Adjective4, Adjective5,
		PluralNoun4, PluralNoun5 string
	}{
		Adjective1: args[0], Noun1: args[1], PluralNoun1: args[2], PluralNoun2: args[3],
		Noun2: args[4], PartOfTheBody: args[5], Adjective2: args[6], Noun3: args[7],
		Adjective3: args[8], PluralNoun3: args[9], Adjective4: args[10], Adjective5: args[11],
		PluralNoun4: args[12], PluralNoun5: args[13],
	}
	if err = tpl.Execute(w, data); err != nil {
		http.Error(w, "executing error", http.StatusInternalServerError)
	}
}

func handleBowlGamesSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	queryParams := url.Values{}
	queryParams.Set("Adjective1", r.Form.Get("Adjective1"))
	queryParams.Add("Noun1", r.Form.Get("Noun1"))
	queryParams.Add("City1", r.Form.Get("City1"))
	queryParams.Add("AnimalPlural1", r.Form.Get("AnimalPlural1"))
	queryParams.Add("AnotherCity1", r.Form.Get("AnotherCity1"))
	queryParams.Add("AnotherAnimalPlural", r.Form.Get("AnotherAnimalPlural"))
	queryParams.Add("TypeOfFlower", r.Form.Get("TypeOfFlower"))
	queryParams.Add("City2", r.Form.Get("City2"))
	queryParams.Add("TypeOfBirdPlural", r.Form.Get("TypeOfBirdPlural"))
	queryParams.Add("AnotherCity2", r.Form.Get("AnotherCity2"))
	queryParams.Add("SomethingAlivePlural", r.Form.Get("SomethingAlivePlural"))
	queryParams.Add("AnotherTypeOfBirdPlural", r.Form.Get("AnotherTypeOfBirdPlural"))
	queryParams.Add("PluralNoun1", r.Form.Get("PluralNoun1"))
	queryParams.Add("TypeOfVegetable", r.Form.Get("TypeOfVegetable"))
	queryParams.Add("City3", r.Form.Get("City3"))
	queryParams.Add("AnotherCity3", r.Form.Get("AnotherCity3"))
	queryParams.Add("AnimalPlural2", r.Form.Get("AnimalPlural2"))
	queryParams.Add("Noun2", r.Form.Get("Noun2"))
	queryParams.Add("FullName", r.Form.Get("FullName"))
	queryParams.Add("TypeOfContainer", r.Form.Get("TypeOfContainer"))
	queryParams.Add("PartOfTheHouse", r.Form.Get("PartOfTheHouse"))
	queryParams.Add("Color", r.Form.Get("Color"))

	redirectURL := r.URL.String() + "_madlib?" + queryParams.Encode()

	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func handleBowlGamesMadlib(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	args := [22]string{
		queryParams.Get("Adjective1"), queryParams.Get("Noun1"),
		queryParams.Get("City1"), queryParams.Get("AnimalPlural1"),
		queryParams.Get("AnotherCity1"), queryParams.Get("AnotherAnimalPlural"),
		queryParams.Get("TypeOfFlower"), queryParams.Get("City2"),
		queryParams.Get("TypeOfBirdPlural"), queryParams.Get("AnotherCity2"),
		queryParams.Get("SomethingAlivePlural"), queryParams.Get("AnotherTypeOfBirdPlural"),
		queryParams.Get("PluralNoun1"), queryParams.Get("TypeOfVegetable"),
		queryParams.Get("City3"), queryParams.Get("AnotherCity3"),
		queryParams.Get("AnimalPlural2"), queryParams.Get("Noun2"),
		queryParams.Get("FullName"), queryParams.Get("TypeOfContainer"),
		queryParams.Get("PartOfTheHouse"), queryParams.Get("Color"),
	}
	for _, arg := range args {
		if arg == "" {
			http.Error(w, "empty arguments", http.StatusBadRequest)
			return
		}
	}
	if len(args) != 22 {
		http.Error(w, "invalid number of arguments", http.StatusInternalServerError)
		return
	}
	tpl, err := template.ParseFiles("./templates/bowl_games_madlib.gohtml")
	if err != nil {
		http.Error(w, "parsing error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Adjective1, Noun1, City1, AnimalPlural1, AnotherCity1, AnotherAnimalPlural,
		TypeOfFlower, City2, TypeOfBirdPlural, AnotherCity2, SomethingAlivePlural,
		AnotherTypeOfBirdPlural, PluralNoun1, TypeOfVegetable, City3, AnotherCity3,
		AnimalPlural2, Noun2, FullName, TypeOfContainer, PartOfTheHouse, Color string
	}{
		Adjective1: args[0], Noun1: args[1], City1: args[2], AnimalPlural1: args[3],
		AnotherCity1: args[4], AnotherAnimalPlural: args[5], TypeOfFlower: args[6],
		City2: args[7], TypeOfBirdPlural: args[8], AnotherCity2: args[9],
		SomethingAlivePlural: args[10], AnotherTypeOfBirdPlural: args[11],
		PluralNoun1: args[12], TypeOfVegetable: args[13], City3: args[14], AnotherCity3: args[15],
		AnimalPlural2: args[16], Noun2: args[17], FullName: args[18], TypeOfContainer: args[19],
		PartOfTheHouse: args[20], Color: args[21],
	}
	if err = tpl.Execute(w, data); err != nil {
		http.Error(w, "executing error", http.StatusInternalServerError)
	}
}

func handleChristmasBlizzardSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	queryParams := url.Values{}
	queryParams.Set("Adjective1", r.Form.Get("Adjective1"))
	queryParams.Add("Adjective2", r.Form.Get("Adjective2"))
	queryParams.Add("PluralNoun", r.Form.Get("PluralNoun"))
	queryParams.Add("Adjective3", r.Form.Get("Adjective3"))
	queryParams.Add("IngVerb", r.Form.Get("IngVerb"))
	queryParams.Add("Adjective4", r.Form.Get("Adjective4"))
	queryParams.Add("Noun1", r.Form.Get("Noun1"))
	queryParams.Add("Noun2", r.Form.Get("Noun2"))
	queryParams.Add("Noun3", r.Form.Get("Noun3"))
	queryParams.Add("Noun4", r.Form.Get("Noun4"))
	queryParams.Add("Noun5", r.Form.Get("Noun5"))
	queryParams.Add("Noun6", r.Form.Get("Noun6"))
	queryParams.Add("Adjective5", r.Form.Get("Adjective5"))

	redirectURL := r.URL.String() + "_madlib?" + queryParams.Encode()

	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func handleChristmasBlizzardMadlib(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	args := [13]string{
		queryParams.Get("Adjective1"), queryParams.Get("Adjective2"),
		queryParams.Get("PluralNoun"), queryParams.Get("Adjective3"),
		queryParams.Get("IngVerb"), queryParams.Get("Adjective4"),
		queryParams.Get("Noun1"), queryParams.Get("Noun2"),
		queryParams.Get("Noun3"), queryParams.Get("Noun4"),
		queryParams.Get("Noun5"), queryParams.Get("Noun6"),
		queryParams.Get("Adjective5"),
	}
	for _, arg := range args {
		if arg == "" {
			http.Error(w, "empty arguments", http.StatusBadRequest)
			return
		}
	}
	if len(args) != 13 {
		http.Error(w, "invalid number of arguments", http.StatusInternalServerError)
		return
	}
	tpl, err := template.ParseFiles("./templates/christmas_blizzard_madlib.gohtml")
	if err != nil {
		http.Error(w, "parsing error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Adjective1, Adjective2, PluralNoun, Adjective3, IngVerb, Adjective4,
		Noun1, Noun2, Noun3, Noun4, Noun5, Noun6, Adjective5 string
	}{
		Adjective1: args[0], Adjective2: args[1], PluralNoun: args[2],
		Adjective3: args[3], IngVerb: args[4], Adjective4: args[5],
		Noun1: args[6], Noun2: args[7], Noun3: args[8], Noun4: args[9],
		Noun5: args[10], Noun6: args[11], Adjective5: args[12],
	}
	if err = tpl.Execute(w, data); err != nil {
		http.Error(w, "executing error", http.StatusInternalServerError)
	}
}

func handleChristmasClothingSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	queryParams := url.Values{}
	queryParams.Set("PluralNoun1", r.Form.Get("PluralNoun1"))
	queryParams.Add("PluralNoun2", r.Form.Get("PluralNoun2"))
	queryParams.Add("Adjective1", r.Form.Get("Adjective1"))
	queryParams.Add("WomanInTheRoom", r.Form.Get("WomanInTheRoom"))
	queryParams.Add("Adjective2", r.Form.Get("Adjective2"))
	queryParams.Add("PluralNoun3", r.Form.Get("PluralNoun3"))
	queryParams.Add("PluralNoun4", r.Form.Get("PluralNoun4"))
	queryParams.Add("PartOfTheBody1", r.Form.Get("PartOfTheBody1"))
	queryParams.Add("PluralNoun5", r.Form.Get("PluralNoun5"))
	queryParams.Add("Color", r.Form.Get("Color"))
	queryParams.Add("Place", r.Form.Get("Place"))
	queryParams.Add("Noun", r.Form.Get("Noun"))
	queryParams.Add("PartOfTheBody2", r.Form.Get("PartOfTheBody2"))
	queryParams.Add("PastTenseVerb", r.Form.Get("PastTenseVerb"))
	queryParams.Add("PartOfTheBody3", r.Form.Get("PartOfTheBody3"))

	redirectURL := r.URL.String() + "_madlib?" + queryParams.Encode()

	http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
}

func handleChristmasClothingMadlib(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	args := [15]string{
		queryParams.Get("PluralNoun1"), queryParams.Get("PluralNoun2"),
		queryParams.Get("Adjective1"), queryParams.Get("WomanInTheRoom"),
		queryParams.Get("Adjective2"), queryParams.Get("PluralNoun3"),
		queryParams.Get("PluralNoun4"), queryParams.Get("PartOfTheBody1"),
		queryParams.Get("PluralNoun5"), queryParams.Get("Color"),
		queryParams.Get("Place"), queryParams.Get("Noun"),
		queryParams.Get("PartOfTheBody2"), queryParams.Get("PastTenseVerb"),
		queryParams.Get("PartOfTheBody3"),
	}
	for _, arg := range args {
		if arg == "" {
			http.Error(w, "empty arguments", http.StatusBadRequest)
			return
		}
	}
	if len(args) != 15 {
		http.Error(w, "invalid number of arguments", http.StatusInternalServerError)
		return
	}
	tpl, err := template.ParseFiles("./templates/christmas_clothing_madlib.gohtml")
	if err != nil {
		http.Error(w, "parsing error", http.StatusInternalServerError)
		return
	}
	data := struct {
		PluralNoun1, PluralNoun2, Adjective1, WomanInTheRoom, Adjective2,
		PluralNoun3, PluralNoun4, PartOfTheBody1, PluralNoun5, Color,
		Place, Noun, PartOfTheBody2, PastTenseVerb, PartOfTheBody3 string
	}{
		PluralNoun1: args[0], PluralNoun2: args[1], Adjective1: args[2],
		WomanInTheRoom: args[3], Adjective2: args[4], PluralNoun3: args[5],
		PluralNoun4: args[6], PartOfTheBody1: args[7], PluralNoun5: args[8],
		Color: args[9], Place: args[10], Noun: args[11], PartOfTheBody2: args[12],
		PastTenseVerb: args[13], PartOfTheBody3: args[14],
	}
	if err = tpl.Execute(w, data); err != nil {
		http.Error(w, "executing error", http.StatusInternalServerError)
	}
}
