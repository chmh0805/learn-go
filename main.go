package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr";
var searchURL string = baseURL + "/zf_user/search/recruit?&searchword=python";

type extractJobItem struct {
	title string;
	postURL string;
	companyName string;
	companyURL string;
	expireDate string;
	location string;
	careerCondition string;
	eduCondition string;
	jobType string;
}

func main() {
	var jobs []extractJobItem;
	totalPages := getPageCount();
	
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i);
		jobs = append(jobs, extractedJobs...);
	}

	fmt.Println(jobs);
}

func getPage(pageNum int) []extractJobItem {
	var jobItems []extractJobItem;

	pageURL := searchURL + "&recruitPage=" + strconv.Itoa(pageNum + 1);
	fmt.Println("RequestURL: ", pageURL);

	res, err := http.Get(pageURL);
	checkErr(err);
	checkCode(res);

	defer res.Body.Close();

	doc, err := goquery.NewDocumentFromReader(res.Body);
	checkErr(err);

	// parse Items
	foundItems := doc.Find(".item_recruit");
	foundItems.Each(func(i int, card *goquery.Selection) {
		var item extractJobItem = extractJob(card);
		jobItems = append(jobItems, item);
	});
	return jobItems;
}

func extractJob(card *goquery.Selection) extractJobItem {
	postTitle, _ := card.Find("h2.job_tit a").Attr("title");
	postURL, _ := card.Find("h2.job_tit a").Attr("href");
	companyName := cleanString(card.Find("strong.corp_name a").Text());
	companyURL, _ := card.Find("strong.corp_name a").Attr("href");
	jobDate := cleanString(card.Find("div.job_date span.date").Text());

	var jobLoc string;
	var careerCond string;
	var eduCond string;
	var jobType string;

	card.Find("div.job_condition span").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			jobLoc = cleanString(s.Text());
			break;
		case 1:
			careerCond = cleanString(s.Text());
			break;
		case 2:
			eduCond = cleanString(s.Text());
			break;
		case 3:
			jobType = cleanString(s.Text());
			break;
		default:
			break;
		}
	});

	return extractJobItem{
		title: cleanString(postTitle),
		postURL: baseURL + cleanString(postURL),
		companyName: companyName,
		companyURL: baseURL + cleanString(companyURL),
		expireDate: jobDate,
		location: jobLoc,
		careerCondition: careerCond,
		eduCondition: eduCond,
		jobType: jobType,
	};
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ");
}

func getPageCount() int {
	pageCount := 0;

	req, reqErr := http.NewRequest("GET", searchURL, nil);
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