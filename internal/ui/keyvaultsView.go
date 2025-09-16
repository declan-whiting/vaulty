package ui

import (
	"github.com/declan-whiting/vaulty/internal/cache"
	"github.com/declan-whiting/vaulty/internal/configuration"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewKeyvaultView() *tview.List {
	list := tview.NewList()
	list.SetTitle("Keyvaults")
	list.SetBorder(true)
	list.ShowSecondaryText(false)

	for i, v := range cache.ReadKeyvaults() {
		list.AddItem(v.Name, v.SubscriptionId, rune('a'+i), nil)
	}

	return list
}

func (ui *Ui) AddKeyvaultViewControls() *Ui {
	ui.KeyvaultView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			ui.App.Stop()
			return tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone)
		}
		if event.Rune() == '/' {
			ui.SearchView.SetText("")
			ui.SecretsView.ScrollToBeginning()
			ui.App.SetFocus(ui.SearchView)
		}
		if event.Key() == tcell.KeyEnter {
			ui.App.SetFocus(ui.SecretsView)
			return nil
		}
		return event
	})

	vault := configuration.GetConfiguration().Keyvaults[0]
	ui.CurrentKeyVault.Name = vault.Name
	ui.CurrentKeyVault.SubscriptionId = vault.SubscriptionId
	UpdateSecretsView(ui)
	ui.SecretsView.SetTitle(ui.CurrentKeyVault.Name + "/secrets")

	return ui
}

func (ui *Ui) KeyvaultSelectionChanged() *Ui {
	ui.KeyvaultView.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		ui.CurrentKeyVault.Name = mainText
		ui.CurrentKeyVault.SubscriptionId = secondaryText
		UpdateSecretsView(ui)
		ui.SecretsView.SetTitle(mainText + "/secrets")
	})

	// //TODO: hack to populate
	// ui.KeyvaultView.SetCurrentItem(0)

	return ui
}
