package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printColor (color string, logLevel string) {
	const (
		Reset  = "\033[0m"
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Blue   = "\033[34m"
		Magenta = "\033[35m"
		Cyan   = "\033[36m"
		Gray   = "\033[37m"
		White  = "\033[97m" // Bright white 
	)

	timestamp := time.Now()

	// Print the random string
	switch logLevel {
		case "CRITICAL":
 		case "SEVERE":
			fmt.Println(timestamp, Red + "CRITICAL " + "Severe Error" + Reset )
 		case "DEBUG":
			fmt.Println(timestamp, Green + "DEBUG " + "This is a Debug line" + Reset )
 		case "EMERGENCY":
			fmt.Println(timestamp, Cyan + "EMERGENCY " + "Emergency Maintenance Notification" + Reset)
 		case "ERROR":
			fmt.Println(timestamp, Magenta + "ERROR "+ "An error occured " + Reset)
 		case "FATAL":
			fmt.Println(timestamp, Red + "FATAL " + "Fatal error. Check infra" + Reset )
 		case "INFO":
			fmt.Println(timestamp, Cyan + "INFO " + "This is for information only" + Reset )
 		case "TRACE":
			fmt.Println(timestamp, White + "TRACE " + "Tracelog: Check logs ... " + Reset )
 		case "WARN":
			fmt.Println(timestamp, Blue + "WARNING " + "Warning: Too many comments " + Reset )
 		case "ALERT":
			fmt.Println(timestamp, Gray + "ALERT " + "Alerting on PagerDuty " + Reset  )
		}
}

func main() {
	colorArray := []string{"red", "green", "yellow", "blue", "magenta", "cyan", "gray", "white"}
	logLevels := []string{"CRITICAL", "DEBUG", "EMERGENCY", "ERROR", "FATAL", "INFO", "SEVERE", "TRACE", "WARN", "ALERT"}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a ticker that triggers every 5 seconds
	ticker := time.NewTicker(5 * time.Second)

	// Run the loop indefinitely
	for range ticker.C {
		// Generate a random color & Log Level
		randomColor := rand.Intn(len(colorArray))
		randomLevel := rand.Intn(len(logLevels))
		printColor (colorArray[randomColor], logLevels[randomLevel])
	}
}
