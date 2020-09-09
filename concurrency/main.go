package main

import (
	"bufio"
	"fmt"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/repositories"
	"github.com/TeplyyMaksim/golang-microservices/src/api/services"
	"github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"
	"os"
	"sync"
)

var (
	success map[string]string
	failed map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result *repositories.CreateRepoResponse
	Error errors.ApiError
}


func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/Users/maksymteplyy/go/src/github.com/TeplyyMaksim/golang-microservices/concurrency/index.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}

	return result
}

func main () {
	requests := getRequests()

	fmt.Printf("about to process %d requests\n", len(requests))

	input := make(chan createRepoResult)
	buff := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(input, &wg)

	for _, request := range requests {
		wg.Add(1)
		buff <- true
		go createRepo(buff, input, request)
	}
	wg.Wait()
	close(input)
}

func handleResults(input chan createRepoResult, wg *sync.WaitGroup) {
	for result := range input {
		if result.Error != nil {
			failed[result.Request.Name] = result.Error
		} else {
			success[""] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(buffer chan bool, output chan createRepoResult, request repositories.CreateRepoRequest) {
	result, err := services.RepositoryService.CreateRepo(request)
	output <- createRepoResult{
		Request: request,
		Result: result,
		Error: err,
	}

	<- buffer
}