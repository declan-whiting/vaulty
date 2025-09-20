package ui

import (
	"log"

	"github.com/declan-whiting/vaulty/internal/controls"
	"github.com/declan-whiting/vaulty/internal/search"
	"github.com/declan-whiting/vaulty/internal/theme"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CurrentKeyVault struct {
	Name           string
	SubscriptionId string
}

type Ui struct {
	App             *tview.Application
	Grid            *tview.Grid
	ControlsView    *controls.ControlsView
	KeyvaultView    *tview.List
	SecretsView     *tview.Table
	SearchView      *search.SearchView
	StatusView      *tview.TextView
	CurrentKeyVault *CurrentKeyVault
	Services        *Services
}

func (ui *Ui) Init(services Services, themer theme.Theme) *Ui {
	ui.Services = &services
	ui.CurrentKeyVault = new(CurrentKeyVault)
	ui.SecretsView = NewSecretsView(services)
	ui.KeyvaultView = NewKeyvaultView(services)
	ui.ControlsView = controls.NewControlsView(themer)
	ui.SearchView = search.NewSearchView()
	ui.StatusView = NewStatusView(services)
	ui.App = tview.NewApplication()

	return ui
}

func BuildUi() {
	services := Services{}
	services.Init()

	theme := theme.NewTheme()
	tview.Styles = theme.GetTheme()

	ui := new(Ui).
		Init(services, theme).
		CreateGrid().
		AddKeyvaultViewControls().
		AddSecretsControls().
		AddStatusControls().
		KeyvaultSelectionChanged().
		SecretSelectedChanged().
		StyleCustomization(theme)

	ui.SearchView.AddSearchControls()
	ui.SearchView.AddObserver(ui)

	ui.App.SetRoot(ui.Grid, true)
	err := ui.App.SetFocus(ui.KeyvaultView).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (ui *Ui) StyleCustomization(theme theme.Theme) *Ui {
	ui.KeyvaultView.SetSelectedBackgroundColor(theme.GetColor("background"))
	ui.KeyvaultView.SetSelectedTextColor(theme.GetColor("pink"))
	ui.SecretsView.SetSelectedStyle(tcell.StyleDefault.
		Background(theme.GetColor("background")).
		Foreground(theme.GetColor("pink")))

	return ui
}
