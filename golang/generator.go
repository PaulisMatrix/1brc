package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const size = 1_000_000_000

var stations []WeatherStations

type WeatherStations struct {
	Station     string  `json:"station"`
	Temperature float32 `json:"temperature"`
}

func constructList() {
	// read the input.txt and construct the stations list
	// from which we will randomly pick station and write it to our csv file 1B times
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		parts := strings.Split(line, ",")

		temp, err := strconv.ParseFloat(parts[1], 32)
		if err != nil {
			panic(err)
		}

		currStation := WeatherStations{
			Station:     parts[0],
			Temperature: float32(temp),
		}

		stations = append(stations, currStation)

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

}

func generate() {
	// func to generate 1B rows

	// get the whole list
	constructList()
	fmt.Println(len(stations))

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	file, err := os.OpenFile("weather1gi.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	// write header
	headerBuf := new(bytes.Buffer)
	header := fmt.Sprintf("%s,%s\n", "name", "temperature")
	headerBuf.WriteString(header)
	file.Write(headerBuf.Bytes())

	for i := 0; i < size; i++ {

		if i > 0 && (i%100_000_000 == 0) {
			fmt.Printf("checkpoint: %d. fsyncing the buffer. Time till now: %.3fsecs\n", i, float64(time.Since(start))/float64(time.Second))
			if err := file.Sync(); err != nil {
				fmt.Println("error in fsyncing", err)
			}
		}

		buf := new(bytes.Buffer)

		pickStation := stations[rand.Intn(len(stations))]
		newStation := fmt.Sprintf("%s,%.3f\n", pickStation.Station, pickStation.Temperature)

		// write to buf
		_, err := buf.WriteString(newStation)
		if err != nil {
			fmt.Println("error in writing to the buffer", err)
			continue
		}

		// write to file
		_, err = file.Write(buf.Bytes())
		if err != nil {
			fmt.Println("error in writing to the file", err)
			continue
		}
	}

	// final fsync
	if err := file.Sync(); err != nil {
		fmt.Println("error in writing to the buffer", err)
	}

}
