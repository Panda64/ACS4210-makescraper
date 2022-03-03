package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/gocolly/colly"
)

type Games struct {
	Game           string `json:"game"`
	CurrentPlayers string `json:"current_players"`
	PeakToday      string `json:"peak_today"`
	GameLink       string `json:"game_link"`
}

func main() {

	var games = []Games{}

	// Instantiate default collector
	c := colly.NewCollector()

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML("#detailStats > table > tbody", func(e *colly.HTMLElement) {

		for count := 3; count < 12; count++ {
			games = append(games, Games{
				Game:           e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(4)"),
				CurrentPlayers: e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(1)"),
				PeakToday:      e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(2)"),
				GameLink:       e.ChildAttr("tr:nth-child("+strconv.Itoa(count)+") > td:nth-child(4) > a", "href"),
			})
		}

		file, _ := json.MarshalIndent(games, "", " ")
		_ = ioutil.WriteFile("game_stats.json", file, 0644)
	})

	// If something goes wrong, print the error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Visit the steam statistics page
	c.Visit("https://store.steampowered.com/stats/Steam-Game-and-Player-Statistics")
}
