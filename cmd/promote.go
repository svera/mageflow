package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	force bool
)

func init() {
	rootCmd.AddCommand(promoteCmd)
	promoteCmd.Flags().BoolVarP(&force, "force", "f", false, "Force promoting ")
}

var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Copies PRs of a job from the team organisation to the mainline",
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
		if force {
			fmt.Printf("Promoted job %s, mainline ID is %s\n", args[0], "42")
		} else {
			fmt.Println("Cannot promote job with unfinished or failed tests. Use --force to promote it anyway\f", args[0])
		}
	},
}
