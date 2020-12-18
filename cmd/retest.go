package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	all bool
)

func init() {
	rootCmd.AddCommand(retestCmd)
	retestCmd.Flags().BoolVarP(&all, "all", "a", false, "Rerun all tests")
}

var retestCmd = &cobra.Command{
	Use:   "retest",
	Short: "Retrigger tests of a specific job (only the failed ones by default)",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires Jenkins job ID")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gathering information...")
		time.Sleep(2 * time.Second)
		if all {
			fmt.Printf("Executed all tests of job %s\n", args[0])
		} else {
			fmt.Printf("Executed failed tests of job %s\n", args[0])
		}

	},
}
