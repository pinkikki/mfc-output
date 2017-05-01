package browser

import (
	"github.com/sclevine/agouti"
	"log"
)

func NewDriver() *agouti.WebDriver {
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return driver
}

func Login(driver *agouti.WebDriver, url string, email string, pass string) *agouti.Page {

	page, err := driver.NewPage()
	if err != nil {
		panic(err)
	}

	Get(page, url)

	page.FindByID("sign_in_session_service_email").Fill(email)
	page.FindByID("sign_in_session_service_password").Fill(pass)
	if err := page.FindByID("login-btn-submit").Submit(); err != nil {
		panic(err)
	}

	return page
}

func Get(page *agouti.Page, url string) {

	if err := page.Navigate(url); err != nil {
		panic(err)
	}
}

func InputDate(page *agouti.Page, data MFCData) {
	page.FindByID("journal_value").Fill(data.Money)
	b := make([]byte, 0, 50)
	b = append(b, data.Location...)
	b = append(b, "\r\n"...)
	b = append(b, data.Content...)
	page.FindByID("journal_remark").Fill(string(b))
	page.FindByID("journal_recognized_at").Fill(data.PayFrom)
	if err := page.FindByClass("ca-btn-save ca-btn-size-xsmall").Submit(); err != nil {
		panic(err)
	}
}

type MFCData struct {
	Money    string
	Location string
	Content  string
	PayFrom  string
}