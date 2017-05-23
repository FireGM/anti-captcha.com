# anti-captcha.com
Приложение для работы с сайтом anti-captcha.com на GOlang

## Как использовать:
Скачиваем данный пакет
```
go get github.com/FireGM/anti-captcha.com
```
Теперь мы можем создать наше приложение
```go
package main

import (
	"fmt"
	"github.com/FireGM/anti-captcha.com/client"
)

func main() {
	  antigate := client.GetClient("APIKEY") //получаем наш клиент для работы с API сайта anti-captcha.com.
    // APIKEY брать из настроек вашего аккаунта по ссылке https://anti-captcha.com/clients/settings/apisetup
    captcha := "captcha.jpg" //что мы будем отправлять на сервер для разгадования.
    captchaText, err := antigate.SendAndGet(captcha) //отправляем на сервер нашу каптчу и получаем сразу разгаданный текст или ошибку
    if err != nil {
        log.Fatal(err) //Если будет ошибка, то мы пишем лог с названием ошибки.
        //Там же будет написана ссылка, по которой можно прочитать, что эта ошибка означает, если это ошибка сайта anti-captcha.com
    }
    fmt.Println(captchaText) //Выводим, что там разгадали.
}
```
###Проверяем баланс
```go 
	money, _ := antigate.GetBalanse() // проверяем баланс
	fmt.Println(money)
```

###Больше власти
Если вам нужно более точечная обработка, то можно посылать капчу на разгадывание, получать её ID, а текст запросить, когда он вам потребуется

Для начала мы загрузим и получим ID капчи. Можем её сохранить куда вам угодно.

```go
	captchaId, err := antigate.UploadCaptcha(captcha) //Посылаем нашу капчу на разгадывание и получаем её ID в формате строки.
	if err != nil {
		log.Fatal(err) //Здесь делаем обработку ошибок, если вдруг появились онные
	}
	fmt.Println(captchaId) //Печатаем ID капчи
```

Мы ещё не запрашивали текст. Мы только послали купчу на разгадывание.
Теперь пришло время её получить, а для этого нам потребается ID капчи, который мы получили выше.

```go
	captchaText, err := antigate.GetCaptchaText(captcha_id) //Здесь нам потребовалось получить текст капчи.
	if err != nil {
		log.Fatal(err) // Обработка ошибок
	}
	fmt.Println(captchaText) //Выводим её или другую логику прикручиваем.
```
Вуаля, мы получили текст.


отправка капчи по урл

```go
	captchaText, err = antigate.SendAndGetByURL("https://upload.wikimedia.org/wikipedia/commons/c/c7/Captcha_voorbeeld.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(captchaText)
```
или получения ID капчи до востребования

```go
	captchaId, err := antigate.UploadCaptchaByURL(captcha) //Скачиваем и посылаем нашу капчу на разгадывание и получаем её ID в формате строки.
	if err != nil {
		log.Fatal(err) //Здесь делаем обработку ошибок, если вдруг появились онные
	}
	fmt.Println(captchaId) //Печатаем ID капчи
```

###Пример
В папке example находится рабочий пример приложения. Для его запуска создайте в папке файл conf.json с данными в виде 

```json
{"api_key": "ВАШКЛЮЧДЛЯАПИ"}
```
Ключ брать по ссылке https://anti-captcha.com/clients/settings/apisetup