package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Debug(format string, args ...interface{}) {
	format = time.Now().Format("2006-01-02 15:04:05 ") + format + "\n"
	fmt.Fprintf(os.Stdout, format, args...)
}

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
                        Debug( "%s %s\n", timestamp, Red + "CRITICAL " + "Critical Error" + Reset )
                case "SEVERE":
                        Debug( "%s %s\n", timestamp, Red + "SEVERE " + "Severe Error" + Reset )
                case "DEBUG":
                        Debug( "%s %s\n", timestamp, Green + "DEBUG " + "This is a Debug line" + Reset )
                case "EMERGENCY":
                        Debug( "%s %s\n", timestamp, Cyan + "EMERGENCY " + "Emergency Maintenance Notification" + Reset )
                case "ERROR":
                        Debug( "%s %s\n", timestamp, Magenta + "ERROR "+ "An error occured " + Reset )
                case "FATAL":
                        Debug( "%s %s\n", timestamp, Red + "FATAL " + "Fatal error. Check infra" + Reset )
                case "INFO":
                        Debug( "%s %s\n", timestamp, Cyan + "INFO " + "This is for information only" + Reset )
                case "TRACE":
                        Debug( "%s %s\n", timestamp, White + "TRACE " + "Tracelog: Check logs ... " + Reset )
                case "WARN":
                        Debug( "%s %s\n", timestamp, Blue + "WARNING " + "Warning: Too many comments " + Reset )
                case "ALERT":
                        Debug( "%s %s\n", timestamp, Gray + "ALERT " + "Alerting on PagerDuty " + Reset )
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
