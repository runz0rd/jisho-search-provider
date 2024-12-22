package main

import (
	"log/slog"

	"github.com/godbus/dbus/v5"
	"github.com/runz0rd/jisho-search-provider/gnome/search/provider"
	"github.com/runz0rd/jisho-search-provider/gnome/search/provider/jisho"
	"github.com/runz0rd/jisho-search-provider/jisho/api"
)

func main() {
	// TODO Received error from D-Bus search provider app.desktop: Gio.DBusError: \
	// GDBus.Error:org.freedesktop.DBus.Error.ServiceUnknown: The name is not activatable
	conn, err := dbus.SessionBus()
	if err != nil {
		slog.Error("failed to connect to session bus", "error", err)
		panic(err)
	}
	defer conn.Close()

	jsp := jisho.New(api.NewClient("https://jisho.org/api/v1"))
	if err = provider.ExportProvider(jsp, conn, jisho.BusName, jisho.ObjectPath); err != nil {
		slog.Error("failed to export jisho search provider", "error", err)
		panic(err)
	}
	slog.Info("jisho search provider running")
	select {}
}
