package main

import (
	"github.com/soulcodex/json-to-csv-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCli := &cobra.Command{Use: "beers-exporter"}
	rootCli.AddCommand(cmd.PunkApiJsonToCsvCmd())
	rootCli.Execute()
}
