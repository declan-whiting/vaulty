package ui

import (
	"time"

	"github.com/declan-whiting/vaulty/internal/events"
	"github.com/declan-whiting/vaulty/internal/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (ui *Ui) CreateGrid() *Ui {
	ui.Grid = tview.NewGrid().SetRows(5, 3, 3).SetColumns(20, 0)
	ui.Grid.AddItem(ui.ControlsView, 0, 0, 1, 2, 0, 0, false)
	ui.Grid.AddItem(ui.StatusView, 0, 2, 1, 1, 0, 0, false)
	ui.Grid.AddItem(ui.KeyvaultView, 1, 0, 3, 1, 0, 0, false)
	ui.Grid.AddItem(ui.SecretsView, 1, 1, 3, 2, 0, 0, false)
	ui.AddStatusControls()
	return ui
}

func (ui *Ui) HideSearch() {
	ui.Grid.RemoveItem(ui.SearchView)
	ui.Grid.RemoveItem(ui.KeyvaultView)
	ui.Grid.RemoveItem(ui.SecretsView)
	ui.Grid.AddItem(ui.KeyvaultView, 1, 0, 3, 1, 0, 0, false)
	ui.Grid.AddItem(ui.SecretsView, 1, 1, 3, 2, 0, 0, false)
}

func (ui *Ui) ShowSearch() {
	ui.App.SetFocus(ui.SearchView)
	ui.Grid.RemoveItem(ui.KeyvaultView)
	ui.Grid.RemoveItem(ui.SecretsView)
	ui.Grid.AddItem(ui.SearchView, 1, 0, 1, 3, 0, 0, false)
	ui.Grid.AddItem(ui.KeyvaultView, 2, 0, 2, 1, 0, 0, false)
	ui.Grid.AddItem(ui.SecretsView, 2, 1, 2, 2, 0, 0, false)
}

func (ui *Ui) AddStatusControls() *Ui {
	ui.Grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlR) {
			start := time.Now()
			ui.Events.NewEvent("\U0001F510 Synchronise Started", "Took: 0.1s")
			var vaults []models.KeyvaultModel

			go func() {
				defer ui.App.QueueUpdateDraw(func() { events.TimedEventLog(start, "\U0001F510 Synchronise Finished", *ui.Events) })
				config := ui.Services.ConfigrationService.GetConfiguration()

				for _, v := range config.Keyvaults {
					vaults = append(vaults, ui.Services.AzureService.AzShowKeyvault(v.Name, v.SubscriptionId))
				}

				for i, v := range vaults {
					vaults[i].Secrets = ui.Services.AzureService.AzGetSecrets(v.Name, v.SubscriptionId)
				}
			}()

		}

		return event
	})

	return ui
}
