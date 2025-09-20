package keyvault

import (
	"github.com/declan-whiting/vaulty/internal/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CacheService interface {
	ReadKeyvaults() []models.KeyvaultModel
}

type ConfigrationService interface {
	GetConfiguration() models.ConfigurationList
}

type CurrentKeyvaultWatcher interface {
	CurrentKeyvaultUpdated(name, subscription string)
}

type KeyvaultView struct {
	*tview.List
	Conf                    ConfigrationService
	Cache                   CacheService
	QuitHandler             func()
	SearchHandler           func()
	SelectedHandler         func()
	CurrentKeyvaultWatchers []CurrentKeyvaultWatcher
}

func NewKeyvaultView(cache CacheService, conf ConfigrationService, quiter, searcher, selecter func()) *KeyvaultView {
	keyvaultView := &KeyvaultView{
		Cache:           cache,
		Conf:            conf,
		QuitHandler:     quiter,
		SearchHandler:   searcher,
		SelectedHandler: selecter,
	}
	keyvaultView.List = tview.NewList()
	keyvaultView.SetTitle("Keyvaults")
	keyvaultView.SetBorder(true)
	keyvaultView.ShowSecondaryText(false)
	keyvaultView.SetBorderPadding(0, 0, 1, 0)

	for _, v := range cache.ReadKeyvaults() {
		keyvaultView.AddItem(v.Name, v.SubscriptionId, rune(0), nil)
	}

	keyvaultView.AddKeyvaultViewControls()
	keyvaultView.KeyvaultSelectionChanged()

	return keyvaultView
}

func (kv *KeyvaultView) AddKeyvaultViewControls() {
	kv.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			kv.QuitHandler()
			return tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone)
		}
		if event.Rune() == '/' {
			kv.SearchHandler()
		}
		if event.Key() == tcell.KeyEnter {
			kv.SelectedHandler()
			return nil
		}
		return event
	})

}

func (kv *KeyvaultView) SetInitialKeyvault() {
	vault := kv.Conf.GetConfiguration().Keyvaults[0]

	for _, v := range kv.CurrentKeyvaultWatchers {
		v.CurrentKeyvaultUpdated(vault.Name, vault.SubscriptionId)
	}
}

func (kv *KeyvaultView) KeyvaultSelectionChanged() {
	kv.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		for _, v := range kv.CurrentKeyvaultWatchers {
			v.CurrentKeyvaultUpdated(mainText, secondaryText)
		}
	})

}

func (kv *KeyvaultView) AddCurrentKeyvaultWatcher(watcher CurrentKeyvaultWatcher) {
	kv.CurrentKeyvaultWatchers = append(kv.CurrentKeyvaultWatchers, watcher)
}
