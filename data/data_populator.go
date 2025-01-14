package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func main() {

	files := []string{"coupons/couponbase1.gz", "coupons/couponbase2.gz", "coupons/couponbase3.gz"}

	for index, filename := range files {
		err := processAndCache(filename, index+1)
		if err != nil {
			log.Fatalf("Error processing file %s: %v", filename, err)
		}
	}

	fmt.Println("Processed all files and cached in redis")
}

func processAndCache(filename string, index int) error {

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("failed to create gzip reader: %v", err)
	}
	defer gzReader.Close()

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	scanner := bufio.NewScanner(gzReader)
	for scanner.Scan() {
		line := scanner.Text()
		//if any empty line
		if line == "" {
			continue
		}

		value := "file" + strconv.Itoa(index)
		err := rdb.SAdd(ctx, line, value).Err()
		if err != nil {
			return fmt.Errorf("failed to write to Redis: %v", err)
		}
		//fmt.Printf("Added to Redis: key=%s, value=%s\n", line, value)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading decompressed content: %v", err)
	}

	return nil
}
