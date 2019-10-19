package main

import (
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchJobList(url string) (JobList, error) {
	jobList := make([]Job, 0)
	doc, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	jobSelection := doc.Find("div.item.job_item")
	jobSelection.Each(func(index int, s *goquery.Selection) {
		job := Job{}
		job.Title = s.Find("h3.item_title").Find("a").Text()
		job.URL, _ = s.Find("a").Attr("href")
		job.Amount = strings.TrimSpace(s.Find("div.entry_data.payment").Find("b.amount").Text())
		job.Expire = strings.TrimSpace(s.Find("div.entry_data.expires").Find("b").Text())
		jobList = append(jobList, job)
	})
	return jobList, nil
}

// for testing purpose
func testFetchJobList() (JobList, error) {
	fileInfos, _ := ioutil.ReadFile("/tmp/test.html")
	stringReader := strings.NewReader(string(fileInfos))

	jobList := make([]Job, 0)
	doc, err := goquery.NewDocumentFromReader(stringReader)

	if err != nil {
		return nil, err
	}

	jobSelection := doc.Find("div.item.job_item")
	jobSelection.Each(func(index int, s *goquery.Selection) {
		job := Job{}
		job.Title = s.Find("h3.item_title").Find("a").Text()
		job.URL, _ = s.Find("a").Attr("href")
		job.Amount = strings.TrimSpace(s.Find("div.entry_data.payment").Find("b.amount").Text())
		job.Expire = strings.TrimSpace(s.Find("div.entry_data.expires").Find("b").Text())
		jobList = append(jobList, job)
	})
	return jobList, nil
}
