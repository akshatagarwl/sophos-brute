package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var passwordList = []string{
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
	//

	err := logout("9919103015")
	if err != nil {
		fmt.Println(err)
	}
}

func login(username string, password string) error {
	url := "http://172.16.68.6:8090/login.xml"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("mode=191&username=%s&password=%s", username, password))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var some Requestresponse
	err = xml.Unmarshal(body, &some)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if !strings.Contains(some.Message, "You are signed in as {username}") {
		return fmt.Errorf("error logging in")
	}
	fmt.Println(some)

	return nil
}

func logout(username string) error {
	url := "http://172.16.68.6:8090/logout.xml"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("mode=193&username=%s", username))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var some Requestresponse
	err = xml.Unmarshal(body, &some)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if !strings.Contains(some.Message, "You&#39;ve signed out") {
		return fmt.Errorf("error logging out")
	}
	fmt.Println(some)

	return nil
}
