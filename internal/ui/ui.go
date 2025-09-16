package ui

import (
	"log"

	"github.com/rivo/tview"
)

type CurrentKeyVault struct {
	Name           string
	SubscriptionId string
}

type Ui struct {
	App             *tview.Application
	Grid            *tview.Grid
	ControlsView    *tview.Table
	KeyvaultView    *tview.List
	SecretsView     *tview.Table
	SearchView      *tview.InputField
	StatusView      *tview.TextView
	CurrentKeyVault *CurrentKeyVault
	Services        *Services
}

func (ui *Ui) Init(services Services) *Ui {
	ui.Services = &services
	ui.CurrentKeyVault = new(CurrentKeyVault)
	ui.SecretsView = NewSecretsView(services)
	ui.KeyvaultView = NewKeyvaultView(services)
	ui.ControlsView = NewControlsView()
	ui.SearchView = NewSearchView()
	ui.StatusView = NewStatusView(services)
	ui.App = tview.NewApplication()

	return ui
}

func BuildUi() {
	services := Services{}
	services.Init()

	ui := new(Ui).
		Init(services).
		CreateGrid().
		AddKeyvaultViewControls().
		AddSecretsControls().
		AddSearchControls().
		AddStatusControls().
		KeyvaultSelectionChanged().
		SecretSelectedChanged()

	ui.App.SetRoot(ui.Grid, true)
	err := ui.App.SetFocus(ui.KeyvaultView).Run()
	if err != nil {
		log.Fatal(err)
	}
}
