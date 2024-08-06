package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func main() {
    // Instantiate the collector
    c := colly.NewCollector()

    // Create a struct to store product data
    type Product struct {
        Title    string
        Subtitle string
        Price    string
        URL      string
        ImageURL string
    }

    // Slice to store all products
    var products []Product

    // On every element with the product-card class
    c.OnHTML(".product-card", func(e *colly.HTMLElement) {
        product := Product{
            Title:    e.ChildText(".product-card__title"),
            Subtitle: e.ChildText(".product-card__subtitle"),
            Price:    e.ChildText(".product-card__price"),
            URL:      e.ChildAttr(".product-card__link-overlay", "href"),
            ImageURL: e.ChildAttr(".product-card__hero-image", "src"),
        }
        products = append(products, product)
    })

    // Start scraping the page
    c.Visit("https://www.nike.com/ca/fr/w/hommes-vetements-6ymx6znik1")

	file,error:=os.Create("output.csv")
	if error!=nil{
		log.Fatalf("Failed to create the file")
	}

	defer file.Close()

	// Create a CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()

	header:=[]string{"Title", "Subtitle", "Price", "URL", "Image URL"}
	if err := writer.Write(header); err != nil {
        log.Fatalf("Failed to write header to CSV: %v", err)
    }

    // Print all products
    for _, product := range products {
        record := []string{product.Title, product.Subtitle, product.Price, product.URL, product.ImageURL}
        if err := writer.Write(record); err != nil {
            log.Fatalf("Failed to write record to CSV: %v", err)
        }
    }
}
