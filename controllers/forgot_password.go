package controllers

import (
	"TestCodelite/api/responses"
	"TestCodelite/auth"
	"TestCodelite/models"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
)




func (server *Server) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var p models.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := models.User{}
	userReceived, err := user.FindUserByEmail(server.DB, p.Email)
	token,_ :=auth.CreateToken(userReceived.ID)
	SendInterfaces(userReceived.Email,token)

	responses.JSON(w, http.StatusOK,nil)
}

func SendInterfaces(email string,token string) {
	to := email
	cc := []string{}
	subject := "Lupa Password"
	url := "Token : "+ token
	message := "To Reset the password follow this link " + url

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent!")
}

func sendMail(recipientEmail string, cc []string, subject, message string) error {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	var CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
	var CONFIG_SMTP_PORT,_ = strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
	var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

	to := []string{recipientEmail}
	msg := []byte("To: "+recipientEmail+"\r\n" +
		"Subject: "+subject+"\r\n" +
		"\r\n" +
		message+"\r\n")

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST,  CONFIG_SMTP_PORT)

	err = smtp.SendMail(smtpAddr, auth, CONFIG_SENDER_NAME, to, msg)
	if err != nil {
		return err
	}

	return nil

}