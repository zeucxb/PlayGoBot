package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	var line, digited string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Show me the code :)")
	for {
		scanner.Scan()

		line = scanner.Text()

		if line == ":end" {
			break
		}

		digited += fmt.Sprintf("%s\n", line)
	}

	fmt.Println("========== PROGRAM ==========")

	fmt.Println(digited)

	fmt.Println("========== RESPONSE ==========")

	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		fmt.Println("Failed to start Selenium:", err)
	}
	page, err := driver.NewPage(agouti.Browser("firefox"))
	if err != nil {
		fmt.Println("Failed to open page:", err)
	}

	if err := page.Navigate("https://play.golang.org/p/m-70CPrfAC"); err != nil {
		fmt.Println("Failed to navigate:", err)
	}

	_, err = page.URL()
	if err != nil {
		fmt.Println("Failed to get page URL:", err)
	}

	page.Find("#code").SendKeys(digited)

	page.Find("#format").Click()

	page.Find("#run").Click()

	time.Sleep(1000 * time.Millisecond)

	output, err := page.Find("#output").Text()
	if err != nil {
		fmt.Println("Failed to get login prompt text:", err)
	}

	fmt.Println(output)

	if err := driver.Stop(); err != nil {
		fmt.Println("Failed to close pages and stop WebDriver:", err)
	}
}
