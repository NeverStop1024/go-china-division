package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cd ChinaDivision

var RestCmd = &cobra.Command{
	Use:     "gain",
	Short:   "gain get zoning code and urban-rural division",
	Example: `./main division -o 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(cd)
		return Run(cd)
	},
}

func init() {
	RestCmd.Flags().IntVarP((*int)(&cd.Option), "option", "o", int(OptionProvinceCityCounty), "option plan")
	RestCmd.Flags().StringVarP(&cd.OutPath, "outPath", "p", "./", "generate file path")
	RestCmd.Flags().StringVarP(&cd.FileName, "fileName", "f", "china_.json", "generate file filename")
}
