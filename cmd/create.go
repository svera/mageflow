package cmd

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new pull request",
	Long:  `Creates a new pull request against the team repository or the mainline one.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		form := tview.NewForm().
			AddDropDown("Version", []string{"2.3-develop", "2.3.1-develop", "2.3.2-develop", "2.4-develop"}, 0, nil).
			AddCheckbox("Magento CE: MC-12345", false, nil).
			AddCheckbox("Magento EE: MC-12346", false, nil).
			AddCheckbox("Magento B2B: MC-12347", false, nil).
			AddCheckbox("Magento Infra: MC-12348", false, nil).
			AddCheckbox("Bypass PR validation", false, nil).
			AddButton("Create", func() {
				app.Stop()
				fmt.Println("Getting delivery profile for 2.4-develop...")
				time.Sleep(2 * time.Second)
				fmt.Println("Pushing local branches to Github...")
				time.Sleep(2 * time.Second)
				fmt.Println("Running Jenkins's CreatePR job...")
				time.Sleep(2 * time.Second)
				fmt.Println("")
				fmt.Println("Jenkins job created with ID #5555 magento-lynx:MC-12345 with the following PRs:")
				fmt.Println("* CE PR: github.com/magento-lynx/magento2ce/pull/10")
				fmt.Println("* EE PR: github.com/magento-lynx/magento2ee/pull/11")
				fmt.Println("* B2B PR: github.com/magento-lynx/magento2b2b/pull/12")
				fmt.Println("")
				fmt.Println("Type 'mageflow status 5555' to check PRs status")
			}).
			AddButton("Cancel", func() {
				app.Stop()
			})
		if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}
