package main

import (
	"bufio"
	"fmt"
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

	writer.WriteString("<html><body><table>\n")
	writer.WriteString("<tr>")
	writer.WriteString("<th>タイトル</th>")
	writer.WriteString("<th>URL</th>")
	writer.WriteString("<th>金額</th>")
	writer.WriteString("<th>期限</th>")
	writer.WriteString("</tr>\n")
	for _, j := range jobList {
		writer.WriteString("<tr>")
		writer.WriteString("<td>" + j.Title + "</td>")
		writer.WriteString("<td>" + j.URL + "</td>")
		writer.WriteString("<td>" + j.Amount + "</td>")
		writer.WriteString("<td>" + j.Expire + "</td>")
		writer.WriteString("</tr>\n")
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

	reader := bufio.NewReaderSize(file, 8192)
	var (
		tmp []byte
	)

	buf := make([]byte, 1024)
	for {
		// nはバイト数を示す
		n, err := reader.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		tmp = append(tmp, buf...)
	}

	str := string(tmp)
	return str
}
