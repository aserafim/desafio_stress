package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
}

func worker(url string, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{}
	for range jobs {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- Result{StatusCode: 0, Duration: duration}
			continue
		}

		results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		resp.Body.Close()
	}
}

func main() {
	// CLI params
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Uso: --url=<url> --requests=<n> --concurrency=<n>")
		return
	}

	fmt.Printf("Iniciando teste de carga:\nURL: %s\nRequests: %d\nConcorrência: %d\n\n", *url, *requests, *concurrency)

	jobs := make(chan int, *requests)
	results := make(chan Result, *requests)
	var wg sync.WaitGroup

	start := time.Now()

	// Iniciar workers
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker(*url, jobs, results, &wg)
	}

	// Enviar jobs
	for i := 0; i < *requests; i++ {
		jobs <- i
	}
	close(jobs)

	// Esperar todos os workers
	wg.Wait()
	close(results)

	// Analisar resultados
	total200 := 0
	statusDist := make(map[int]int)

	for result := range results {
		statusDist[result.StatusCode]++
		if result.StatusCode == 200 {
			total200++
		}
	}

	totalTime := time.Since(start)

	// Relatório final
	fmt.Println("\nRelatório de Teste:")
	fmt.Printf("Tempo total: %s\n", totalTime)
	fmt.Printf("Total de requests: %d\n", *requests)
	fmt.Printf("HTTP 200: %d\n", total200)

	fmt.Println("Distribuição de status HTTP:")
	for code, count := range statusDist {
		fmt.Printf("  %d: %d\n", code, count)
	}
}
