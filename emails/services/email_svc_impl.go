package services

import (
	"fmt"
	"net/http"
	"nganterin-go/emails/dto"
	"nganterin-go/pkg/exceptions"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type CompServicesImpl struct{}

func NewComponentServices() CompServices {
	return &CompServicesImpl{}
}

func (s *CompServicesImpl) SendEmail(data dto.EmailRequest) *exceptions.Exception {
	email := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	server := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	i, err := strconv.Atoi(smtpPort)
	if err != nil {
		return exceptions.NewException(http.StatusInternalServerError, exceptions.ErrInternalServer)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", data.Email)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	d := gomail.NewDialer(server, i, email, password)

	if err := d.DialAndSend(m); err != nil {
		return exceptions.NewException(http.StatusBadGateway, exceptions.ErrEmailSendFailed)
	}

	return nil
}

func (s *CompServicesImpl) VerificationEmail(data dto.EmailVerification) *exceptions.Exception {
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>Email Verification</title>

			<style>
			@import url("https://fonts.googleapis.com/css2?family=Poppins:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;0,900;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800;1,900&display=swap");

			body {
				font-family: "Poppins", sans-serif;
				display: flex;
				flex-direction: column;
				justify-content: center;
				align-items: center;
				min-height: 100vh;
				margin: 0;
				padding: 0;
				background-color: #f6f6f6;
			}
			.container {
				max-width: 600px;
				margin: 0 auto;
				padding: 30px;
				border-radius: 20px;
				background-color: #ffffff;
				box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
			}
			.header {
				text-align: center;
				margin-bottom: 20px;
			}
			.header img {
				max-width: 100px;
				margin-bottom: 15px;
			}
			.title {
				font-size: 28px;
				font-weight: 600;
				color: #171717;
				margin-bottom: 10px;
			}
			.message {
				font-size: 15px;
				line-height: 1.6;
				color: #171717;
				margin-bottom: 20px;
				text-align: center;
			}
			.verify-button {
				text-align: center;
				margin: 0 auto;
			}
			.verify-button a {
				display: inline-block;
				text-align: center;
				background: linear-gradient(90deg, #38bdf8, #0284c7);
				color: #ffffff;
				text-decoration: none;
				font-size: 15px;
				font-weight: bold;
				padding: 10px;
				border-radius: 10px;
				box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
				margin: 20px auto;
			}
			.verify-button a:hover {
				background: linear-gradient(90deg, #0ea5e9, #075985);
			}
			</style>
		</head>
		<body>
			<div class="container">
			<div class="header">
				<p class="title">Verify Your Email Address</p>
			</div>
			<p class="message" style="font-weight: 600">Dear %s,</p>
			<p class="message">
				Thank you for registering with our platform. To complete your
				registration, please click the button below:
			</p>
			<div class="verify-button">
				<a href="%s">Verify My Email</a>
			</div>
			<p
				class="message"
				style="font-size: 14px; color: rgb(255, 85, 85); margin-top: -15px"
			>
				This link will expire in 2 hours. If you did not request this
				verification, please ignore this email.
			</p>
			</div>
			<p class="footer" style="font-size: 14px">
			© 2024 Nganterin. All Rights Reserved.,<br />Best regards, Nganterin
			</p>
		</body>
		</html>

	`, data.Email, data.VerificationURL)

	emailData := dto.EmailRequest{
		Email:   data.Email,
		Subject: data.Subject,
		Body:    body,
	}

	err := s.SendEmail(emailData)
	if err != nil {
		return err
	}

	return nil
}
