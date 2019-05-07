package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"googlemaps.github.io/maps"
)

/*Park is a struct containing name and address of state park*/
type Park struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	ogAddress, destAddress := ReadAddressData("StateParkAddresses.csv", 10)
	WriteDistanceData(ogAddress, destAddress, "Insert API Key here")
	Prims("DistanceResponse.txt")
}

/*ReadAddressData s address data from file name, with limit on origin/dest addresses limit*/
func ReadAddressData(name string, limit int) (originAddresses []string, destAddresses []string) {
	dat, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
		return
	}
	r := csv.NewReader(dat)
	var parks []Park
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		parks = append(parks, Park{
			Name:    record[0],
			Address: record[1],
		})
	}
	var parkAddresses []string

	for i, park := range parks {
		//Limit below Free Google Maps API
		if i < limit {
			parkAddresses = append(parkAddresses, park.Address)
		}

	}
	return parkAddresses, parkAddresses
}

/*WriteDistanceData writes DistanceMatrix API data to file*/
func WriteDistanceData(originAddresses []string, destAddresses []string, apiKey string) {

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		fmt.Println("uh oh google maps")
		log.Fatal(err)
	}
	request := &maps.DistanceMatrixRequest{
		Origins:       originAddresses,
		Destinations:  destAddresses,
		DepartureTime: `now`,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := c.DistanceMatrix(ctx, request)
	defer cancel()
	if err != nil {
		fmt.Printf("Error was not null, error was %+v", err)
		return
	}

	adjMatrix := make([][]int, 10)
	for z := 0; z < 10; z++ {
		adjMatrix[z] = make([]int, 10)
	}

	f, err := os.Create("DistanceResponse.txt")
	b, err := json.Marshal(resp)
	f.Write(b)
}
