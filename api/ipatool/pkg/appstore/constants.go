package appstore

const (
	FailureTypeInvalidCredentials     = "-5000"
	FailureTypePasswordTokenExpired   = "2034"
	FailureTypeLicenseNotFound        = "9610"
	FailureTypeTemporarilyUnavailable = "2059"

	CustomerMessageBadLogin             = "MZFinance.BadLogin.Configurator_message"
	CustomerMessageSubscriptionRequired = "Subscription Required"

	iTunesAPIDomain     = "itunes.apple.com"
	iTunesAPIPathSearch = "/search"
	iTunesAPIPathLookup = "/lookup"

	PrivateAppStoreAPIDomainPrefixWithoutAuthCode = "p25"
	PrivateAppStoreAPIDomainPrefixWithAuthCode    = "p71"
	PrivateAppStoreAPIDomainPrefixWithAuthCode1   = "p73"
	PrivateAppStoreAPIDomain                      = "buy." + iTunesAPIDomain
	PrivateAppStoreAPIPathAuthenticate            = "/WebObjects/MZFinance.woa/wa/authenticate"
	PrivateAppStoreAPIPathPurchase                = "/WebObjects/MZBuy.woa/wa/buyProduct"
	PrivateAppStoreAPIPathDownload                = "/WebObjects/MZFinance.woa/wa/volumeStoreDownloadProduct"
	PrivateAppStoreAPIAccountInfo                 = "/WebObjects/MZFinance.woa/wa/accountSummary" ///WebObjects/MZFinance.woa/wa/com.apple.jingle.app.finance.DirectAction/accountSummary

	//PrivateAppStoreAPIAccountInfo = "/WebObjects/MZFinance.woa/wa/com.apple.jingle.app.finance.DirectAction/accountSummary"

	HTTPHeaderStoreFront = "X-Set-Apple-Store-Front"

	PricingParameterAppStore    = "STDQ"
	PricingParameterAppleArcade = "GAME"
)
