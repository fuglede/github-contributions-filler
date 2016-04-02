package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

// getUserRegistrationDate produces the time at which a given GitHub user
// created their account. 
func getUserRegistrationDate(username string) (creationTime time.Time, err error) { 
	gitHubApiUrl := "https://api.github.com/users/" + username
	response, err := http.Get(gitHubApiUrl)
	if err != nil {
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	var user map[string]interface{}
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return
	}
	creationTime, err = time.Parse(time.RFC3339Nano, user["created_at"].(string))
	return 
}

// makeDates creates a slice of Poisson distributed times between given dates with a
// specified mean waiting time (given in milliseconds).
func makeDates(startingDate time.Time, endingDate time.Time, meanWaitingTime float64) []time.Time {
	currentDate := startingDate
	for currentDate.Before(endingDate) {
		// Poisson waiting times are exponentially distributed
		waitingTime := - meanWaitingTime * math.Log(rand.Float64())
		duration := time.Duration(time.Duration(waitingTime) * 1000000 * time.Nanosecond)
		currentDate = currentDate.Add(duration)
		fmt.Println(currentDate)
	}
	return nil
}


func main() {
	registrationDate, err := getUserRegistrationDate("fuglede")
	lastDate, _ := time.Parse(time.RFC3339Nano, "2025-01-01T00:42:42Z")
	if err != nil {
		log.Fatal("Could not get registration date from GitHub: ", err)
	}
	// We use a mean of 18000000 milliseconds = five hours
	makeDates(registrationDate, lastDate, 18000000)
}