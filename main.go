package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//  Ref: https://www.golanglearn.com/http-get-rest-example-using-golang/
//  Ref: https://www.soberkoder.com/consume-rest-api-go/
//  https://api.ipgeolocation.io/ipgeo?apiKey=2caffbb7123943d7933b1fe729994186&ip=

// Response is the response from the api
type Response []struct {
	IP            string `json:"ip"`
	Continentcode string `json:"continent_code"`
	Continentname string `json:"continent_name"`
	Countrycode2  string `json:"country_code2"`
	Countrycode3  string `json:"country_code3"`
	Countryname   string `json:"country_name"`
	//Countrycapital string `json:"country_capital"`
	//Stateprov      string `json:"state_prov"`
	//District       string `json:"district"`
	//City           string `json:"city"`
	//Zipcode        string `json:"zipcode"`
	//Latitude       string `json:"latitude"`
	//Longitude      string `json:"longitude"`
	//Iseu           string `json:"is_eu"`
	//Callingcode    string `json:"calling_code"`
	//Countrytld     string `json:"country_tld"`
	//Languages      string `json:"languages"`
	//Countryflag    string `json:"country_flag"`
	//Geonameid      string `json:"geoname_id"`
	//Isp            string `json:"isp"`
	//Connectiontype string `json:"connection_type"`
	//Organization   string `json:"organization"`
	//Currency       string `json:"currency"`
	//	Code            string `json:"code"`
	//	Currencyname    string `json:"name"`
	//	Symbol          string `json:"symbol"`
	//Timezone string `json:"time_zone"`
	//TzName          string `json:"name"`
	//Offset          string `json:"offset"`
	//Currenttime     string `json:"current_time"`
	//Currenttimeunix string `json:"current_time_unix"`
	//Isdst           string `json:"is_dst"`
	//Dstsavings      string `json:"dst_savings"`
}

func main() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://api.ipgeolocation.io/ipgeo?apiKey=2caffbb7123943d7933b1fe729994186")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var responsestruct Response
	errj := json.Unmarshal(bodyBytes, &responsestruct)
	if errj != nil {
		fmt.Println(errj)
		return
	}
	fmt.Printf("API Response as struct %+v\n", responsestruct)
}
