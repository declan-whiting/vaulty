package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/declan-whiting/vaulty/internal/models"
)

type ConfigurationService interface {
	GetConfiguration() models.ConfigurationList
}

type CacheService struct {
	config ConfigurationService
}

func NewCacheService(config ConfigurationService) *CacheService {
	cs := new(CacheService)
	cs.config = config
	return cs
}

func (cs *CacheService) getKeyvaultFilePath(name string) string {
	path := "bin/cache/" + name + "-kv.json"
	return path
}

func (cs *CacheService) getSecretsFilePath(name string) string {
	path := "bin/cache/" + name + "-secrets.json"
	return path
}

func (cs *CacheService) getLastSyncPath() string {
	path := "bin/cache/lastsync.txt"
	return path
}
func (cs *CacheService) WriteLastSync(contents []byte) {
	fileName := cs.getLastSyncPath()
	err := os.WriteFile(fileName, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (cs *CacheService) WriteKeyvault(name string, contents []byte) {
	fileName := cs.getKeyvaultFilePath(name)
	err := os.WriteFile(fileName, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (cs *CacheService) WriteSecrets(name string, contents []byte) {
	fileName := cs.getSecretsFilePath(name)
	err := os.WriteFile(fileName, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (cs *CacheService) EnsureCache() {
	path := filepath.Join(".", "bin/cache/")
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func (cs *CacheService) ReadKeyvaults() []models.KeyvaultModel {
	cs.EnsureCache()
	config := cs.config.GetConfiguration()
	var cachedVaults []models.KeyvaultModel

	for i, v := range config.Keyvaults {
		if _, err := os.Stat(cs.getKeyvaultFilePath(v.Name)); errors.Is(err, os.ErrNotExist) {
			return nil
		} else {
			cacheVaultFile, err := os.Open(cs.getKeyvaultFilePath(v.Name))
			if err != nil {
				log.Fatal(err)
			}
			var vault models.KeyvaultModel
			out, _ := io.ReadAll(cacheVaultFile)

			err = json.Unmarshal(out, &vault)
			if err != nil {
				log.Fatal(err)
			}
			vault.SubscriptionId = config.Keyvaults[i].SubscriptionId
			cachedVaults = append(cachedVaults, vault)

		}
	}
	return cachedVaults
}

func (cs *CacheService) ReadSecrets(keyvaultName string) []models.SecretModel {
	var cachedSecrets []models.SecretModel
	if _, err := os.Stat(cs.getSecretsFilePath(keyvaultName)); errors.Is(err, os.ErrNotExist) {
		return nil
	} else {
		cachedSecretsFile, err := os.Open(cs.getSecretsFilePath(keyvaultName))
		if err != nil {
			fmt.Println(err)
		}
		var secret []models.SecretModel

		out, _ := io.ReadAll(cachedSecretsFile)

		if err := json.Unmarshal(out, &secret); err != nil {
			log.Fatal(err)
		}

		cachedSecrets = secret

	}

	return cachedSecrets
}

func (cs *CacheService) ReadLastSync() string {
	if _, err := os.Stat(cs.getLastSyncPath()); errors.Is(err, os.ErrNotExist) {
		panic("No cache for last sync")
	} else {
		lastSyncFile, err := os.Open(cs.getLastSyncPath())
		if err != nil {
			fmt.Println(err)
		}

		out, _ := io.ReadAll(lastSyncFile)
		return string(out)
	}
}
