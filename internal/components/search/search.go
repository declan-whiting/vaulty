package search

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Oberservers interface {
	NotifyUpdate(content string)
}

type Themer interface {
	GetColor(color string) tcell.Color
}

type SearchView struct {
	*tview.InputField
	FocusSecretsHandler func()
	OnEscape            func()
	oberservers         []Oberservers
}

func (sv *SearchView) AddObserver(o Oberservers) {
	sv.oberservers = append(sv.oberservers, o)
}

func NewSearchView(focusSecretsHandler, onEscape func(), theme Themer) *SearchView {
	search := &SearchView{
		FocusSecretsHandler: focusSecretsHandler,
		OnEscape:            onEscape,
	}
	search.InputField = tview.NewInputField()
	search.SetBorder(true)
	search.SetTitle("Search")
	search.SetFieldBackgroundColor(theme.GetColor("background"))
	search.AddSearchControls()
	return search
}

func (sv *SearchView) AddSearchControls() {
	sv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			sv.FocusSecretsHandler()
			return nil
		}
		if event.Key() == tcell.KeyEscape {
			sv.OnEscape()
		}
		return event
	})

	sv.SetChangedFunc(func(text string) {

		for _, v := range sv.oberservers {
			v.NotifyUpdate(text)
		}

	})

}
