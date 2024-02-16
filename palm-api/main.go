package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
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

type MyRespEnvelope struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName                             xml.Name
	QueryRaceTrainData_AfterNumResponse QueryRaceTrainData_AfterNumResponse `xml:"QueryRaceTrainData_AfterNumResponse"`
}

type QueryRaceTrainData_AfterNumResponse struct {
	XMLName                           xml.Name
	QueryRaceTrainData_AfterNumResult string `xml:"QueryRaceTrainData_AfterNumResult"`
}

type config struct {
	Url       string `yaml:"url"`
	SID       string `yaml:"sID"`
	SPW       string `yaml:"sPW"`
	SType     string `yaml:"sType"`
	VerCode   string `yaml:"verCode"`
	Areacode  string `yaml:"areacode"`
	Koipc     string `yaml:"koipc"`
	Status    string `yaml:"status"`
	Startfrom string `yaml:"startfrom"`
	Loftdigit string `yaml:"loft"`
}

var (
	conf *config
)

func init() {
	// Retrieve config options.
	conf = getConf()
}

func main() {
	/*
		<sID>`jon`</sID>
		<sPW>6180</sPW>
		<sType>T</sType>
		<sGetAfterNum>1</sGetAfterNum>
		<verCode>Rex2022KouYi</verCode>
	*/
	//fmt.Println(Getwinapi("http://60.248.233.195:86/", "jon", "6180", "T", "199", "Rex2022KouYi"))
	Runapiserver(conf.Url, conf.SID, conf.SPW, conf.SType, conf.VerCode, conf.Areacode, conf.Koipc, conf.Status, conf.Startfrom)
	//Runapiserver("http://60.248.233.195:86/", "wtph", "3678", "T", "Rex2022KouYi", "04", "00001", conf.Status)
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

func Runapiserver(url string, sID string, sPW string, sType string, verCode string, areacode string, koipc string, status string, start string) {
	endtime := time.Now().Hour() + 12
	sGetAfterNum, _ := strconv.Atoi(start)
	if endtime > 24 {
		endtime = endtime - 24
	}
	for time.Now().Hour() != endtime {
		data := Getwinapi(url, sID, sPW, sType, sGetAfterNum, verCode)
		datanum := analyzedata(data, areacode, koipc, status)
		if datanum > 0 {
			sGetAfterNum = sGetAfterNum + datanum
		}
		fmt.Println("\n", sGetAfterNum)
		time.Sleep(20 * time.Second)
	}
}

func analyzedata(data string, areacode string, koipc string, status string) int {
	sdata := strings.Split(data, "'")
	datanum := len(sdata)
	for i := 1; i <= datanum-2; i = i + 1 {
		fmt.Println(string(sdata[i]))
		webinfo := datatojson(analyzesdata(sdata[i], areacode, koipc, status))
		//fmt.Println("webinfo", string(webinfo))
		Posttosite(webinfo)
		//time.Sleep(5*time.Second)
	}
	return datanum - 2
}

func analyzesdata(sdata string, areacode string, koipc string, status string) WebInfo {
	ssdata := strings.Split(sdata, ";")

	//receivessdata := strings.Split(ssdata[2], " ")
	webreceivedatey := ssdata[5][17:19]
	webreceivedatem := ssdata[5][19:21]
	webreceivedated := ssdata[5][21:23]
	webreceivedate := fmt.Sprintf("%s%s-%s-%s", "20", webreceivedatey, webreceivedatem, webreceivedated)
	webreceivetimeh := ssdata[5][23:25]
	webreceivetimem := ssdata[5][25:27]
	webreceivetimes := ssdata[5][27:29]
	webreceivetime := fmt.Sprintf("%s:%s:%s", webreceivetimeh, webreceivetimem, webreceivetimes)
	weblongitude := ssdata[4][1:8]
	weblatitude := ssdata[4][8:16]
	webnumber := ssdata[5][7:14]
	inductiondate := strings.Split(ssdata[3], " ")
	webinductiondate := inductiondate[0]
	webinductiontime := inductiondate[1]
	fmt.Printf(ssdata[1])
	webpigeonloft := makepigeonloft(ssdata[1], conf.Loftdigit)
	uid := fmt.Sprintf("%s_%s_%s_%s_%s_%s_%s", areacode, "01", koipc, status, webpigeonloft, webnumber, webreceivedate)
	webinfo := WebInfo{
		Uid:              uid,
		Pigeon_club:      koipc,
		Pigeonloft:       webpigeonloft,
		Pigeonnumber:     webnumber,
		Webreceivedate:   webreceivedate,
		Webreceivetime:   webreceivetime,
		Weblongitude:     weblongitude,
		Weblatitude:      weblatitude,
		Webnumber:        webnumber,
		Webinductiondate: webinductiondate,
		Webinductiontime: webinductiontime,
		Webcheck:         "1",
	}
	return webinfo
}

func makepigeonloft(pigeonloft string, digits string) string {
	if digits == "4" {
		pigeonloft = pigeonloft[0:4]
	} else if digits == "A3" {
		if !IsNum(pigeonloft[0:1]) {
			pigeonloft = fmt.Sprintf("%s%s", "0", pigeonloft[1:4])
		}
	} else if digits == "3" {
		pigeonloft = fmt.Sprintf("%s%s", "0", pigeonloft[0:3])
	} else if digits == "2" {
		pigeonloft = fmt.Sprintf("%s%s", "00", pigeonloft[0:2])
	}
	//pigeonloft = fmt.Sprintf("%s%s", "0", pigeonloft[0:3])
	//fmt.Println(pigeonloft)
	return pigeonloft
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func datatojson(webinfo WebInfo) []byte {
	webinfojson, err := json.Marshal(webinfo)
	if err != nil {
		fmt.Println(err)
	}
	return webinfojson
}

func Posttosite(webinfo []byte) {
	httpposturl := "http://172.104.98.231:8282/api/v1/pigeons/web/"
	fmt.Println("webinfo", string(webinfo))
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

func Getwinapi(url string, sID string, sPW string, sType string, intGetAfterNum int, verCode string) string {
	sGetAfterNum := strconv.Itoa(intGetAfterNum)
	reqBody :=
		`<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
		<soap12:Body>
			<QueryRaceTrainData_AfterNum xmlns="http://tempuri.org/">
					<sID>` + sID + `</sID>
					<sPW>` + sPW + `</sPW>
					<sType>` + sType + `</sType>
					<sGetAfterNum>` + sGetAfterNum + `</sGetAfterNum>
					<verCode>` + verCode + `</verCode>
			</QueryRaceTrainData_AfterNum>
		</soap12:Body>
	</soap12:Envelope>`

	res, err := http.Post(url, "application/soap+xml; charset=utf-8", strings.NewReader(reqBody))
	if nil != err {
		fmt.Println("http post err:", err)
	}
	defer res.Body.Close()
	// return status
	if http.StatusOK != res.StatusCode {
		fmt.Println("WebService soap1.2 request fail, status: %s\n", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if nil != err {
		fmt.Println("ioutil ReadAll err:", err)
	}

	v := MyRespEnvelope{}
	err = xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return v.Body.QueryRaceTrainData_AfterNumResponse.QueryRaceTrainData_AfterNumResult
}
