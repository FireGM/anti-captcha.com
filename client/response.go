package client

import "fmt"

const url_error = "https://anti-captcha.com/apidoc"

type CaptchaError struct {
	What string
}

func (c CaptchaError) Error() string {
	return fmt.Sprintf("\nERROR! %s. \nLook description on %s", c.What, url_error)
}

func (c CaptchaError) GetCode() string {
	return c.What
}
