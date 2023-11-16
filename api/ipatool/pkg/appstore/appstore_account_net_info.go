package appstore

import (
	"api/ipatool/pkg/http"
	"encoding/json"
	"fmt"
	"strings"
)

func (t *appstore) NetAccountInfo() (AccountInfoOutput, error) {
	data, err := t.keychain.Get("account")
	if err != nil {
		return AccountInfoOutput{}, fmt.Errorf("failed to get account: %w", err)
	}

	var acc Account

	err = json.Unmarshal(data, &acc)

	if err != nil {
		return AccountInfoOutput{}, fmt.Errorf("failed to unmarshal json: %w", err)
	}
	macAddr, err := t.machine.MacAddress()
	if err != nil {
		return AccountInfoOutput{}, fmt.Errorf("failed to get mac address: %w", err)
	}

	guid := strings.ReplaceAll(strings.ToUpper(macAddr), ":", "")

	req := t.netAccountInfoRequest(acc, guid)

	res, err := t.netInfoClient.Send(req)

	if err != nil {
		return AccountInfoOutput{}, fmt.Errorf("request failed: %w", err)
	}

	fmt.Println("StatusCode:", res.StatusCode)

	return AccountInfoOutput{
		Account: acc,
	}, nil
}

func (*appstore) netAccountInfoRequest(acc Account, guid string) http.Request {
	host := fmt.Sprintf("%s-%s", PrivateAppStoreAPIDomainPrefixWithAuthCode1, PrivateAppStoreAPIDomain)

	return http.Request{
		URL:            fmt.Sprintf("https://%s%s?guid=%s", host, PrivateAppStoreAPIAccountInfo, guid),
		Method:         http.MethodGET,
		ResponseFormat: http.ResponseFormatXML,
		Headers: map[string]string{
			"Accept":              "*/*",
			"Content-Type":        "text/html,application/xhtml+xml,application/xml",
			"iCloud-DSID":         acc.DirectoryServicesID,
			"X-Dsid":              acc.DirectoryServicesID,
			"X-Token":             acc.PasswordToken,
			"X-Apple-Store-Front": acc.StoreFront,
		},
	}
}
