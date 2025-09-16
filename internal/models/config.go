package models

type ConfigurationList struct {
	Keyvaults []KeyvaultConfiguration `yaml:"Keyvaults"`
}

type KeyvaultConfiguration struct {
	Name           string `yaml:"Name"`
	SubscriptionId string `yaml:"Subscription"`
}
