package services

import (
	"cdc_mailer/models"
	"cdc_mailer/utils"
	"context"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
	"sync"
)

func scrapeWebPage(url string, dataHolder *string, errorHolder *error, attempts int, waitGrp *sync.WaitGroup) {
	scraper := colly.NewCollector()

	var response = true
	var retryCount = 0

	scraper.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})

	scraper.OnError(func(request *colly.Response, err error) {
		fmt.Println("Error Occurred....Re-Visiting", url)
		response = false
		retryCount += 1
	})

	scraper.OnHTML("body", func(styleElement *colly.HTMLElement) {
		bodyString := styleElement.Text
		*dataHolder = utils.CompanyDetailsParser(bodyString)
		*errorHolder = nil
		defer waitGrp.Done()
		return
	})

	for retryCount <= attempts {
		_ = scraper.Visit(url)
		if response == true {
			break
		}
	}

	if retryCount > attempts {
		*dataHolder = ""
		*errorHolder = errors.New("web page not scraped")
		return
	}
}

func getLLMResponse(companyProfile string, templateHolder *string) {
	// Clean the Data and get company profile.
	llmContext := context.Background()
	client, llmError := genai.NewClient(llmContext, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if llmError != nil {
		log.Fatal("[ERROR] connection to LLM was not established")
	}
	defer client.Close()

	llmModel := client.GenerativeModel("gemini-1.0-pro")

	promptParts := utils.GeneratePrompt(companyProfile)

	response, _ := llmModel.GenerateContent(llmContext, promptParts...)

	for _, candidate := range response.Candidates {
		for _, part := range candidate.Content.Parts {
			*templateHolder += fmt.Sprintf("%s", part)
		}
	}
}

func GenerateTemplates(companyData *models.Company) error {
	// Get URLs
	companyCareers := companyData.CompanyCareers
	companyAbout := companyData.CompanyAbout

	// Initialize Wait Group and scrape
	waitGrp := new(sync.WaitGroup)
	waitGrp.Add(2)

	var scrapedCompanyCareers, scrapedCompanyAbout string
	var scrapeCareerError, scrapeAboutError error

	go scrapeWebPage(companyCareers, &scrapedCompanyCareers, &scrapeCareerError, 3, waitGrp)
	go scrapeWebPage(companyAbout, &scrapedCompanyAbout, &scrapeAboutError, 3, waitGrp)
	waitGrp.Wait()

	// Log errors
	if scrapeCareerError != nil {
		fmt.Println("Scrape failed for URL ", companyCareers, "Error: ", scrapeCareerError)
	}
	if scrapeAboutError != nil {
		fmt.Println("Scrape failed for URL ", companyAbout, "Error: ", scrapeAboutError)
	}

	_ = os.WriteFile("car.txt", []byte(scrapedCompanyCareers), 0666)
	_ = os.WriteFile("abt.txt", []byte(scrapedCompanyAbout), 0666)

	var mailTemplate string

	companyProfile := "Company About Us Data \n" + scrapedCompanyAbout + "\nCompany Careers Data\n" + scrapedCompanyCareers

	getLLMResponse(companyProfile, &mailTemplate)

	fmt.Println(mailTemplate)

	return nil

}
