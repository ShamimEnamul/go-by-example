package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"log"
	"os"
	"strings"
)

// Connect to Redis
func connectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	return rdb
}

// Load CSV data into Redis
func loadCSVToRedis(rdb *redis.Client, filePath string, key string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		userID := strings.TrimSpace(scanner.Text())              // Read and trim each line
		err := rdb.SAdd(context.Background(), key, userID).Err() // Add to Redis set
		if err != nil {
			log.Printf("Failed to add user ID %s to Redis: %v", userID, err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Printf("Loaded data into Redis set '%s' successfully.\n", key)
}
func readCsv(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("error occured while opening the file %s", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.Read()
	if err != nil {
		log.Fatalf("Error occured during file reading %s", err)
	}
	fmt.Println(data)
}

// Check if a user ID is eligible
func isUserEligible(rdb *redis.Client, key string, userID string) bool {
	isMember, err := rdb.SIsMember(context.Background(), key, userID).Result()
	if err != nil {
		log.Printf("Error checking user ID %s: %v", userID, err)
		return false
	}
	return isMember
}

func main() {
	// Connect to Redis
	rdb := connectRedis()
	defer rdb.Close()

	// Load CSV data into Redis
	filePath := "data.csv" // Replace with your CSV file path
	redisKey := "promo_user_ids"
	loadCSVToRedis(rdb, filePath, redisKey)

	// Example: Check eligibility of a user
	testUserID := "12345" // Replace with the user ID you want to check
	if isUserEligible(rdb, redisKey, testUserID) {
		fmt.Printf("User ID %s is eligible for the promo.\n", testUserID)
	} else {
		fmt.Printf("User ID %s is not eligible for the promo.\n", testUserID)
	}
	//readCsv("data.csv")
}
