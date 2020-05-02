package ch4

import (
	"encoding/json"
	"fmt"
	"log"
)

//Movie
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //解析时用color若是为空则忽略,在这里若是为false则忽略
	Actors []string
}

func TestMovie() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: false,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var newMovie []Movie
	if err := json.Unmarshal(data, &newMovie); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(newMovie) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
