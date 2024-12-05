package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
)



func loadEnvVars() (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", fmt.Errorf("error loading .env file: %v", err)
	}
	google := os.Getenv("GOOGLE")
	pass := os.Getenv("GMAILPASS")
	if google == "" || pass == "" {
		return "", "", fmt.Errorf("missing required environment variables")
	}
	return google, pass, nil
}

func createVerificationEmail(code string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: Arial, sans-serif;
			background-color: #f9f9f9;
			margin: 0;
			padding: 0;
		  }
		  .container {
			max-width: 600px;
			margin: 30px auto;
			background-color: #ffffff;
			padding: 20px;
			border-radius: 10px;
			box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		  }
		  h1 {
			color: #333;
			font-size: 22px;
		  }
		  p {
			color: #555;
			line-height: 1.6;
			font-size: 16px;
		  }
		  .code {
			margin: 20px 0;
			padding: 15px;
			background-color: #f4f4f4;
			border-left: 4px solid #007BFF;
			font-size: 20px;
			font-weight: bold;
			color: #333;
		  }
		  .footer {
			margin-top: 20px;
			font-size: 14px;
			color: #888;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <h1>Welcome to Tendanz!</h1>
		  <p>Thank you for joining us. Please use the verification code below to activate your account:</p>
		  <div class="code">%s</div>
		  <p>If you didn’t request this, please ignore this email.</p>
		  <div class="footer">
			&copy; 2024 Tendanz. All rights reserved.
		  </div>
		</div>
	  </body>
	</html>
	`, code)
}

func SendCode(target string, code string) error {
	google, pass, err := loadEnvVars()
	if err != nil {
		log.Printf("Error loading environment variables: %v", err)
		return err
	}

	htmlTemplate := createVerificationEmail(code)

	m := gomail.NewMessage()
	m.SetHeader("From", google)
	m.SetHeader("To", target)
	m.SetHeader("Subject", "Tendanz Verification Code")
	m.SetBody("text/html", htmlTemplate)

	d := gomail.NewDialer("smtp.gmail.com", 587, google, pass)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}

	log.Println("Verification email sent successfully!")
	return nil
}





func ContractPendingMail(target string, contractType string, client string) error {
	google, pass, err := loadEnvVars()
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", google)
	m.SetHeader("To", target)
	m.SetHeader("Subject", "New "+contractType+" Insurance Contract Created")
	m.SetBody("text/html", fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: Arial, sans-serif;
			background-color: #f9f9f9;
			margin: 0;
			padding: 0;
		  }
		  .container {
			max-width: 600px;
			margin: 30px auto;
			background-color: #ffffff;
			padding: 20px;
			border-radius: 10px;
			box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		  }
		  h1 {
			color: #333;
			font-size: 22px;
		  }
		  p {
			color: #555;
			line-height: 1.6;
			font-size: 16px;
		  }
		  .code {
			margin: 20px 0;
			padding: 15px;
			background-color: #f4f4f4;
			border-left: 4px solid #007BFF;
			font-size: 20px;
			font-weight: bold;
			color: #333;
		  }
		  .footer {
			margin-top: 20px;
			font-size: 14px;
			color: #888;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <h1>New `+contractType+` Insurance Contract Created</h1>
		  <p>A new `+contractType+` insurance contract has been created for client %s. Check the dashboard for more details</p>

		  <p>Thank you for using Tendanz.</p>
		  <div class="footer">
			&copy; 2024 Tendanz. All rights reserved.
		  </div>
		</div>
	  </body>
	</html>
	`, client))

	d := gomail.NewDialer("smtp.gmail.com", 587, google, pass)

	if err := d.DialAndSend(m); err != nil	{
		log.Println(err)
		return err
	}

	return nil
}



func ContractAccepted(target string, contractType string, client string) error {
	google, pass, err := loadEnvVars()
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", google)
	m.SetHeader("To", target)
	m.SetHeader("Subject", contractType+" Insurance Contract Accepted")
	m.SetBody("text/html", fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: Arial, sans-serif;
			background-color: #f9f9f9;
			margin: 0;
			padding: 0;
		  }
		  .container {
			max-width: 600px;
			margin: 30px auto;
			background-color: #ffffff;
			padding: 20px;
			border-radius: 10px;
			box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		  }
		  h1 {
			color: #333;
			font-size: 22px;
		  }
		  p {
			color: #555;
			line-height: 1.6;
			font-size: 16px;
		  }
		  .code {
			margin: 20px 0;
			padding: 15px;
			background-color: #f4f4f4;
			border-left: 4px solid #007BFF;
			font-size: 20px;
			font-weight: bold;
			color: #333;
		  }
		  .footer {
			margin-top: 20px;
			font-size: 14px;
			color: #888;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <h1>`+contractType+` Insurance Contract Accepted</h1>
		  <p>Your `+contractType+` insurance contract has been accepted. Check the dashboard for more details</p>

		  <p>Thank you for using Tendanz.</p>
		  <div class="footer">
			&copy; 2024 Tendanz. All rights reserved.
		  </div>
		</div>
	  </body>
	</html>
	`))

	d := gomail.NewDialer("smtp.gmail.com", 587, google, pass)

	if err := d.DialAndSend(m); err != nil	{
		log.Println(err)
		return err
	}	
		
	return nil
}


func ContractRejected(target string, contractType string, client string) error {

	google, pass, err := loadEnvVars()
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", google)
	m.SetHeader("To", target)
	m.SetHeader("Subject", contractType+" Insurance Contract Rejected")
	m.SetBody("text/html", fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: Arial, sans-serif;
			background-color: #f9f9f9;
			margin: 0;
			padding: 0;
		  }
		  .container {
			max-width: 600px;
			margin: 30px auto;
			background-color: #ffffff;
			padding: 20px;
			border-radius: 10px;
			box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		  }
		  h1 {
			color: #333;
			font-size: 22px;
		  }
		  p {
			color: #555;
			line-height: 1.6;
			font-size: 16px;
		  }
		  .code {
			margin: 20px 0;
			padding: 15px;
			background-color: #f4f4f4;
			border-left: 4px solid #007BFF;
			font-size: 20px;
			font-weight: bold;
			color: #333;
		  }
		  .footer {
			margin-top: 20px;
			font-size: 14px;
			color: #888;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <h1>`+contractType+` Insurance Contract Rejected</h1>
		  <p>Your `+contractType+` insurance contract has been rejected. Check the dashboard for more details</p>

		  <p>Thank you for using Tendanz.</p>
		  <div class="footer">
			&copy; 2024 Tendanz. All rights reserved.
		  </div>
		</div>
	  </body>
	</html>
	`))

	d := gomail.NewDialer("smtp.gmail.com", 587, google, pass)

	

	if err := d.DialAndSend(m); err != nil	{
		log.Println(err)
		return err
	}

	return nil
}
