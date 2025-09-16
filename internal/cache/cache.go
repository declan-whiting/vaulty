package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/declan-whiting/vaulty/internal/configuration"
	"github.com/declan-whiting/vaulty/internal/models"
)

func getKeyvaultFilePath(name string) string {
	path := "bin/cache/" + name + "-kv.json"
	return path
}

func getSecretsFilePath(name string) string {
	path := "bin/cache/" + name + "-secrets.json"
	return path
}

func getLastSyncPath() string {
	path := "bin/cache/lastsync.txt"
	return path
}
func WriteLastSync(contents []byte) {
	fileName := getLastSyncPath()
	err := os.WriteFile(fileName, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteKeyvault(name string, contents []byte) {
	fileName := getKeyvaultFilePath(name)
	err := os.WriteFile(fileName, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteSecrets(name string, contents []byte) {
	fileName := getSecretsFilePath(name)
	err := os.WriteFile(fileName, contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func EnsureCache() {
	path := filepath.Join(".", "bin/cache/")
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadKeyvaults() []models.KeyvaultModel {
	EnsureCache()
	config := configuration.GetConfiguration()
	var cachedVaults []models.KeyvaultModel

	for i, v := range config.Keyvaults {
		if _, err := os.Stat(getKeyvaultFilePath(v.Name)); errors.Is(err, os.ErrNotExist) {
			return nil
		} else {
			cacheVaultFile, err := os.Open(getKeyvaultFilePath(v.Name))
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

func ReadSecrets(keyvaultName string) []models.SecretModel {
	var cachedSecrets []models.SecretModel
	if _, err := os.Stat(getSecretsFilePath(keyvaultName)); errors.Is(err, os.ErrNotExist) {
		return nil
	} else {
		cachedSecretsFile, err := os.Open(getSecretsFilePath(keyvaultName))
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

func ReadLastSync() string {
	if _, err := os.Stat(getLastSyncPath()); errors.Is(err, os.ErrNotExist) {
		panic("No cache for last sync")
	} else {
		lastSyncFile, err := os.Open(getLastSyncPath())
		if err != nil {
			fmt.Println(err)
		}

		out, _ := io.ReadAll(lastSyncFile)
		return string(out)
	}
}
