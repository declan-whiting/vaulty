package search

import (
	"github.com/declan-whiting/vaulty/internal/theme"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Oberservers interface {
	NotifyUpdate(content string)
}

type SearchView struct {
	*tview.InputField
	FocusSecretsHandler func()
	oberservers         []Oberservers
}

func (sv *SearchView) AddObserver(o Oberservers) {
	sv.oberservers = append(sv.oberservers, o)
}

func NewSearchView(focusSecretsHandler func()) *SearchView {
	search := &SearchView{
		FocusSecretsHandler: focusSecretsHandler,
	}
	search.InputField = tview.NewInputField()
	search.SetBorder(true)
	search.SetTitle("Search")

	return search
}

func (sv *SearchView) AddSearchControls() {
	sv.SetFieldBackgroundColor(theme.NewTheme().GetColor("background"))

	sv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			sv.FocusSecretsHandler()
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
