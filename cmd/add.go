/*
Copyright © 2022 Hikaru Imamoto
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// フラグバインド用の変数
var date, name string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Register anniversary",
	Long:  "Register anniversary",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := time.Parse("20060102", date); err != nil {
			return fmt.Errorf("invalid date format: %s", err.Error())
		}
		if len(name) == 0 {
			return errors.New("anniversary name is empty")
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get user home dir path: %s", err.Error())
		}

		dir := fmt.Sprintf("%s/.anniv", home)
		if s, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0777); err != nil {
				return fmt.Errorf("cannot create directory: %s", err.Error())
			}
		} else if !s.IsDir() {
			return fmt.Errorf("%s is not directory", dir)
		}

		f := fmt.Sprintf("%s/data.csv", dir)
		fp, err := os.OpenFile(f, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("cannot create or open file: %s", err.Error())
		}
		defer fp.Close()

		w := csv.NewWriter(fp)
		w.Write([]string{date, name, tag})
		w.Flush()
		if err := w.Error(); err != nil {
			return err
		}

		fmt.Println("Success!!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	//フラグの値を変数にバインド
	addCmd.Flags().StringVarP(&date, "date", "D", "", "Anniversary date in the format 'YYYYMMDD'")
	addCmd.Flags().StringVar(&name, "name", "", "Anniversary date name")

	//必須のフラグに指定
	addCmd.MarkFlagRequired("date")
	addCmd.MarkFlagRequired("name")
}
