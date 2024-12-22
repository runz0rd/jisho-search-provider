package provider

import (
	"fmt"
	"log/slog"

	"github.com/godbus/dbus/v5"
)

// https://developer.gnome.org/documentation/tutorials/search-provider.html
type Provider interface {
	GetInitialResultSet(terms []string) ([]string, *dbus.Error)
	GetSubsearchResultSet(previousResults, terms []string) ([]string, *dbus.Error)
	GetResultMetas(identifiers []string) ([]map[string]dbus.Variant, *dbus.Error)
	ActivateResult(identifier string, terms []string, timestamp uint32) *dbus.Error
	LaunchSearch(terms []string, timestamp uint32) *dbus.Error
}

const searchProviderInterface = "org.gnome.Shell.SearchProvider2"

func ExportProvider(provider Provider, conn *dbus.Conn, busName, objectPath string) error {
	reply, err := conn.RequestName(busName, dbus.NameFlagDoNotQueue)

	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		return fmt.Errorf("Failed to request name: %v", err)
	}

	slog.Info("dbus", "reply", reply)
	return conn.Export(provider, dbus.ObjectPath(objectPath), searchProviderInterface)
}
