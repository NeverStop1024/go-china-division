/**
* @Author: TheLife
* @Date: 2021/8/7 下午2:18
 */
package main

import (
	"github.com/spf13/cobra"
	"go-china-division/app"
	"log"
)

//go:generate go build -o main
//go:generate ./go-china-division division

func main() {
	var rootCmd = &cobra.Command{Use: "china"}
	rootCmd.AddCommand(app.RestCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
