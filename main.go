package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Parse command line flags
	var domain string
	var output string
	flag.StringVar(&domain, "d", "", "The target domain to collect js files from <proto:port>")
	flag.StringVar(&output, "o", "", "Output file name")
	flag.Parse()

	// Check if domain flag is provided
	if domain == "" {
		fmt.Println("Please provide the target domain using the -d flag")
		os.Exit(1)
	}

	// Parse domain URL
	u, err := url.Parse(domain)
	if err != nil {
		log.Fatalf("Error parsing domain: %v", err)
	}

	// Send HTTP request to domain URL
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Parse response HTML with goquery
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalf("Error parsing response HTML: %v", err)
	}

	// Collect script URLs
	var urls []string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		if src, ok := s.Attr("src"); ok {
			if strings.HasPrefix(src, "//") {
				// Handle URLs that start with double slashes
				src = fmt.Sprintf("%s:%s", u.Scheme, src)
			} else if strings.HasPrefix(src, "/") {
				// Handle absolute URLs
				src = fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, src)
			} else if !strings.HasPrefix(src, "http") {
				// Handle relative URLs
				src = fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, src)
			}
			if strings.HasSuffix(src, ".js") && strings.Contains(src, u.Host) {
				urls = append(urls, src)
			}
		}
	})

	// Filter out duplicate URLs
	urls = unique(urls)

	// Write URLs to output file
	if output != "" {
		file, err := os.Create(output)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer file.Close()
		for _, url := range urls {
			fmt.Fprintln(file, url)
		}
	} else {
		for _, url := range urls {
			fmt.Println(url)
		}
	}
}

// unique returns a slice of unique strings from the input slice
func unique(s []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
