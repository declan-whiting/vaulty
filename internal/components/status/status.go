package status

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusView struct {
	*tview.Table
}

type Theme interface {
	GetColor(string) tcell.Color
}

func NewStatusView(t Theme) *StatusView {
	statusView := &StatusView{}
	statusView.Table = tview.NewTable()
	statusView.SetBorder(true)
	statusView.SetTitle("Log")
	return statusView
}

func (sv *StatusView) OnNewEvent(eventAction, eventTook string) {
	sv.InsertRow(0)
	sv.SetCellSimple(0, 0, eventAction)
	sv.SetCellSimple(0, 1, eventTook)
	sv.ScrollToBeginning()
}
