package main

import (
	"bufio"
	"fmt"
	"go-image-downloader/downloader"
	"log"
	"os"
	"sync"
)

const (
	numWorkers = 5
	folder     = "images"
)

func readUrls(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			urls = append(urls, url)
		}
	}
	return urls, scanner.Err()
}

func main() {
	urls, err := readUrls("urls.txt")
	if err != nil {
		fmt.Println("Error reading URLs:", err)
		return
	}

	var wg sync.WaitGroup
	tasks := make(chan string)

	// Запускаем воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for url := range tasks {
				fmt.Printf("[Worker %d] Downloading: %s\n", workerId, url)
				err := downloader.DownloadImage(url, folder)
				if err != nil {
					log.Printf("[Worker %d] Error: %s\n", workerId, err)
					appendError(url, err)
				} else {
					fmt.Printf("[Worker %d] Done: %s\n", workerId, url)
				}
			}
		}(i)
	}

	// Отправляем задачи
	for _, url := range urls {
		tasks <- url
	}
	close(tasks)

	wg.Wait()
	fmt.Println("✅ All downloads completed.")
}

func appendError(url string, err error) {
	f, ferr := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ferr != nil {
		log.Println("Failed to write error log:", ferr)
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "%s – %s\n", url, err)
}
