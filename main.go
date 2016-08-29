package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	var digited, imports string

	fmt.Println("What you want to import?")
	fmt.Scan(&imports)

	fmt.Println("Show me the code :)")
	fmt.Scan(&digited)

	ims := strings.Split(imports, ",")

	imports = "import("

	for i := range ims {
		imports += fmt.Sprintf("\"%s\"\n", ims[i])
	}

	imports += ")"

	digited = fmt.Sprintf("package main\n%s\nfunc main() { %s }", imports, digited)

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

	loginPrompt, err := page.Find("#output").Text()
	if err != nil {
		fmt.Println("Failed to get login prompt text:", err)
	}

	fmt.Println(loginPrompt)

	// loginPrompt, err = page.Find("#code").Text()
	// if err != nil {
	// 	fmt.Println("Failed to get login prompt text:", err)
	// }

	// fmt.Println("code:", loginPrompt)

	// if err := driver.Stop(); err != nil {
	// 	fmt.Println("Failed to close pages and stop WebDriver:", err)
	// }
}
