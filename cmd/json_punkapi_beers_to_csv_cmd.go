package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/soulcodex/json-to-csv-cli/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	output        = "output"
	defaultOutput = "data/beers-%d.csv"
	baseUrl       = "https://api.punkapi.com/v2%s"
)

func PunkApiJsonToCsvCmd() *cobra.Command {
	punkApiJsonToCsvCmd := &cobra.Command{
		Use:   "beers",
		Short: "Extract beers from Punk API data and dump to CSV file",
		Run:   runJsonBeersToCsv(),
	}

	punkApiJsonToCsvCmd.Flags().StringP(output, "o", "", "Destination path for CSV file")
	return punkApiJsonToCsvCmd
}

func runJsonBeersToCsv() model.CobraFn {
	return func(cmd *cobra.Command, args []string) {
		o, e := cmd.Flags().GetString(output)
		destination := map[bool]string{
			true:  o,
			false: fmt.Sprintf(defaultOutput, time.Now().UnixMilli()),
		}[e == nil && len(o) > 0]

		c := http.Client{Timeout: time.Duration(1) * time.Second}

		resp, e := c.Get(fmt.Sprintf(baseUrl, "/beers"))

		if e != nil {
			fmt.Printf("Error: %s", e)
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var beers []model.Beer

		if err := json.Unmarshal(body, &beers); err != nil {
			log.Fatal(err)
		}

		var beersRecords [][]string
		headers := []string{"id", "name", "tagline", "description", "contributed_by"}
		beersRecords = append(beersRecords, headers)

		for _, beer := range beers {
			beerRow := []string{
				strconv.Itoa(beer.Id),
				beer.Name,
				beer.Tagline,
				beer.Description,
				beer.ContributedBy,
			}
			beersRecords = append(beersRecords, beerRow)
		}

		WriteBeersInCsv(beersRecords, destination)
	}
}

func WriteBeersInCsv(records [][]string, destination string) {
	f, err := os.Create(destination)

	w := csv.NewWriter(f)
	err = w.WriteAll(records)

	if err != nil {
		log.Fatal(err)
	}

	w.Flush()
}
