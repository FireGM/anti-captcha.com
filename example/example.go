// example
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"io/ioutil"

	"github.com/FireGM/anti-captcha.com/client"
)

func main() {
	conf := getConf()                         //получаем настройки приложения из файла conf.json
	antigate := client.GetClient(conf.ApiKey) //получаем наш клиент для работы с API сайта anti-captcha.com.
	// APIKEY брать из настроек вашего аккаунта по ссылке https://anti-captcha.com/panel/settings/account
	captcha := "captcha.jpg"                         //что мы будем отправлять на сервер для разгадования.
	captchaText, err := antigate.SendAndGet(captcha) //отправляем на сервер нашу каптчу и получаем сразу разгаданный текст или ошибку
	if err != nil {
		log.Fatal(err) //Если будет ошибка, то мы пишем лог с названием ошибки.
		//Там же будет написана ссылка, по которой можно прочитать, что эта ошибка означает, если это ошибка сайта anti-captcha.com
	}
	fmt.Println(captchaText) //Выводим, что там разгадали.

	money, _ := antigate.GetBalanse() // проверяем баланс
	fmt.Println(money)

	captchaId, err := antigate.UploadCaptcha(captcha) //Посылаем нашу капчу на разгадывание и получаем её ID в формате строки.
	if err != nil {
		log.Fatal(err) //Здесь делаем обработку ошибок, если вдруг появились онные
	}
	//Здесь проводим логику работы.
	//
	captchaText, err = antigate.GetCaptchaText(captchaId) //Здесь нам потребовалось получить текст капчи.
	if err != nil {
		log.Fatal(err) // Обработка ошибок
	}
	fmt.Println(captchaText) //Выводим её или другую логику прикручиваем.

	//отправляем каптчу с ссылки и получаем разгаданную
	captchaText, err = antigate.SendAndGetByURL("https://upload.wikimedia.org/wikipedia/commons/c/c7/Captcha_voorbeeld.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(captchaText)
}

type Conf struct {
	ApiKey string `json:"api_key"`
}

func getConf() Conf {
	file, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}
	var c Conf
	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}
	return c
}
