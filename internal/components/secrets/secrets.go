package secrets

import (
	"strings"

	"github.com/declan-whiting/vaulty/internal/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type SecretsView struct {
	*tview.Table
	Cache                        CacheService
	QuitHandler                  func()
	SearchHandler                func()
	BackHandler                  func()
	SelectedSecretChangedHandler func(secret, keyvault, subscription string)
	CurrentKeyvaultName          string
	CurrentKeyvaultSubscription  string
}

type CacheService interface {
	ReadSecrets(keyvault string) []models.SecretModel
}

type Themer interface {
	GetColor(color string) tcell.Color
	SetTableCellTheme(table *tview.Table, row int, col int, foreground, background string)
}

func NewSecretsView(cache CacheService,
	quitHandler, searchHandler, backHandler func(),
	selectedChangedHandler func(string, string, string), theme Themer) *SecretsView {
	secretsView := &SecretsView{}
	secretsView.Cache = cache
	secretsView.Table = tview.NewTable()
	secretsView.SetSelectable(true, false)
	secretsView.SetBorder(true)
	secretsView.SetTitle("Secrets")
	secretsView.QuitHandler = quitHandler
	secretsView.SearchHandler = searchHandler
	secretsView.BackHandler = backHandler
	secretsView.SelectedSecretChangedHandler = selectedChangedHandler
	secretsView.SetSelectable(true, false)

	secretsView.SetSelectedStyle(tcell.StyleDefault.
		Background(theme.GetColor("background")).
		Foreground(theme.GetColor("pink")))

	secretsView.AddSecretsControls()
	secretsView.SecretSelectedChanged()
	return secretsView
}

func (sv *SecretsView) AddSecretsControls() {
	sv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			sv.QuitHandler()
			return tcell.NewEventKey(tcell.KeyRune, 's', tcell.ModNone)
		}
		if event.Rune() == '/' {

			sv.SearchHandler()
		}
		if event.Rune() == 'b' {

			sv.BackHandler()
			return tcell.NewEventKey(tcell.KeyRune, 'b', tcell.ModNone)
		}
		return event
	})
}

func (sv *SecretsView) SecretSelectedChanged() {
	sv.SetSelectedFunc(func(row, column int) {
		secretName := sv.GetCell(row, column).Text
		sv.SelectedSecretChangedHandler(secretName, sv.CurrentKeyvaultName, sv.CurrentKeyvaultSubscription)
	})
}

func (sv *SecretsView) NotifyUpdate(content string) {
	sv.Clear()
	i := 0
	for _, v := range sv.Cache.ReadSecrets(sv.CurrentKeyvaultName) {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(content)) {
			sv.SetCell(i, 0, tview.NewTableCell(v.Name))
			i++
		}

	}
}

func (sv *SecretsView) CurrentKeyvaultUpdated(name, subscription string) {
	sv.CurrentKeyvaultName = name
	sv.CurrentKeyvaultSubscription = subscription
	sv.SetTitle(name + "/secrets")
	for i, v := range sv.Cache.ReadSecrets(name) {
		sv.SetCell(i, 0, tview.NewTableCell(v.Name))
	}
}
