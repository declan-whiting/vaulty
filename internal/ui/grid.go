package ui

import "github.com/rivo/tview"

func (ui *Ui) CreateGrid() *Ui {
	ui.Grid = tview.NewGrid().SetRows(5, 0, 3).SetColumns(20, 0)
	ui.Grid.AddItem(ui.ControlsView, 0, 0, 1, 2, 0, 0, false)
	ui.Grid.AddItem(ui.StatusView, 0, 2, 1, 1, 0, 0, false)
	ui.Grid.AddItem(ui.SearchView, 2, 0, 1, 3, 0, 0, false)
	ui.Grid.AddItem(ui.KeyvaultView, 1, 0, 1, 1, 0, 0, true)
	ui.Grid.AddItem(ui.SecretsView, 1, 1, 1, 2, 0, 0, false)

	return ui
}
