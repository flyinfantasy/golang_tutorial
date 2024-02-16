package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/spf13/viper"
)

type WebInfo struct {
	Uid              string `json:"uid"`
	Pigeon_club      string `json:"pigeon_club"`
	Pigeonloft       string `json:"pigeon_loft"`
	Pigeonnumber     string `json:"pigeon_number"`
	Webreceivedate   string `json:"get_web_receive_date"`
	Webreceivetime   string `json:"get_web_receive_time"`
	Weblongitude     string `json:"get_web_longitude"`
	Weblatitude      string `json:"get_web_latitude"`
	Webnumber        string `json:"get_web_number"`
	Webinductiondate string `json:"get_web_Induction_date"`
	Webinductiontime string `json:"get_web_Induction_time"`
	Webcheck         string `json:"get_web_check"`
}

type config struct {
	Day      string `yaml:"day"`
	Url      string `yaml:"url"`
	Urlf     string `yaml:"urlf"`
	Urlb     string `yaml:"urlb"`
	Areacode string `yaml:"areacode"`
	Koipc    string `yaml:"koipc"`
	Status   string `yaml:"status"`
}

var (
	conf *config
)

func init() {
	// Retrieve config options.
	conf = getConf()
}

func main() {
	fmt.Print(GetCreatorPigeonClubData(conf.Url))
}

func getConf() *config {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

func GetCreatorPigeonClubData(url string) int {
	page := GetPigeonGodPages(url)
	count := 0
	for i := 0; i <= page; i++ {
		count = count + GetCreatorPigeonClubPageData(i)
	}
	return count
}

func datatojson(webinfo WebInfo) []byte {
	webinfojson, err := json.Marshal(webinfo)
	if err != nil {
		fmt.Println(err)
	}
	return webinfojson
}

func GetPigeonGodPages(Url string) int {
	var page int

	c := colly.NewCollector()

	c.OnHTML(".inline", func(e *colly.HTMLElement) {
		format := "Page: 1 / %d　Total: 6927"
		//fmt.Println(e.Text)
		fmt.Sscanf(e.Text, format, &page)
		// e.Text 印出 <a> tag 裡面的文字，也就是文章標題
		// e.Attr("href") 則是找到 <a> tag裡面的 href元素
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(Url)

	return page
}

func GetCreatorPigeonClubPageData(page int) int {
	c := colly.NewCollector()
	count := 0
	var webinductiondate string
	var webinductiontime string
	var webpigeonloft string
	var webnumber string
	c.OnHTML("table > tbody > tr", func(e *colly.HTMLElement) {
		format := "%d"
		var id int
		fmt.Sscanf(e.ChildText("td:nth-child(1)"), format, &id)
		if id > 0 {
			fmt.Println(id, " ")
			//fmt.Print(e.ChildText("td:nth-child(2)"), " c ")
			//fmt.Print(e.ChildText("td:nth-child(3)"), " d ")
			//fmt.Print(e.ChildText("td:nth-child(4)"), " e ")
			fmt.Print(conf.Day, "XXXXX")
			webinductiondate = conf.Day
			timetemp := strings.Split(e.ChildText("td:nth-child(5)"), ".")
			webinductiontime = timetemp[0]
			webpigeonloft = e.ChildText("td:nth-child(3)")
			webnumber = e.ChildText("td:nth-child(4)")
			//fmt.Print(e.ChildText("td:nth-child(5)"), " f ")
			fmt.Println()
			uid := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s", conf.Areacode, "01", conf.Koipc, conf.Status, webpigeonloft, webnumber, webinductiondate)
			//c.Visit("https://www.11568.com.tw/index.asp?Page=1&QSite=&QSysid=1507&QRaceDate=2022%2F03%2F19&QSort=0&QSiteCode=&QMode=train&qsize=1000&p=N&QPass=")
			webinfo := WebInfo{
				Uid:              uid,
				Pigeon_club:      conf.Koipc,
				Pigeonloft:       webpigeonloft,
				Pigeonnumber:     webnumber,
				Webnumber:        webnumber,
				Webinductiondate: webinductiondate,
				Webinductiontime: webinductiontime,
				Webcheck:         "1",
			}

			webinfoj := datatojson(webinfo)
			fmt.Println(string(webinfoj))
			Posttosite(webinfoj)
			count++
		}
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(GetPigeonGodUrl(page))

	return count
}

func GetPigeonGodUrl(page int) string {
	return fmt.Sprintf("%s%d%s", conf.Urlf, page, conf.Urlb)
}

func Posttosite(webinfo []byte) {
	httpposturl := "http://172.104.98.231:8282/api/v1/pigeons/web/"
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(webinfo))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
