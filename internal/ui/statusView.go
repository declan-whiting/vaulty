package ui

import (
	"fmt"
	"sync"
	"time"

	"github.com/declan-whiting/vaulty/internal/azure"
	"github.com/declan-whiting/vaulty/internal/cache"
	"github.com/declan-whiting/vaulty/internal/configuration"
	"github.com/declan-whiting/vaulty/internal/models"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewStatusView() *tview.TextView {
	statusView := tview.NewTextView()
	statusView.SetText(cache.ReadLastSync()).SetTextAlign(tview.AlignCenter)
	statusView.SetBorder(true)
	statusView.SetTitle("Status")

	return statusView
}

func (ui *Ui) AddStatusControls() *Ui {
	ui.Grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlR) {
			loading := true

			go func() {
				frames := []string{"[   ]", "[=  ]", "[== ]", "[===]", "[ ==]", "[  =]", "[   ]"}
				i := 0
				for {
					if loading {
						ui.StatusView.SetText(fmt.Sprintf("Loading %s", frames[i%len(frames)]))
						time.Sleep(150 * time.Millisecond)
						i++
						ui.App.Draw()
					}
				}
			}()

			go func() {
				ui.StatusView.SetText(UpdateFromAzure())
				loading = false
				ui.App.Draw()

			}()
			return event
		}

		return event
	})

	return ui
}

func UpdateFromAzure() string {
	var vaults []models.KeyvaultModel
	start := time.Now()

	config := configuration.GetConfiguration()

	for _, v := range config.Keyvaults {
		vaults = append(vaults, azure.AzShowKeyvault(v.Name, v.SubscriptionId))
	}

	var wg sync.WaitGroup
	wg.Add(len(vaults))
	for i, v := range vaults {
		go func() {
			defer wg.Done()
			vaults[i].Secrets = azure.AzGetSecrets(v.Name, v.SubscriptionId)
		}()
	}
	wg.Wait()

	elapsed := time.Since(start)
	took := fmt.Sprintf("\nTook: %.1fs", elapsed.Seconds())
	return (start.Format("Last Sync: "+time.ANSIC) + took)
}
