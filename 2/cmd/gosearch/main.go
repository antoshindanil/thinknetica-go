package main

import (
	"flag"
	"fmt"
	"go-search/pkg/crawler"
	"go-search/pkg/crawler/spider"
	"strings"
)

func main() {
	service := spider.New()

	searchedValue := flag.String("s", "", "Строка для поиска")
	flag.Parse()

	if *searchedValue == "" {
		fmt.Println("Укажите строку для поиска с флагом -s")
		return
	}

	urls := [2]string{"https://golang.com", "https://go.dev"}

	scanned := []crawler.Document{}

	for _, url := range urls {
		data, _ := service.Scan(url, 2)
		scanned = append(scanned, data...)
	}

	filteredData := []crawler.Document{}

	for _, v := range scanned {
		if strings.Contains(v.Title, *searchedValue) {
			filteredData = append(filteredData, v)
		}
	}

	if len(filteredData) == 0 {
		fmt.Println("Ничего не найдено")
		return
	}

	for _, v := range filteredData {
		fmt.Printf("Заголовок: %s\nURL: %s\n\n", v.Title, v.URL)
	}

}
