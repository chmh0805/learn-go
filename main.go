package main

import (
	"fmt"
	"learngo/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New(); // initialize.

	addListener(e);

	e.Logger.Fatal(e.Start(":1323"));
}

func handleHome(c echo.Context) error {
	return c.File("views/home.html");
}

func handleScrape(c echo.Context) error {
	keyword := strings.ToLower(scrapper.CleanString(c.FormValue("keyword")));
	fmt.Println(keyword);
	
	defer os.Remove(scrapper.ResultFileName);

	return c.Attachment(scrapper.ResultFileName, fmt.Sprintf("jobs_result_%s.csv", keyword));
}

func addListener(e *echo.Echo) {
	e.GET("/", handleHome);
	e.POST("/scrape", handleScrape);
}