package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type URLRequest struct {
    URL       string `json:"url"`
    Operation string `json:"operation"`
}

type URLResponse struct {
    ProcessedURL string `json:"processed_url"`
}

func processURL(url, operation string) string {
    log.Printf("Processing URL: %s with operation: %s", url, operation) 
    switch operation {
    case "canonical":
        return canonicalURL(url)
    case "redirection":
        return redirectionURL(url)
    case "all":
        url = canonicalURL(url)
        return redirectionURL(url)
    default:
        return url
    }
}

func canonicalURL(url string) string {
    if idx := strings.Index(url, "?"); idx != -1 {
        url = url[:idx]
    }
    return strings.TrimRight(url, "/")
}

func redirectionURL(inputURL string) string {
    inputURL = strings.ToLower(inputURL)

    parsedURL, err := url.Parse(inputURL)
    if err != nil {
        log.Printf("Error parsing URL: %v", err) 
        return ""
    }

    parsedURL.Scheme = "https"
    parsedURL.Host = "www.byfood.com"
    parsedURL.RawQuery = ""
    parsedURL.Fragment = ""

    resultURL := parsedURL.String()
    log.Printf("Redirection URL result: %s", resultURL)
    return resultURL
}



// UrlHandler processes a URL based on the provided operation
// @Summary Process a URL
// @Description Processes a URL based on the operation specified in the request
// @Tags URL Processing
// @Accept json
// @Produce json
// @Param request body URLRequest true "URL Request"
// @Success 200 {object} URLResponse "URL successfully processed"
// @Failure 400 {string} string "Invalid request"
// @Failure 405 {string} string "Only POST method is allowed"
// @Router /process-url [post]
func UrlHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received request: %s", r.URL.Path) 
    if r.Method != "POST" {
        log.Printf("Invalid request method: %s", r.Method) 
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var request URLRequest
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        log.Printf("Error decoding request: %v", err) 
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    processedURL := processURL(request.URL, request.Operation)
    response := URLResponse{ProcessedURL: processedURL}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)

    log.Printf("Processed and responded with URL: %s", processedURL) // Log the response
}