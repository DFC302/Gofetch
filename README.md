# Gofetch
A simple script to take a list of URLs and fetch the status code for each URL

# Install
git clone https://github.com/DFC302/Gofetch.git \
go build gofetch.go

# Usage
Create a file, full of URLs, with each URL on a new line. 

gofetch -f [file] 

gofetch will then proceed down the list of URLs and return the status code for each url. 

gofetch -h
Usage of gofetch:
  -f string \
      Specify file to read URLs from. (default "file") 
  -file string \
      Specify file to read URLs from. (default "file") 
      
# Example
![Example](https://github.com/DFC302/Gofetch/blob/master/images/example.png)

# Author
Written by: Matthew Greer

# Version
Version 1.0

