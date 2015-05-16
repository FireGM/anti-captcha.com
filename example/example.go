// example
package main

import (
	"fmt"
	"log"

	"github.com/FireGM/anti-captcha.com/client"
)

func main() {
	antigate := client.GetClient("APIKEY") //получаем наш клиент для работы с API сайта anti-captcha.com.
	// APIKEY брать из настроек вашего аккаунта по ссылке https://anti-captcha.com/panel/settings/account
	captcha := "captcha.jpg"                          //что мы будем отправлять на сервер для разгадования.
	captcha_text, err := antigate.SendAndGet(captcha) //отправляем на сервер нашу каптчу и получаем сразу разгаданный текст или ошибку
	if err != nil {
		log.Fatal(err) //Если будет ошибка, то мы пишем лог с названием ошибки.
		//Там же будет написана ссылка, по которой можно прочитать, что эта ошибка означает, если это ошибка сайта anti-captcha.com
	}
	fmt.Println(captcha_text) //Выводим, что там разгадали.

	money, _ := antigate.GetBalanse() // проверяем баланс
	fmt.Println(money)

	captcha_id, err := antigate.UploadCaptcha(captcha) //Посылаем нашу капчу на разгадывание и получаем её ID в формате строки.
	if err != nil {
		log.Fatal(err) //Здесь делаем обработку ошибок, если вдруг появились онные
	}
	//Здесь проводим логику работы.
	//
	captcha_text, err := antigate.GetCaptchaText(captcha_id) //Здесь нам потребовалось получить текст капчи.
	if err != nil {
		log.Fatal(err) // Обработка ошибок
	}
	fmt.Println(captcha_text) //Выводим её или другую логику прикручиваем.
}
