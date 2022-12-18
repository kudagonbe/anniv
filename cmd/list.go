/*
Copyright © 2022 Hikaru Imamoto
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Refer anniversary list",
	Long:  "Refer anniversary list",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get user home dir path: %s", err.Error())
		}

		f := fmt.Sprintf("%s/.anniv/data.csv", home)
		if s, err := os.Stat(f); os.IsNotExist(err) {
			fmt.Println("date,name,tag")
			return nil
		} else if s.IsDir() {
			return fmt.Errorf("%s is directory", f)
		}

		fp, err := os.Open(f)
		if err != nil {
			return fmt.Errorf("cannot open file: %s", err.Error())
		}
		defer fp.Close()

		r := csv.NewReader(fp)

		fmt.Println("date,name,tag")
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return fmt.Errorf("cannot read file: %s", err.Error())
			}
			if tag == "" || (len(record) >= 3 && record[2] == tag) {
				fmt.Println(strings.Join(record, ","))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
