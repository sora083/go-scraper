package main

import (
	"fmt"
	"net/smtp"
	"os"
	"time"

	"bytes"
	"encoding/base64"
	"strings"

	"crypto/rand"
	"encoding/binary"
	"strconv"

	"github.com/google/uuid"
)

const layout = "2006-01-02"

func sendMail(textBody, htmlBody string) {

	day := time.Now()

	// TODO 環境変数などから取得
	mailFrom := "test@gmail.com"
	mailTo := "test@gmail.com"
	host := "localhost"
	subject := fmt.Sprintf("ジョブリスト(%s)", day.Format(layout))

	// UUID生成
	randomId, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uid := randomId.String()

	var textBodyBuffer bytes.Buffer
	textBodyBuffer.WriteString(textBody)

	var htmlBodyBuffer bytes.Buffer
	htmlBodyBuffer.WriteString(htmlBody)

	var message bytes.Buffer
	boundary := ""

	// mail header
	message.WriteString("From: " + mailFrom + "\r\n")
	message.WriteString("To: " + mailTo + "\r\n")
	message.WriteString(encodeSubject(subject))
	message.WriteString("Mime-Version: 1.0\r\n")

	if textBody != "" && htmlBody != "" {
		message.WriteString("Content-Type: multipart/alternative; boundary=\"" + uid + "\"\r\n")
		message.WriteString("\r\n")

		boundary = "--" + uid + "\r\n"
	}

	// text/plain mail body
	if textBody != "" {
		message.WriteString(boundary)
		message.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString("\r\n")
		message.WriteString(add76crlf(base64.StdEncoding.EncodeToString(textBodyBuffer.Bytes())))
		message.WriteString("\r\n")
		message.WriteString("\r\n")
	}

	// text/html mail body
	if htmlBody != "" {
		message.WriteString(boundary)
		message.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
		message.WriteString("Content-Transfer-Encoding: base64\r\n")
		message.WriteString("\r\n")
		message.WriteString(add76crlf(base64.StdEncoding.EncodeToString(htmlBodyBuffer.Bytes())))
		message.WriteString("\r\n")
		message.WriteString("\r\n")
	}

	// mail body end
	if textBody != "" && htmlBody != "" {
		message.WriteString("--" + uid + "--\r\n")
	}

	if err := smtp.SendMail(host+":25", nil, mailFrom, []string{mailTo}, []byte(message.String())); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("sending mail to %s\n", mailTo)
}

// サブジェクトをMIMEエンコードする
func encodeSubject(subject string) string {
	var buffer bytes.Buffer
	buffer.WriteString("Subject:")
	for _, line := range utf8Split(subject, 13) {
		buffer.WriteString(" =?utf-8?B?")
		buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		buffer.WriteString("?=\r\n")
	}
	return buffer.String()
}

// 76バイト毎にCRLFを挿入する
func add76crlf(msg string) string {
	var buffer bytes.Buffer
	for k, c := range strings.Split(msg, "") {
		buffer.WriteString(c)
		if k%76 == 75 {
			buffer.WriteString("\r\n")
		}
	}
	return buffer.String()
}

// UTF8文字列を指定文字数で分割
func utf8Split(utf8string string, length int) []string {
	resultString := []string{}
	var buffer bytes.Buffer
	for k, c := range strings.Split(utf8string, "") {
		buffer.WriteString(c)
		if k%length == length-1 {
			resultString = append(resultString, buffer.String())
			buffer.Reset()
		}
	}
	if buffer.Len() > 0 {
		resultString = append(resultString, buffer.String())
	}
	return resultString
}

// ランダムな文字列を生成
func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}
