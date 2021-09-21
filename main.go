package main

import (
	"os"
	"strings"

	"github.com/junhoKim/learngo/scrapper"
	"github.com/labstack/echo"
)

const flleName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "job.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScape)
	e.Logger.Fatal(e.Start(":1323"))
}
