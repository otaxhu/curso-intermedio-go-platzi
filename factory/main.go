package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS

type SMSNotification struct {
}

func (s SMSNotification) SendNotification() {
	fmt.Println("Sending Notification SMS")
}

func (s SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (s SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (s SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Email

type EmailNotification struct{}

func (e EmailNotification) SendNotification() {
	fmt.Println("Sending Notification Email")
}

func (e EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

func (e EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (e EmailNotificationSender) GetSenderChannel() string {
	return "Outlook"
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("error:\nNo notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func getChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := GetNotificationFactory("SMS")
	emailFactory, _ := GetNotificationFactory("Email")
	sendNotification(smsFactory)
	sendNotification(emailFactory)
	getMethod(smsFactory)
	getMethod(emailFactory)
	getChannel(smsFactory)
	getChannel(emailFactory)
}
