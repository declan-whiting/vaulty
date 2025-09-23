package secretDetails

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SecretDetailsView struct {
	*tview.TextView
}

func CreateSecretsDetailView(title, content string) *SecretDetailsView {
	details := &SecretDetailsView{}
	details.TextView = tview.NewTextView()
	details.SetBorder(true)
	details.SetTitle(title)
	details.SetText(content)

	return details
}

func (sdv *SecretDetailsView) AddControls(closer func(tv *tview.TextView)) {
	sdv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'b' || event.Key() == tcell.KeyEscape {
			closer(sdv.TextView)
			return tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone)
		}
		return event
	})
}
