package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const uri = "http://anti-captcha.com/in.php"

type Client struct {
	Key string
}

func (c *Client) UploadCaptcha(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(part, file)

	err = writer.WriteField("key", c.Key)
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.DefaultClient.Post(uri, writer.FormDataContentType(), body)
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	captcha_id, err := GetErrorOrResponse(string(res))
	if err != nil {
		return "", err
	}
	return captcha_id, nil
}

func (c *Client) UploadCaptchaByURL(url string) (string, error) {
	path, err := saveFromUrl(url)
	if err != nil {
		return "", err
	}
	return c.UploadCaptcha(path)
}

func (c *Client) GetCaptchaText(id string) (string, error) {
	url_response := "http://anti-captcha.com/res.php"
	values := url.Values{}
	values.Add("key", c.Key)
	values.Add("action", "get")
	values.Add("id", id)
	for {
		resp, _ := http.DefaultClient.PostForm(url_response, values)
		defer resp.Body.Close()
		res, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		captcha_text, err := GetErrorOrResponse(string(res))
		if err != nil {
			if string(res) == "CAPCHA_NOT_READY" {
				time.Sleep(time.Second * 3)
				continue
			}
			return "", err
		} else {
			return captcha_text, nil
		}

	}
}

func GetErrorOrResponse(res string) (string, error) {
	splited := strings.Split(res, "|")
	if splited[0] == "OK" {
		captcha_id := splited[1]
		return captcha_id, nil
	} else {
		err := CaptchaError{What: res}
		return "", err
	}
}

func (c *Client) SendAndGet(path string) (string, error) {
	response, err := c.UploadCaptcha(path)
	if err != nil {
		return "", err
	}
	captcha_text, err := c.GetCaptchaText(response)
	if err != nil {
		return "", err
	}
	return captcha_text, nil
}

func (c *Client) SendAndGetByURL(url string) (string, error) {
	path, err := saveFromUrl(url)
	if err != nil {
		return "", err
	}
	return c.SendAndGet(path)
}

func GetClient(key string) *Client {
	return &Client{Key: key}
}

func (c *Client) GetBalanse() (float64, error) {
	values := url.Values{}
	values.Add("key", c.Key)
	values.Add("action", "getbalance")
	u := url.URL{}
	u.Host = "anti-captcha.com"
	u.Scheme = "http"
	u.Path = "/res.php"
	u.RawQuery = values.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	b, err := ioutil.ReadAll(response.Body)
	money, _ := strconv.ParseFloat(string(b), 10)
	return money, nil
}
