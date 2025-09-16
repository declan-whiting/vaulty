package models

type KeyvaultModel struct {
	ID             string        `json:"id"`
	Location       string        `json:"location"`
	Name           string        `json:"name"`
	ResourceGroup  string        `json:"resourceGroup"`
	Properties     Properties    `json:"properties"`
	SubscriptionId string        // appended from cache
	Secrets        []SecretModel //appended from cache or queried from azure
}

type Properties struct {
	ResourceGroup                string `json:"resourceGroup"`
	CreateMode                   bool   `json:"createMode"`
	EnablePurgeProtection        bool   `json:"enablePurgeProtection"`
	EnableRbacAuthorization      bool   `json:"enableRbacAuthorization"`
	EnableSoftDelete             bool   `json:"enableSoftDelete"`
	EnabledForDeployment         bool   `json:"enabledForDeployment"`
	EnabledForDiskEncryption     bool   `json:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment bool   `json:"enabledForTemplateDeployment"`
	HsmPoolResourceId            string `json:"hsmPoolResourceId"`
	ProvisioningState            string `json:"provisioningState"`
	PublicNetworkAccess          string `json:"publicNetworkAccess"`
	Sku                          Sku    `json:"sku"`
}

type Sku struct {
	Family string `json:"family"`
	Name   string `json:"name"`
}
