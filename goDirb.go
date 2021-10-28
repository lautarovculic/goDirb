package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	//"strings"
	//"sync"
)

func scan() {

	//You must put HTTPS: // or HTTP: // without / at the end. Example: https://lautarovculic.com
	fmt.Print("URL: ")
	var firstUrl string
	fmt.Scan(&firstUrl)

	//Full path where the file is located without quotation marks. Example: C:\Users\lauta\Desktop\wordlist.txt
	fmt.Print("File path: ")
	var firstPath string
	fmt.Scan(&firstPath)

	readFile, err := os.Open(firstPath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()
	/*
		//Ask
		go func(msg string) {
			print("Do you want to do another enumeration? [y/n]: ")
			var answer string
			fmt.Scan(&answer)

			switch strings.ToLower(answer) {
			case "y", "yes":
				go scan()
			case "n", "no":
				break
			default:
				fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter: ")

			}
		}("Other enumeration...")
	*/

	for _, eachline := range fileTextLines {

		resp, err := http.Get(firstUrl + "/" + eachline)
		if err != nil {
			fmt.Println("Must be with 'HTTP://' or 'HTTPS://' the URL.")
			log.Fatalln(err)
		}

		//Leave the status code you want to find.
		if resp.StatusCode == 200 || resp.StatusCode == 204 || resp.StatusCode == 301 || resp.StatusCode == 302 || resp.StatusCode == 307 || resp.StatusCode == 401 || resp.StatusCode == 403 {
			fmt.Println("["+firstUrl+"/"+eachline+"] "+"HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		}

	}

}

func main() {

	println("|| Lautaro Villarreal Culic' ||")
	println("|| https://lautarovculic.com ||")
	println("|||||||| goDirb - v0.2 ||||||||")
	println("")
	println("")

	scan()

}
