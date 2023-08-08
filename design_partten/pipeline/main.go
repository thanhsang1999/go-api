package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type City struct {
	Name       string
	Population int
}

func main() {

	// dọc từ 1 nguồn liên tục
	// chỉ đo luồng code sử lý

	f, err := os.Open("design_partten/pipeline/cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rows := genRows(f)
	filterSmallCity := filterByMinPopulation(3797699)
	ur1 := upperCityName(filterSmallCity(rows))
	ur2 := upperCityName(filterSmallCity(rows))
	ur3 := upperCityName(filterSmallCity(rows))
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		for c := range fanIn(ur1) {
			log.Println("ur1 ::: ", c)
		}
		wg.Done()
	}()
	go func() {
		for c := range fanIn(ur2) {
			log.Println("ur2 ::: ", c)
		}
		wg.Done()
	}()
	go func() {
		for c := range fanIn(ur3) {
			log.Println("ur3 ::: ", c)
		}
		wg.Done()
	}()
	wg.Wait()
}
func genRows(r io.Reader) chan City {
	out := make(chan City)

	go func() {
		reader := csv.NewReader(r)
		if _, err := reader.Read(); err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out)

	}()

	return out
}
func filterByMinPopulation(min int) func(<-chan City) <-chan City {
	return func(cities <-chan City) <-chan City {
		out := make(chan City)
		go func() {
			for c := range cities {
				if c.Population > min {
					out <- City{Name: c.Name, Population: c.Population}
				}
			}
			close(out)
		}()

		return out
	}
}
func upperCityName(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			time.Sleep(time.Second)
			out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
		}
		close(out)
	}()

	return out
}
func fanIn(chans ...<-chan City) <-chan City {
	out := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))

	for _, cities := range chans {
		go func(city <-chan City) {
			for c := range city {
				out <- c
			}
			wg.Done()
		}(cities)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
