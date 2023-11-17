package appstore

import (
	"api/ipatool/pkg/http"
	"api/ipatool/pkg/keychain"
	"api/ipatool/pkg/util/machine"
	"api/ipatool/pkg/util/operatingsystem"
)

type AppStore interface {
	// Login authenticates with the App Store.
	Login(input LoginInput) (LoginOutput, error)
	// AccountInfo returns the information of the authenticated account.
	AccountInfo() (AccountInfoOutput, error)
	// NetAccountInfo returns the information of the authenticated account.
	NetAccountInfo() (AccountInfoOutput, error)
	// Revoke revokes the active credentials.
	Revoke() error
	// Lookup looks apps up based on the specified bundle identifier.
	Lookup(input LookupInput) (LookupOutput, error)
	// Search searches the App Store for apps matching the specified term.
	Search(input SearchInput) (SearchOutput, error)

	GetCountryCode(storeFront string) (string, error)
	// Purchase acquires a license for the desired app.
	// Note: only free apps are supported.
	Purchase(input PurchaseInput) error
	// Download downloads the IPA package from the App Store to the desired location.
	Download(input DownloadInput) (DownloadOutput, error)
	// ReplicateSinf replicates the sinf for the IPA package.
	ReplicateSinf(input ReplicateSinfInput) error
}

type appstore struct {
	keychain       keychain.Keychain
	loginClient    http.Client[loginResult]
	searchClient   http.Client[searchResult]
	purchaseClient http.Client[purchaseResult]
	downloadClient http.Client[downloadResult]
	netInfoClient  http.Client[loginResult]
	httpClient     http.Client[interface{}]
	machine        machine.Machine
	os             operatingsystem.OperatingSystem
}

type Args struct {
	Keychain        keychain.Keychain
	CookieJar       http.CookieJar
	OperatingSystem operatingsystem.OperatingSystem
	Machine         machine.Machine
}

func NewAppStore(args Args) AppStore {
	clientArgs := http.Args{
		CookieJar: args.CookieJar,
	}

	return &appstore{
		keychain:       args.Keychain,
		loginClient:    http.NewClient[loginResult](clientArgs),
		searchClient:   http.NewClient[searchResult](clientArgs),
		purchaseClient: http.NewClient[purchaseResult](clientArgs),
		downloadClient: http.NewClient[downloadResult](clientArgs),
		netInfoClient:  http.NewClient[loginResult](clientArgs),
		httpClient:     http.NewClient[interface{}](clientArgs),
		machine:        args.Machine,
		os:             args.OperatingSystem,
	}
}
