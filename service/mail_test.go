package service

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"testing"
)

func TestSendMail(t *testing.T) {
	from := mail.NewEmail("唐新尧", "shaytang@163.com")
	//"smtp.sendgrid.net"
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("唐新尧", "2578873659@qq.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	//key := "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"
	key := "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"
	client := sendgrid.NewSendClient(key)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
