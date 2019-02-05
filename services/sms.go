package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var SendSMS = func() {
	// Set account keys & information
	accountSid := "AC0ac28a1fef741ca6a50387754dec3206"
	authToken := "4e3531b6dae646e4232957924664757a"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Create possible message bodies
	quotes := [7]string{"I urge you to please notice when you are happy, and exclaim or murmur or think at some point, 'If this isn't nice, I don't know what is.'",
		"Peculiar travel suggestions are dancing lessons from God.",
		"There's only one rule that I know of, babies—God damn it, you've got to be kind.",
		"Many people need desperately to receive this message: 'I feel and think much as you do, care about many of the things you care about, although most people do not care about them. You are not alone.'",
		"That is my principal objection to life, I think: It's too easy, when alive, to make perfectly horrible mistakes.",
		"So it goes.",
		"We must be careful about what we pretend to be."}

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", "4028922865")
	msgData.Set("From", "4023477087")
	msgData.Set("Body", quotes[rand.Intn(len(quotes))])
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
