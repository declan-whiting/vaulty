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
	configService := configuration.NewConfigurationService()
	cacheService := cache.NewCacheService(configService)
	azureService := azure.NewAzureService(cacheService)

	startUp(configService, cacheService, azureService)

	ui.BuildUi()
}

type ConfigurationService interface {
	GetConfiguration() models.ConfigurationList
}

type CacheService interface {
	ReadKeyvaults() []models.KeyvaultModel
	ReadSecrets(string) []models.SecretModel
	WriteLastSync([]byte)
}

type AzureService interface {
	AzShowKeyvault(string, string) models.KeyvaultModel
	AzGetSecrets(string, string) []models.SecretModel
}

func startUp(conf ConfigurationService, cache CacheService, azure AzureService) []models.KeyvaultModel {
	fmt.Printf("Starting vaulty...\n")
	fmt.Println("Attempting read from cache..")
	vaults := cache.ReadKeyvaults()

	if vaults == nil {
		config := conf.GetConfiguration()
		fmt.Printf("No cache found, reading configuration...\n")
		for _, v := range config.Keyvaults {
			fmt.Printf("Getting %s from Azure...\n", v.Name)
			vaults = append(vaults, azure.AzShowKeyvault(v.Name, v.SubscriptionId))
		}

		writeLastSync(cache)
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

	writeLastSync(cache)
	fmt.Printf("Finished loading!")
	return vaults
}

func writeLastSync(cs CacheService) {
	bytes := []byte(fmt.Sprintf("Last Sync: %s", time.Now().Format(time.ANSIC)))
	cs.WriteLastSync(bytes)
}
