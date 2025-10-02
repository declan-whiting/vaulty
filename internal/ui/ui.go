package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/declan-whiting/vaulty/internal/components/controls"
	"github.com/declan-whiting/vaulty/internal/components/keyvault"
	"github.com/declan-whiting/vaulty/internal/components/search"
	"github.com/declan-whiting/vaulty/internal/components/secretDetails"
	"github.com/declan-whiting/vaulty/internal/components/secrets"
	"github.com/declan-whiting/vaulty/internal/components/status"
	"github.com/declan-whiting/vaulty/internal/events"
	"github.com/declan-whiting/vaulty/internal/theme"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Ui struct {
	App          *tview.Application
	Grid         *tview.Grid
	ControlsView *controls.ControlsView
	KeyvaultView *keyvault.KeyvaultView
	SecretsView  *secrets.SecretsView
	SearchView   *search.SearchView
	StatusView   *status.StatusView
	Services     *Services
	Events       *events.EventStore
}

func BuildUi() {
	start := time.Now()
	services := Services{}
	services.Init()

	theme := theme.NewTheme()
	tview.Styles = theme.GetTheme()

	ui := new(Ui).Init(services, theme).CreateGrid()

	ui.SearchView.AddObserver(ui.SecretsView)

	ui.KeyvaultView.AddCurrentKeyvaultWatcher(ui.SecretsView)
	ui.KeyvaultView.SetInitialKeyvault()

	ui.Events.AddNewEventObserver(ui.StatusView)
	events.TimedEventLog(start, "\U0001F308 UI Built Sucessfully", *ui.Events)

	ui.App.SetRoot(ui.Grid, true)
	err := ui.App.SetFocus(ui.KeyvaultView).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (ui *Ui) Init(services Services, themer Themer) *Ui {
	ui.App = tview.NewApplication()
	ui.Events = events.NewEventStore()
	ui.Services = &services

	ui.SecretsView = secrets.NewSecretsView(
		services.CacheService,
		ui.HandleQuit,
		ui.HandleSearch,
		ui.Handleback,
		ui.HandleSecretsSelectedChanged,
		themer)

	ui.KeyvaultView = keyvault.NewKeyvaultView(
		services.CacheService,
		services.ConfigrationService,
		ui.HandleQuit,
		ui.HandleSearch,
		ui.FocusSecretsView, themer)

	ui.ControlsView = controls.NewControlsView(themer)
	ui.SearchView = search.NewSearchView(ui.FocusSecretsView, ui.EscapeSearch, themer)
	ui.StatusView = status.NewStatusView(themer)

	return ui
}

func (ui *Ui) EscapeSearch() {
	ui.App.SetFocus(ui.KeyvaultView)
	ui.HideSearch()
}

func (ui *Ui) HandleQuit() {
	ui.App.Stop()
}

func (ui *Ui) HandleSearch() {
	ui.ShowSearch()
	ui.SearchView.SetText("")
	ui.SecretsView.ScrollToBeginning()
	ui.App.SetFocus(ui.SearchView)
	ui.SecretsView.Select(0, 0)
}

func (ui *Ui) FocusSecretsView() {
	ui.App.SetFocus(ui.SecretsView)
	ui.HideSearch()
}

func (ui *Ui) Handleback() {
	ui.App.SetFocus(ui.KeyvaultView)
}

func (ui *Ui) CloseSecretDetailsView(view *tview.TextView) {
	ui.Grid.RemoveItem(view)
	ui.Grid.AddItem(ui.SecretsView, 1, 1, 3, 2, 0, 0, false)
	ui.App.SetFocus(ui.SecretsView)
	ui.SearchView.SetText("")
}

func (ui *Ui) HandleSecretsSelectedChanged(secret, keyvault, subscription string) {
	ui.Grid.RemoveItem(ui.SecretsView)
	defer events.TimedEventLog(time.Now(), fmt.Sprintf("\U0001F916 Got %s ", secret), *ui.Events)

	secretText := ui.Services.AzureService.AzShowSecret(secret, keyvault, subscription)
	details := secretDetails.CreateSecretsDetailView(fmt.Sprintf("%s/%s", keyvault, secret), secretText)
	ui.App.SetFocus(details)
	details.AddControls(ui.CloseSecretDetailsView)

	ui.Grid.AddItem(details, 1, 1, 3, 2, 0, 0, false)
}

type Themer interface {
	GetColor(color string) tcell.Color
	SetTableCellTheme(table *tview.Table, row int, col int, foreground, background string)
}
