package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// CW: Web開発・システム設計
var baseURL = "https://crowdworks.jp/public/jobs/category/241"

var outputFile = "/tmp/output.html"

type Job struct {
	Title  string
	URL    string
	Amount string
	Expire string
}

type JobList []Job

func main() {
	//jobList, err := fetchJobList(baseURL)
	jobList, err := testFetchJobList()
	if err != nil {
		panic(err)
	}

	err = exportHTML(outputFile, jobList)
	if err != nil {
		panic(err)
	}

	htmlBody := readFile(outputFile)
	if err != nil {
		panic(err)
	}
	// TODO: textメール本文を作成
	sendMail("", htmlBody)

	fmt.Println("finish")
}

func exportHTML(filepath string, jobList JobList) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)

	writer.WriteString("<html><body><table><tr>")
	writer.WriteString("<th>タイトル</th>")
	writer.WriteString("<th>URL</th>")
	writer.WriteString("<th>金額</th>")
	writer.WriteString("<th>期限</th>")
	writer.WriteString("</tr>")
	for _, j := range jobList {
		writer.WriteString("<tr>")
		writer.WriteString("<td>" + j.Title + "</td>")
		writer.WriteString("<td>" + j.URL + "</td>")
		writer.WriteString("<td>" + j.Amount + "</td>")
		writer.WriteString("<td>" + j.Expire + "</td>")
		writer.WriteString("</tr>")
	}
	writer.WriteString("</tr></table></html>")

	writer.Flush()
	return nil
}

func readFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, 1000)
	var tmp []byte
	for {
		line, err := reader.ReadBytes(100)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)
	}
	str := string(tmp)
	return str
}
