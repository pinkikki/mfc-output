package main

import (
	"flag"
	"os"
	"log"
	"encoding/csv"
	"io"
	"./browser"
	"time"
)

const (
	loginUrl = "[記帳アプリのurl]"
	email = "[記帳アプリのe-mail]"
	pass = "[記帳アプリのパスワード]"
	inputUrl = "[記帳アプリの入力url]"
)

func main() {
	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	fail(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	var dataSlice []browser.MFCData
	for {
		row, err := reader.Read()
		if (err == io.EOF) {
			break
		}
		fail(err)
		log.Printf("%#v", row)
		data := browser.MFCData{
			Money:row[4],
			Location:row[2],
			Content:row[3],
			PayFrom:row[6],
		}

		dataSlice = append(dataSlice, data)
	}

	driver := browser.NewDriver()
	defer driver.Stop()
	page := browser.Login(driver, loginUrl, email, pass)
	browser.Get(page, inputUrl)

	for _, data := range dataSlice {
		browser.InputDate(page, data)
	}

	time.Sleep(5 * time.Second)

}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}
