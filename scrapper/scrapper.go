package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractJobItem struct {
	title string;
	postURL string;
	companyName string;
	companyURL string;
	expireDate string;
	location string;
	careerRequirement string;
	eduRequirement string;
	jobType string;
}

// Scrape Saramin using a keyword
func Scrape(keyword string) {
	var baseURL string = "https://www.saramin.co.kr";
	var searchURL string = baseURL + "/zf_user/search/recruit?&searchword=";

	fmt.Println("Enter Keyword what you want to search:");
	fmt.Scanln(&keyword);

	searchURL = searchURL + keyword;
	fmt.Println("It will return result of ", searchURL, "...");

	var jobs []extractJobItem;
	mainChan := make(chan []extractJobItem);

	totalPages := getPageCount(searchURL);
	
	for i := 0; i < totalPages; i++ {
		go getPage(i, mainChan, baseURL, searchURL);
	}
	
	for i := 0; i < totalPages; i++ {
		extractedJobs := <- mainChan;
		jobs = append(jobs, extractedJobs...);
	}

	writeJobs(jobs);
}

func writeJobs(jobs []extractJobItem) {
	file, err := os.Create("jobs.csv");
	checkErr(err);

	w := csv.NewWriter(file);
	defer w.Flush();

	headers := []string{
		"Title",
		"PostURL",
		"CompanyName",
		"CompanyURL",
		"Expire Date",
		"Location",
		"Career Requirement",
		"Education Requirement",
		"Job Type",
	};

	writeDataToFile(w, headers);

	for _, job := range jobs {
		writeJobDataToFile(file, w, job);
	}
}

func writeDataToFile(w *csv.Writer, data []string) {
	wErr := w.Write(data);
	checkErr(wErr);
}

func writeJobDataToFile(file *os.File, w *csv.Writer, job extractJobItem) {
	jobStr := []string{
		job.title,
		job.postURL,
		job.companyName,
		job.companyURL,
		job.expireDate,
		job.location,
		job.careerRequirement,
		job.eduRequirement,
		job.jobType,
	};
	jwErr := w.Write(jobStr);
	checkErr(jwErr);
	// for UTF-8 Encoding
	utf8bom := []byte{0xEF, 0xBB, 0xBF};
	file.Write(utf8bom);
}

func getPage(pageNum int, mainChan chan<- []extractJobItem, baseURL string, searchURL string) {
	var jobItems []extractJobItem;
	c := make(chan extractJobItem);

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
		go extractJob(card, c, baseURL);
	});

	for i := 0; i < foundItems.Length(); i++ {
		job := <-c;
		jobItems = append(jobItems, job);
	}

	mainChan <- jobItems;
}

func extractJob(card *goquery.Selection, c chan<- extractJobItem, baseURL string) {
	postTitle, _ := card.Find("h2.job_tit a").Attr("title");
	postURL, _ := card.Find("h2.job_tit a").Attr("href");
	companyName := CleanString(card.Find("strong.corp_name a").Text());
	companyURL, _ := card.Find("strong.corp_name a").Attr("href");
	jobDate := CleanString(card.Find("div.job_date span.date").Text());

	var jobLoc string;
	var careerReq string;
	var eduReq string;
	var jobType string;

	card.Find("div.job_condition span").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			jobLoc = CleanString(s.Text());
			break;
		case 1:
			careerReq = CleanString(s.Text());
			break;
		case 2:
			eduReq = CleanString(s.Text());
			break;
		case 3:
			jobType = CleanString(s.Text());
			break;
		default:
			break;
		}
	});

	c <- extractJobItem{
		title: CleanString(postTitle),
		postURL: baseURL + CleanString(postURL),
		companyName: companyName,
		companyURL: baseURL + CleanString(companyURL),
		expireDate: jobDate,
		location: jobLoc,
		careerRequirement: careerReq,
		eduRequirement: eduReq,
		jobType: jobType,
	};
}

// CleanString cleans a string, like trim()
func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ");
}

func getPageCount(searchURL string) int {
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