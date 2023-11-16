package common

import (
	"errors"
	"fmt"

	"os"
	"path/filepath"
	"strings"

	"api/ipatool/pkg/appstore"
	"api/ipatool/pkg/http"
	"api/ipatool/pkg/keychain"
	"api/ipatool/pkg/log"
	"api/ipatool/pkg/util"
	"api/ipatool/pkg/util/machine"
	"api/ipatool/pkg/util/operatingsystem"
	"github.com/99designs/keyring"
	cookiejar "github.com/juju/persistent-cookiejar"

	"golang.org/x/term"
)

const (
	ConfigDirectoryName = ".ipatool"
	CookieJarFileName   = "cookies"
	KeychainServiceName = "ipatool-auth.service"
)

type Commontem interface {
}

var Dependenciexs = Dependencies{}
var keychainPassphrase string

type Dependencies struct {
	Logger    log.Logger
	OS        operatingsystem.OperatingSystem
	Machine   machine.Machine
	CookieJar http.CookieJar
	Keychain  keychain.Keychain
	AppStore  appstore.AppStore
}

// newCookieJar returns a new cookie jar instance.
func newCookieJar(machine machine.Machine) http.CookieJar {
	return util.Must(cookiejar.New(&cookiejar.Options{
		Filename: filepath.Join(machine.HomeDirectory(), ConfigDirectoryName, CookieJarFileName),
	}))
}

// newKeychain returns a new keychain instance.
func newKeychain(machine machine.Machine, interactive bool) keychain.Keychain {
	ring := util.Must(keyring.Open(keyring.Config{
		AllowedBackends: []keyring.BackendType{
			keyring.KeychainBackend,
			keyring.SecretServiceBackend,
			keyring.FileBackend,
		},
		ServiceName: KeychainServiceName,
		FileDir:     filepath.Join(machine.HomeDirectory(), ConfigDirectoryName),
		FilePasswordFunc: func(s string) (string, error) {
			if keychainPassphrase == "" && !interactive {
				return "", errors.New("keychain passphrase is required when not running in interactive mode; use the \"--keychain-passphrase\" flag")
			}

			if keychainPassphrase != "" {
				return keychainPassphrase, nil
			}

			path := strings.Split(s, " unlock ")[1]
			fmt.Printf("enter passphrase to unlock %s (this is separate from your Apple ID password): \n", path)
			bytes, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				return "", fmt.Errorf("failed to read password: %w", err)
			}

			password := string(bytes)
			password = strings.Trim(password, "\n")
			password = strings.Trim(password, "\r")

			return password, nil
		},
	}))

	return keychain.New(keychain.Args{Keyring: ring})
}

// initWithCommand initializes the dependencies of the command.
func InitWithCommand() {

	Dependenciexs.OS = operatingsystem.New()
	Dependenciexs.Machine = machine.New(machine.Args{OS: Dependenciexs.OS})
	Dependenciexs.CookieJar = newCookieJar(Dependenciexs.Machine)
	Dependenciexs.Keychain = newKeychain(Dependenciexs.Machine, false)
	Dependenciexs.AppStore = appstore.NewAppStore(appstore.Args{
		CookieJar:       Dependenciexs.CookieJar,
		OperatingSystem: Dependenciexs.OS,
		Keychain:        Dependenciexs.Keychain,
		Machine:         Dependenciexs.Machine,
	})

	util.Must("", createConfigDirectory(Dependenciexs.OS, Dependenciexs.Machine))
}

// createConfigDirectory creates the configuration directory for the CLI tool, if needed.
func createConfigDirectory(os operatingsystem.OperatingSystem, machine machine.Machine) error {
	configDirectoryPath := filepath.Join(machine.HomeDirectory(), ConfigDirectoryName)
	_, err := os.Stat(configDirectoryPath)

	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(configDirectoryPath, 0700)
		if err != nil {
			return fmt.Errorf("failed to create config directory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("could not read metadata: %w", err)
	}

	return nil
}
