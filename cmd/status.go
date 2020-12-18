package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks the status of the pull requests created by a job",
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

		app := tview.NewApplication()
		table := tview.NewTable().SetBorders(true)
		table.SetCell(0, 1, tview.NewTableCell(" MAGENTO CE (No approvals)"))
		table.SetCell(0, 2, tview.NewTableCell(" MAGENTO EE (1 approval)").SetTextColor(tcell.ColorGreen))
		table.SetCell(0, 3, tview.NewTableCell(" MAGENTO B2B (1 change request)").SetTextColor(tcell.ColorRed))
		table.SetCell(0, 4, tview.NewTableCell(" MAGENTO INFRA "))

		table.SetCell(1, 0, tview.NewTableCell("ci/jenkins/with-exts/Functional-Tests-B2B-PR "))
		table.SetCell(2, 0, tview.NewTableCell("ci/jenkins/with-exts/Functional-Tests-EE-PR "))
		table.SetCell(3, 0, tview.NewTableCell("ci/jenkins/with-exts/Integration-Tests-PR "))
		table.SetCell(4, 0, tview.NewTableCell("ci/jenkins/with-exts/Magento-Health-Index-PR "))
		table.SetCell(5, 0, tview.NewTableCell("ci/jenkins/with-exts/Semantic-Version-Checker-PR "))
		table.SetCell(6, 0, tview.NewTableCell("ci/jenkins/with-exts/Static-Tests-PR "))
		table.SetCell(7, 0, tview.NewTableCell("ci/jenkins/with-exts/Unit-Tests-PR "))
		table.SetCell(8, 0, tview.NewTableCell("ci/jenkins/with-exts/WebAPI-Tests-PR "))
		table.SetCell(9, 0, tview.NewTableCell("ci/jenkins/with-exts/Functional-Tests-CE-PR "))
		table.SetCell(10, 0, tview.NewTableCell("ci/jenkins/with-exts/Database-Compare-PR "))

		cols, rows := 5, 11
		for r := 1; r < rows; r++ {
			for c := 1; c < cols; c++ {
				color := tcell.ColorGreen
				fgColor := tcell.ColorWhite
				if r == 4 && c == 3 {
					color = tcell.ColorYellow
					fgColor = tcell.ColorBlack
				}
				if r == 4 && c == 4 {
					color = tcell.ColorRed
					fgColor = tcell.ColorWhite
				}
				table.SetCell(r, c,
					tview.NewTableCell(" VIEW ").
						SetTextColor(fgColor).
						SetBackgroundColor(color))
			}
		}
		table.Select(1, 1).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
			if key == tcell.KeyEnter {
				table.SetSelectable(true, true)
			}
		})
		if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
			panic(err)
		}
		app.Stop()
	},
}
