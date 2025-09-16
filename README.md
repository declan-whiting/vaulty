# Vaulty

[![Go Report Card](https://goreportcard.com/badge/github.com/declan-whiting/vaulty)](https://goreportcard.com/report/github.com/declan-whiting/vaulty)

An Azure Keyvault TUI, written in Golang using [tview](https://github.com/rivo/tview/tree/master). 

![Screenshot](vaulty.gif)

Vaulty is under active development and is subject to change. 

## Building from source

1. Clone the repository
2. Update vaulty.conf
   - You must configure at least one Keyvault. Remove any unused Keyvaults from configuration. 

``` yaml
Keyvaults:
  - Name: <keyvault-name>
    Subscription: <keyvault-subscription-id>
  - Name: <keyvault-name>
    Subscription: <keyvault-subscription-id>
  - Name: <keyvault-name>
    Subscription: <keyvault-subscription-id>
```
3. Build and run the executable

``` bash
cd vaulty
make && ./bin/vaulty
```

# Roadmap
- Bug fixing
- Global search
- Secret modification/deletion
- Certificates management
- Keys mananagement
- Themes






