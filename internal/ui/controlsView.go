package ui

import "github.com/rivo/tview"

func NewControlsView() *tview.Table {
	table := tview.NewTable()
	table.SetBorder(true)
	table.SetTitle("Controls")

	i := 0
	col := 0
	for _, v := range CreateControlsHelp() {

		if i > 2 {
			i = 0
			col++
			table.SetCell(i, col, tview.NewTableCell(v.Message+" "+v.Key))
		} else {
			table.SetCell(i, col, tview.NewTableCell(v.Message+" "+v.Key))
		}
		i++
	}

	return table
}

func CreateControlsHelp() []ControlsHelp {
	return []ControlsHelp{
		{Key: "<q>", Message: "Quit"},
		{Key: "<Enter>", Message: "Show Secrets"},
		{Key: "<s>", Message: "Show Keyvault"},
		{Key: "<b>", Message: "Back to Menu"},
		{Key: "</>", Message: "Search"},
		{Key: "<ctrl+r>", Message: "Reload"},
	}
}

type ControlsHelp struct {
	Key     string
	Message string
}
