package Handler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/ikediufomadu/Go_Spotify_Script/Music"
	"log"
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

	/*
		Define the HTML elements to be scraped and extract music data.
		This will need to be manually changed depending on the website user wants to data scrap.
	*/
	c.OnHTML("div.track-list div[class=track-collection-item__details]", func(e *colly.HTMLElement) {
		music := Music.Music{
			SongName:   e.ChildText("h2[class=track-collection-item__title]"),
			Delimiter:  "|",
			Artist:     e.ChildText("ul[class=artist-list]"),
			Delimiter2: "|",
			Genre:      e.ChildText("[class=genre-list__link]"),
			Delimiter3: "|",
			Published:  e.ChildText("[class=pub-date]"),
		}
		musicData = append(musicData, music)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL, "\n")
	})

	err := c.Visit(website)
	if err != nil {
		return err, website
	}

	// Store the scraped music data to be filtered
	Music.StoreMusicData(musicData)

	return nil, website
}