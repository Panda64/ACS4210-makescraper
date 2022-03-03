package main

import (
	"fmt"

	"strconv"

	"github.com/gocolly/colly"
)

type Games struct {
	current_players string
	peak_today      string
	game            string
	game_link       string
}

func main() {

	var games = []Games{}

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#detailStats > table > tbody", func(e *colly.HTMLElement) {

		for count := 3; count < 12; count++ {
			games = append(games, Games{
				current_players: e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(1)"),
				peak_today:      e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(2)"),
				game:            e.ChildText("tr:nth-child(" + strconv.Itoa(count) + ") > td:nth-child(4)"),
				game_link:       e.ChildAttr("tr:nth-child("+strconv.Itoa(count)+") > td:nth-child(4) > a", "href"),
			})
		}

		// Print games
		fmt.Println(games)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// If something goes wrong, print the error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Visit the steam statistics page
	c.Visit("https://store.steampowered.com/stats/Steam-Game-and-Player-Statistics")
}
