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
}

func (ui *Ui) Init() *Ui {
	ui.CurrentKeyVault = new(CurrentKeyVault)
	ui.SecretsView = NewSecretsView()
	ui.KeyvaultView = NewKeyvaultView()
	ui.ControlsView = NewControlsView()
	ui.SearchView = NewSearchView()
	ui.StatusView = NewStatusView()
	ui.App = tview.NewApplication()
	return ui
}

func BuildUi() {
	ui := new(Ui).
		Init().
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
