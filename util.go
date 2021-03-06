package independentreserveclient

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func getAndRead(URL string) ([]byte, error) {
	got, err := http.Get(URL)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(got.Body)
}

func getBody(request string) ([]byte, error) {
	resp, err := http.Get(request)
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (c Client) sign(URI, body string) (int64, string) {
	now := time.Now().Unix() * 1000 //milliseconds
	return now, c.hashEncode(URI + "\n" + strconv.FormatInt(now, 10) + "\n" + body)
}

func (c Client) hashEncode(message string) string {
	mac := hmac.New(sha512.New, c.decodedSecret)
	mac.Write([]byte(message))
	data := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return data
}

func (c Client) setupHeaders(req *http.Request, timestamp int64, signature string) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "UTF-8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.apikey)
	req.Header.Set("timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Set("signature", signature)
}

func (c Client) signAndPost(URI string, i interface{}) ([]byte, error) {
	return c.signAnd(URI, i, "POST")
}
func (c Client) signAndGet(URI string) ([]byte, error) {
	return c.signAnd(URI, nil, "GET")
}

func (c Client) sign(irr IndependentReserveRequest) (string, error) {
	f, err := irr.SignFormat()
	if err != nil {
		return "", err
	}
	return c.hashEncode(f)
}

func (c Client) hashEncode(message string) (string, error) {
	mac := hmac.New(sha512.New, []byte(c.private))
	mac.Write([]byte(message))

	data, err := hex.DecodeString(string(mac.Sum(nil)))
	if err != nil {
		return "", err
	}
	//1mac.Sum(nil)
	//data := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return string(data), nil
}

func (c Client) setupHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func (c Client) signAndPost(i IndependentReserveRequest) ([]byte, error) {
	return c.signAnd(i, "POST")
}
func (c Client) signAndGet(i IndependentReserveRequest) ([]byte, error) {
	return c.signAnd(i, "GET")
}

func (c Client) signAnd(i IndependentReserveRequest, do string) ([]byte, error) {
	var body []byte
	var err error
	if i != nil {
		body, err = json.Marshal(i)
		if err != nil {
			return nil, err
		}
	} else {
		body = []byte("")
	}

	//spew.Dump(i)
	client := http.Client{}
	signature, err := c.sign(i)
	if err != nil {
		return body, err
	}
	i.SetSignature(signature)
	req, err := http.NewRequest(do, i.GetURL(), bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("Error creating new Request;" + err.Error())
	}
	c.setupHeaders(req)
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Error doing request;" + err.Error())
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Error reading response;" + err.Error())
	}
	if response.StatusCode/100 != 2 {
		return nil, errors.New("StatusCode not 2xx; " + strconv.Itoa(response.StatusCode))
	}
	return body, err
}

func lookup(sl []string, cur string) int {
	for i, s := range sl {
		if s == cur {
			return i
		}
	}
	return -1
}
