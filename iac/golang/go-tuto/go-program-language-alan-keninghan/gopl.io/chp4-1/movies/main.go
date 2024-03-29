package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

// var mymovies = []Movie{Movie{"jaw", 1982, true, []string{"tom", "jean", "cecile"}},}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	{
		// +Marshal
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshalling failed %s", err)
		}
		fmt.Printf("%s", data)
	}

	{
		//+Marshal with Indention
		data, err := json.MarshalIndent(movies, "", "   ")
		if err != nil {
			log.Fatalf("JSON marshalling with indention failed %s", err)
		}
		fmt.Printf("%s", data)

		//-Unmarshal
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("Json unmarshaling failed %s", err)
		}
		fmt.Println(titles)

	}

}
