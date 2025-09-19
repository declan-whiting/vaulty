package controls

import (
	"github.com/rivo/tview"
)

type ControlsView struct {
	*tview.Table
}

type CellThemer interface {
	SetTableCellTheme(table *tview.Table, row int, col int, foreground, background string)
}

func NewControlsView(themer CellThemer) *ControlsView {
	cv := &ControlsView{}
	cv.Table = tview.NewTable()
	cv.SetBorder(true)
	cv.SetTitle("Controls")

	col := 0
	index := 0
	for _, v := range cv.CreateControlsHelp() {
		if index == 3 {
			index = 0
			col += 2
		}

		cv.SetCell(index, col, tview.NewTableCell(v.Message))
		cv.SetCell(index, col+1, tview.NewTableCell(v.Key))
		themer.SetTableCellTheme(cv.Table, index, col+1, "orange", "background")
		index++
	}

	return cv
}

func (cv ControlsView) CreateControlsHelp() []ControlsHelp {
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
