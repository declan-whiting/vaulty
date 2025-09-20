package ui

import (
	"log"

	"github.com/declan-whiting/vaulty/internal/controls"
	"github.com/declan-whiting/vaulty/internal/keyvault"
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
	KeyvaultView    *keyvault.KeyvaultView
	SecretsView     *tview.Table
	SearchView      *search.SearchView
	StatusView      *tview.TextView
	CurrentKeyVault *CurrentKeyVault
	Services        *Services
}

func (ui *Ui) Init(services Services, themer theme.Theme) *Ui {
	ui.App = tview.NewApplication()
	ui.Services = &services
	ui.CurrentKeyVault = new(CurrentKeyVault)
	ui.SecretsView = NewSecretsView(services)

	ui.KeyvaultView = keyvault.NewKeyvaultView(
		services.CacheService,
		services.ConfigrationService,
		ui.HandleQuit,
		ui.HandleSearch,
		ui.HandleKeyvaultSelection)

	ui.ControlsView = controls.NewControlsView(themer)
	ui.SearchView = search.NewSearchView()
	ui.StatusView = NewStatusView(services)

	return ui
}

func (ui *Ui) HandleQuit() {
	ui.App.Stop()
}
func (ui *Ui) HandleSearch() {
	ui.SearchView.SetText("")
	ui.SecretsView.ScrollToBeginning()
	ui.App.SetFocus(ui.SearchView)
}
func (ui *Ui) HandleKeyvaultSelection() {
	ui.App.SetFocus(ui.SecretsView)
}

func BuildUi() {
	services := Services{}
	services.Init()

	theme := theme.NewTheme()
	tview.Styles = theme.GetTheme()

	ui := new(Ui).
		Init(services, theme).
		CreateGrid().
		AddSecretsControls().
		AddStatusControls().
		SecretSelectedChanged().
		StyleCustomization(theme)

	ui.SearchView.AddSearchControls()
	ui.SearchView.AddObserver(ui)

	ui.KeyvaultView.AddCurrentKeyvaultWatcher(ui)
	ui.KeyvaultView.SetInitialKeyvault()

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
