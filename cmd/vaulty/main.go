package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/declan-whiting/vaulty/internal/azure"
	"github.com/declan-whiting/vaulty/internal/cache"
	"github.com/declan-whiting/vaulty/internal/configuration"
	"github.com/declan-whiting/vaulty/internal/models"
	"github.com/declan-whiting/vaulty/internal/ui"
)

func main() {
	startUp()
	ui.BuildUi()
}

func startUp() []models.KeyvaultModel {
	fmt.Printf("Starting vaulty...\n")
	fmt.Println("Attempting read from cache..")
	vaults := cache.ReadKeyvaults()

	if vaults == nil {
		config := configuration.GetConfiguration()
		fmt.Printf("No cache found, reading configuration...\n")
		for _, v := range config.Keyvaults {
			fmt.Printf("Getting %s from Azure...\n", v.Name)
			vaults = append(vaults, azure.AzShowKeyvault(v.Name, v.SubscriptionId))
		}

		writeLastSync()
	}

	var wg sync.WaitGroup
	for i, v := range vaults {
		fmt.Printf("Getting secrets...\n")
		if secrets := cache.ReadSecrets(v.Name); secrets != nil {
			fmt.Printf("Reading secrets from cache\n")
			vaults[i].Secrets = secrets
		} else {
			fmt.Printf("No secrets cache found for %s, reading from azure...\n", v.Name)

			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Printf("Gettings secrets for %s...\n", v.Name)
				vaults[i].Secrets = azure.AzGetSecrets(v.Name, v.SubscriptionId)
				fmt.Printf("Got secrets for %s...\n", v.Name)
			}()
		}
	}
	wg.Wait()

	writeLastSync()
	fmt.Printf("Finished loading!")
	return vaults
}

func writeLastSync() {
	bytes := []byte(fmt.Sprintf("Last Sync: %s", time.Now().Format(time.ANSIC)))
	cache.WriteLastSync(bytes)
}
