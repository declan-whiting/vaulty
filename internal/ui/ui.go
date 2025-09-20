package ui

import (
	"fmt"
	"log"

	"github.com/declan-whiting/vaulty/internal/controls"
	"github.com/declan-whiting/vaulty/internal/keyvault"
	"github.com/declan-whiting/vaulty/internal/search"
	"github.com/declan-whiting/vaulty/internal/secrets"
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
	SecretsView     *secrets.SecretsView
	SearchView      *search.SearchView
	StatusView      *tview.TextView
	CurrentKeyVault *CurrentKeyVault
	Services        *Services
}

func (ui *Ui) Init(services Services, themer theme.Theme) *Ui {
	ui.App = tview.NewApplication()
	ui.Services = &services
	ui.CurrentKeyVault = new(CurrentKeyVault)

	ui.SecretsView = secrets.NewSecretsView(
		services.CacheService,
		ui.HandleQuit,
		ui.HandleSearch,
		ui.Handleback,
		ui.HandleSecretsSelectedChanged)

	ui.KeyvaultView = keyvault.NewKeyvaultView(
		services.CacheService,
		services.ConfigrationService,
		ui.HandleQuit,
		ui.HandleSearch,
		ui.FocusSecretsView)

	ui.ControlsView = controls.NewControlsView(themer)
	ui.SearchView = search.NewSearchView(ui.FocusSecretsView)
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

func (ui *Ui) FocusSecretsView() {
	ui.App.SetFocus(ui.SecretsView)
}
func (ui *Ui) Handleback() {
	ui.App.SetFocus(ui.KeyvaultView)
}

func (ui *Ui) HandleSecretsSelectedChanged(secret, keyvault, subscription string) {
	ui.Grid.RemoveItem(ui.SecretsView)
	secretsDetailsView := tview.NewTextView()
	secretsDetailsView.SetTitle(fmt.Sprintf("%s/%s", keyvault, secret))
	secretsDetailsView.SetBorder(true)
	secretsDetailsView.SetText(ui.Services.AzureService.AzShowSecret(secret, keyvault, subscription))

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
}

func BuildUi() {
	services := Services{}
	services.Init()

	theme := theme.NewTheme()
	tview.Styles = theme.GetTheme()

	ui := new(Ui).
		Init(services, theme).
		CreateGrid().
		AddStatusControls().
		StyleCustomization(theme)

	ui.SearchView.AddSearchControls()
	ui.SearchView.AddObserver(ui.SecretsView)

	ui.KeyvaultView.AddCurrentKeyvaultWatcher(ui.SecretsView)
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
