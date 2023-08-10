package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func main() {
	args := os.Args

	// "Uso: ./sendmail archivo.zip id_job nombre_job desinatario"

	pathToZipFile := args[1]
	buildID := args[2]
	jobName := args[3]
	toMessage := args[4]

	if pathToZipFile == "" || buildID == "" || jobName == "" || toMessage == "" {
		log.Fatal("Faltan argumentos, estos son obligatorios")
	}

	// email
	m := gomail.NewMessage()

	m.SetHeader("From", "servidorjenkins@example.com")

	m.SetHeader("To", toMessage)

	subjectMessage := fmt.Sprintf("Resultados del job: %s", jobName)
	m.SetHeader("Subject", subjectMessage)

	message := fmt.Sprintf("<h1>Este es un mail generado por el servidor de Jenkins</h1><h2>ID: %s<h2><h2>Resultados del job: %s</h2>", buildID, jobName)
	m.SetBody("text/html", message)

	m.Attach(pathToZipFile)

	// Send email function
	d := gomail.NewDialer("", 587, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}

}
