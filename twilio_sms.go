package golang_helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type PhoneInfo struct {
	Valid     string
	Number164 string
	Local     string
	Country   string
	Carrier   string
	Type      string
	Cost      string
}

type Lookup struct {
	CallerName  string `json:"caller_name"`
	PhoneNumber string `json:"phone_number"`
	Carrier     *struct {
		ErrorCode         *int   `json:"error_code"`
		MobileCountryCode string `json:"mobile_country_code"`
		MobileNetworkCode string `json:"mobile_network_code"`
		Name              string `json:"name"`
		Type              string `json:"type"`
	} `json:"carrier"`
	CountryCode    string `json:"country_code"`
	NationalFormat string `json:"national_format"`
	AddOns         string `json:"add_ons"`
	URL            string `json:"url"`
}

func SendSMS(send_to, content string) int {

	sms_sent := 0
	if len(content) > 160 {
		content = content[:160]
	}

	twilio_phone := "+17899876543" //fake US Phone Number

	accountSid := "get_from_twillio_account"
	authToken := "get_from_twillio_account"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", send_to)
	v.Set("From", twilio_phone)
	v.Set("Body", content)
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make request
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)

		if err == nil {
			sms_sent = 1
			log.Println(data)
		} else {
			log.Println(err)
		}
	}

	return sms_sent
}
