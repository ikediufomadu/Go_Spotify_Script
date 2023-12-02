package Handler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/ikediufomadu/Go_Spotify_Script/Music"
	"github.com/ikediufomadu/Go_Spotify_Script/Query"
	"log"
	"strings"
)

// WebConnector initiates the web scraping process for the target website
func WebConnector(targetWebsite string) {
	err, website := webScrapper(targetWebsite)
	if err != nil {
		log.Fatal("Error connecting to ", website)
	}
}

// webScrapper performs web scraping on the provided website
func webScrapper(website string) (error, string) {
	c := colly.NewCollector()

	var musicData []Music.Music

	// Looks for the div that stores all the songs
	c.OnHTML("div.o-chart-results-list-row-container", func(e *colly.HTMLElement) {
		// Looks for the H3 element that has the class 'c-title'
		e.ForEach("h3.c-title", func(_ int, h3 *colly.HTMLElement) {
			// Grabs the parent element of the H3 element
			parent := h3.DOM.Parent()
			// Checks if the parent is the li element that contains the song name and artist name children elements.
			if parent.Is("li.o-chart-results-list__item") {
				// Formats retrieved data
				songName := strings.TrimSpace(h3.Text)

				// Gets the span element with the class 'c-label', this is where the artist name is stored.
				artist := strings.TrimSpace(parent.Find("span.c-label").Text())

				music := Music.Music{
					SongName: songName,
					Artist:   artist,
				}
				musicData = append(musicData, music)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL, "\n")
	})

	err := c.Visit(website)
	if err != nil {
		return err, website
	}

	// Perform Spotify queries with the scraped music data
	Query.ScrapedMusic(musicData)

	return nil, website
}
