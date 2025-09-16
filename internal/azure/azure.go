package azure

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/declan-whiting/vaulty/internal/models"
)

type CacheService interface {
	WriteKeyvault(string, []byte)
	WriteSecrets(string, []byte)
}

type AzureService struct {
	CacheService CacheService
}

func NewAzureService(cache CacheService) *AzureService {
	azure := new(AzureService)
	azure.CacheService = cache
	return azure
}

func (az *AzureService) AzShowKeyvault(name, subscriptionId string) models.KeyvaultModel {
	out, _ := exec.Command("az", "keyvault", "show", "--name", name, "--subscription", subscriptionId, "--output", "json").CombinedOutput()
	az.CacheService.WriteKeyvault(name, out)
	var kv models.KeyvaultModel
	kv.SubscriptionId = subscriptionId
	err := json.Unmarshal(out, &kv)
	if err != nil {
		fmt.Println("Failed to parse JSON for keyvaults")
		fmt.Println(err)
	}

	return kv
}

func (az *AzureService) AzGetSecrets(name, subscriptionId string) []models.SecretModel {
	out, _ := exec.Command("az", "keyvault", "secret", "list", "--vault-name", name, "--subscription", subscriptionId, "--output", "json").CombinedOutput()
	az.CacheService.WriteSecrets(name, out)
	var response []models.SecretModel
	err := json.Unmarshal(out, &response)
	if err != nil {
		fmt.Println("Failed to parse JSON for secrets")
		fmt.Println(err)
	}

	return response
}

func (az *AzureService) AzShowSecret(secretName, vaultName, subscriptionId string) string {
	out, _ := exec.Command("az", "keyvault", "secret", "show", "--vault-name", vaultName, "--name", secretName, "--subscription", subscriptionId, "--output", "json").CombinedOutput()
	return string(out)
}
