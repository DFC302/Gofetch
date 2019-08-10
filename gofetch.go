// Gofetch

package main

import (
		"fmt"
		"bufio"
		"flag"
		"os"
		"log"	
		"net/http"
		"strings"
)

func main() {

	// Set file flag, used to specify file to be opened by OS
	var file string

	flag.StringVar(&file, "file", "file", "Specify file to read URLs from.")
	flag.StringVar(&file, "f", "file", "Specify file to read URLs from.")

	flag.Parse()

	// Print title info at begining

	fmt.Println("===================================================")
	fmt.Println("Gofetch")
	fmt.Println("Version 1.0")
	fmt.Println("===================================================")
	fmt.Println("Written by: Matthew Greer")
	fmt.Println("Github: https://github.com/DFC302")
	fmt.Println("Twitter: https://twitter.com/Vail_302")
    fmt.Println("===================================================")
    fmt.Println("Scanning URL's from file:", file)
    fmt.Println("===================================================")
    fmt.Println()

    // Open file, handle errors
	data, err := os.Open(file)

	if err != nil{
		log.Fatal(err)
	}

	// Close file
	defer data.Close()

	// Read in data
	scanner := bufio.NewScanner(data)
    

    for scanner.Scan() {
        line := scanner.Text()
       	
       	// If url has http:// or https://, leave it alone
        if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
        	line = line
        } else if line == "" { // if line is blank, pass it
            continue
        } else { // else, assume https, append it to begining of url
        	https := "https://"
        	line = (https + line) // assume https
        }

        // try to make a request for each line

        resp, err := http.Get(line)

        // if error should occur, print it, no fatal, continue with requests
        if err != nil{
            fmt.Println(err)
            fmt.Println()
        	continue
        }

        //fmt.Println(line, resp.StatusCode)
        fmt.Println("URL:\t\t", line)
        fmt.Println("Status Code:\t", resp.StatusCode)
        fmt.Println()
    }

    // handle scan errors
	 if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    os.Exit(0)
}
