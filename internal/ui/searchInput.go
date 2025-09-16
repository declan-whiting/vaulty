package ui

import (
	"strings"

	"github.com/declan-whiting/vaulty/internal/cache"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewSearchView() *tview.InputField {
	search := tview.NewInputField()
	search.SetBorder(true)
	search.SetTitle("Search")

	return search
}

func (ui *Ui) AddSearchControls() *Ui {
	ui.SearchView.SetFieldBackgroundColor(ui.KeyvaultView.GetBackgroundColor())

	ui.SearchView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			ui.App.SetFocus(ui.SecretsView)
			return nil
		}
		return event
	})

	ui.SearchView.SetChangedFunc(func(text string) {

		ui.SecretsView.Clear()
		i := 0

		for _, v := range cache.ReadSecrets(ui.CurrentKeyVault.Name) {
			if strings.Contains(strings.ToLower(v.Name), strings.ToLower(text)) {
				ui.SecretsView.SetCell(i, 0, tview.NewTableCell(v.Name))
				i++
			}

		}

	})

	return ui
}
