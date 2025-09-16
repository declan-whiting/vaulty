package ui

import (
	"github.com/declan-whiting/vaulty/internal/azure"
	"github.com/declan-whiting/vaulty/internal/cache"
	"github.com/declan-whiting/vaulty/internal/configuration"
	"github.com/declan-whiting/vaulty/internal/models"
)

type ConfigurationService interface {
	GetConfiguration() models.ConfigurationList
}

type CacheService interface {
	ReadKeyvaults() []models.KeyvaultModel
	ReadSecrets(string) []models.SecretModel
	WriteLastSync([]byte)
	ReadLastSync() string
	WriteKeyvault(string, []byte)
	WriteSecrets(string, []byte)
}

type AzureService interface {
	AzShowKeyvault(string, string) models.KeyvaultModel
	AzGetSecrets(string, string) []models.SecretModel
	AzShowSecret(string, string, string) string
}

type Services struct {
	AzureService        AzureService
	CacheService        CacheService
	ConfigrationService ConfigurationService
}

func (s *Services) Init() {
	s.ConfigrationService = configuration.NewConfigurationService()
	s.CacheService = cache.NewCacheService(s.ConfigrationService)
	s.AzureService = azure.NewAzureService(s.CacheService)
}
