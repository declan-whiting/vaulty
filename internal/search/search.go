package search

import (
	"github.com/declan-whiting/vaulty/internal/theme"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Oberservers interface {
	NotifyFocus()
	NotifyUpdate(content string)
}

type SearchView struct {
	*tview.InputField
	oberservers []Oberservers
}

func (sv *SearchView) AddObserver(o Oberservers) {
	sv.oberservers = append(sv.oberservers, o)
}

func NewSearchView() *SearchView {
	search := &SearchView{}
	search.InputField = tview.NewInputField()
	search.SetBorder(true)
	search.SetTitle("Search")

	return search
}

func (sv *SearchView) AddSearchControls() {
	sv.SetFieldBackgroundColor(theme.NewTheme().GetColor("background"))

	sv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			//ui.App.SetFocus(ui.SecretsView)
			for _, v := range sv.oberservers {
				v.NotifyFocus()
			}
			return nil
		}
		return event
	})

	sv.SetChangedFunc(func(text string) {

		for _, v := range sv.oberservers {
			v.NotifyUpdate(text)
		}

	})

}
