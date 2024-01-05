package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	file, err := os.OpenFile("../DONOTOPEN/weather1gi.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	// write header
	headerBuf := new(bytes.Buffer)
	header := fmt.Sprintf("%s,%s\n", "name", "temperature")
	headerBuf.WriteString(header)
	file.Write(headerBuf.Bytes())

	wrBuf := bufio.NewWriter(file)

	for i := 0; i < size; i++ {

		if i > 0 && (i%100_000_000 == 0) {
			fmt.Printf("checkpoint: %d. flushing and fsyncing the buffer. Time till now: %.3fsecs\n", i, float64(time.Since(start))/float64(time.Second))

			if err := wrBuf.Flush(); err != nil {
				fmt.Println("error in flushing to io writer", err)
			}

			if err := file.Sync(); err != nil {
				fmt.Println("error in fsyncing", err)
			}
		}

		pickStation := stations[rand.Intn(len(stations))]
		newStation := fmt.Sprintf("%s,%.3f\n", pickStation.Station, measurement(pickStation.Temperature))

		// write to buf
		_, err := wrBuf.WriteString(newStation)
		if err != nil {
			fmt.Println("error in writing to the buffer", err)
			continue
		}

		// write to file
		//_, err = file.Write(buf.Bytes())
		//if err != nil {
		//	fmt.Println("error in writing to the file", err)
		//	continue
		//}
	}

	// final flush
	if err := wrBuf.Flush(); err != nil {
		fmt.Println("error in flushing to io writer", err)
	}

	// final fsync
	if err := file.Sync(); err != nil {
		fmt.Println("error in writing to the buffer", err)
	}

}

func measurement(meanTemperature float32) float32 {
	m := rand.NormFloat64()*10 + float64(meanTemperature)
	return float32(math.Round(m*10.0) / 10.0)
}

/*
unbuffered writes:
checkpoint: 100000000. fsyncing the buffer. Time till now: 255.106secs
checkpoint: 200000000. fsyncing the buffer. Time till now: 503.731secs
checkpoint: 300000000. fsyncing the buffer. Time till now: 744.789secs

buffered writes: buffered write will automatically flush the buffer on reaching default size of 4KB and will empty the buffer for new writes.
checkpoint: 100000000. flushing and fsyncing the buffer. Time till now: 26.980secs
checkpoint: 200000000. flushing and fsyncing the buffer. Time till now: 53.791secs
checkpoint: 300000000. flushing and fsyncing the buffer. Time till now: 80.649secs
*/
