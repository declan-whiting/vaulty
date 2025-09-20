package ui

import (
	"fmt"
	"strings"

	"github.com/declan-whiting/vaulty/internal/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewSecretsView(service Services) *tview.Table {
	keyvault := models.KeyvaultModel{}
	table := tview.NewTable()
	table.SetSelectable(true, false)
	table.SetBorder(true)

	table.SetTitle("Secrets")
	for i, v := range service.CacheService.ReadSecrets(keyvault.Name) {
		table.SetCell(i, 0, tview.NewTableCell(v.Name))
	}

	table.SetSelectable(true, false)
	return table
}

func UpdateSecretsView(ui *Ui) *tview.Table {
	ui.SecretsView.SetBorder(true)
	ui.SecretsView.SetTitle("Secrets")

	for i, v := range ui.Services.CacheService.ReadSecrets(ui.CurrentKeyVault.Name) {
		ui.SecretsView.SetCell(i, 0, tview.NewTableCell(v.Name))
	}

	return ui.SecretsView
}

func (ui *Ui) AddSecretsControls() *Ui {
	ui.SecretsView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			ui.App.Stop()
			return tcell.NewEventKey(tcell.KeyRune, 's', tcell.ModNone)
		}
		if event.Rune() == '/' {
			ui.SearchView.SetText("")
			ui.SecretsView.ScrollToBeginning()
			ui.App.SetFocus(ui.SearchView)
		}
		if event.Rune() == 'b' {
			ui.App.SetFocus(ui.KeyvaultView)
			return tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone)
		}
		return event
	})
	return ui
}

func (ui *Ui) SecretSelectedChanged() *Ui {
	ui.SecretsView.SetSelectedFunc(func(row, column int) {
		secretName := ui.SecretsView.GetCell(row, column).Text

		ui.Grid.RemoveItem(ui.SecretsView)
		secretsDetailsView := tview.NewTextView()
		secretsDetailsView.SetTitle(fmt.Sprintf("%s/%s", ui.CurrentKeyVault.Name, secretName))
		secretsDetailsView.SetBorder(true)
		secretsDetailsView.SetText(ui.Services.AzureService.AzShowSecret(secretName, ui.CurrentKeyVault.Name, ui.CurrentKeyVault.SubscriptionId))

		ui.App.SetFocus(secretsDetailsView)

		secretsDetailsView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'b' {
				ui.Grid.RemoveItem(secretsDetailsView)
				ui.Grid.AddItem(ui.SecretsView, 1, 1, 1, 2, 0, 0, false)
				ui.App.SetFocus(ui.SecretsView)
				return tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone)
			}
			return event
		})

		ui.Grid.AddItem(secretsDetailsView, 1, 1, 1, 2, 0, 0, false)
	})
	return ui
}

func (ui *Ui) NotifyUpdate(content string) {
	ui.SecretsView.Clear()
	i := 0
	for _, v := range ui.Services.CacheService.ReadSecrets(ui.CurrentKeyVault.Name) {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(content)) {
			ui.SecretsView.SetCell(i, 0, tview.NewTableCell(v.Name))
			i++
		}

	}
}
func (ui *Ui) NotifyFocus() {
	ui.App.SetFocus(ui.SecretsView)
}
