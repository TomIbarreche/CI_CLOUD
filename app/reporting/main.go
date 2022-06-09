package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	teams "github.com/atc0005/go-teams-notify/v2"
	_ "github.com/go-sql-driver/mysql"
)

func getProductCount(uri string) (int, error) {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	count := 0
	err = db.QueryRow("SELECT COUNT(*) FROM products;").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func sendMessageToTeams(webhookURL string, productCount int) error {
	// setup message card
	msgCard := teams.NewMessageCard()
	msgCard.Title = "Daily reporting"
	msgCard.Text = fmt.Sprintf("%d products in database\n", productCount)
	msgCard.ThemeColor = "#DF813D"

	// send
	mstClient := teams.NewClient()
	mstClient.SkipWebhookURLValidationOnSend(true)

	return mstClient.Send(webhookURL, msgCard)
}

func main() {
	os.Setenv("WEBHOOK_URL", "https://outlook.office.com/webhookb2/6d742bc5-60f4-4f1d-ae3a-c3ad31391d4f@901cb4ca-b862-4029-9306-e5cd0f6d9f86/IncomingWebhook/f8b6dfdf1dc7460ca74a4076b400913b/85354c20-2276-43f6-ab80-15a8b176d0df")
	os.Setenv("DB_URI", "root:TVlTUUxAQWRtaW4wNDUx@tcp(mysql:3306)/laravel")
	dbURI := os.Getenv("DB_URI")
	webhookURL := os.Getenv("WEBHOOK_URL")

	if len(dbURI) == 0 {
		log.Fatal("Missing DB_URI environment variable")
	}
	if len(webhookURL) == 0 {
		log.Fatal("Missing WEBHOOK_URL environment variable")
	}

	count, err := getProductCount(dbURI)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(count)

	err = sendMessageToTeams(webhookURL, count)
	if err != nil {
		log.Fatal(err.Error())
	}
}
