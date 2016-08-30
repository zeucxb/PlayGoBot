package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	var videoURL string
	var viewsQTD int
	// var videoTime int

	fmt.Println("What is the url of your video?")
	fmt.Scan(&videoURL)

	fmt.Println("How mane views do you like?")
	fmt.Scan(&viewsQTD)

	// fmt.Println("What is the length of your video? (ex. 1.30 = 01:30 min)")
	// fmt.Scan(&videoTime)

	// videoTime *= 1000

	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		fmt.Println("Failed to start Selenium:", err)
	}
	page, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		fmt.Println("Failed to open page:", err)
	}

	for viewsQTD > 0 {
		if err := page.Navigate(videoURL); err != nil {
			fmt.Println("Failed to navigate:", err)
		}

		time.Sleep(1000 * time.Millisecond)

		views, _ := page.Find(".watch-view-count").Text()

		fmt.Println("VIEWS:", views)

		viewsQTD--

		fmt.Printf("Views remaining %d\n", viewsQTD)
	}

	if err := driver.Stop(); err != nil {
		fmt.Println("Failed to close pages and stop WebDriver:", err)
	}
}
