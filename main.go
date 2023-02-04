package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type player struct {
	Name     string
	Team     string
	Position string
	Height   string
	Weight   string
	Age      string
}

func parseCSV(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	r, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for _, r := range r {
		player := player{
			Name:     r[0],
			Team:     r[1],
			Position: r[2],
			Height:   r[3],
			Weight:   r[4],
			Age:      r[5],
		}
		//fmt.Println(player.Name + " | " + player.Team + " | " + player.Position + " | " + player.Height + " | " + player.Weight + " | " + player.Age)
		fmt.Fprintln(w, player.Name+"\t"+player.Team+"\t"+player.Position+"\t"+player.Height+"\t"+player.Weight+"\t"+player.Age)
	}
	w.Flush()
}

func main() {
	parseCSV("./sample.csv")

}
