package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python";

func main() {
	totalPages := getPageCount();
	
	for i := 0; i < totalPages; i++ {
		getPage(i);
	}
}

func getPage(pageNum int) {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(pageNum + 1);
	fmt.Println("RequestURL: ", pageURL);
}

func getPageCount() int {
	pageCount := 0;

	req, reqErr := http.NewRequest("GET", baseURL, nil);
	checkErr(reqErr);

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36");

	client := &http.Client{}
	res, err := client.Do(req);
	checkErr(err);
	checkCode(res);

	defer res.Body.Close();

	doc, err := goquery.NewDocumentFromReader(res.Body);
	checkErr(err);

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pageCount = s.Find("a").Length();
	})

	return pageCount;
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err);
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode);
	}
	// log.Println("Request Status Code: ", res.StatusCode);
}