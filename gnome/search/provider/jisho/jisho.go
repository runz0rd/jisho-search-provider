package jisho

import (
	"fmt"
	"log/slog"

	"github.com/godbus/dbus/v5"
	"github.com/runz0rd/jisho-search-provider/jisho/api"
)

const BusName = "org.jisho.SearchProvider"
const ObjectPath = "/org/jisho/SearchProvider"

type Provider struct {
	client *api.Client
}

func New(client *api.Client) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) GetInitialResultSet(terms []string) ([]string, *dbus.Error) {
	slog.Info("GetInitialResultSet", "terms", terms)
	// TODO fix
	// ctx, cancel := context.WithCancel(context.Background())
	// select {
	// case <-ctx.Done():
	// 	return []string{"Thinking..."}, nil
	// case <-time.After(sp.debounceTime):
	// 	return []string{sp.performSearch(terms)}, nil
	// }
	return nil, nil
}

func (p *Provider) GetSubsearchResultSet(previousResults, terms []string) ([]string, *dbus.Error) {
	slog.Info("GetSubsearchResultSet", "terms", terms)
	result, err := p.client.Search(terms[0])
	if err != nil {
		return nil, dbus.MakeFailedError(err)
	}
	var resultSet []string
	for _, data := range result.Data {
		resultSet = append(resultSet, fmt.Sprintf("%v, %v - %v",
			data.Japanese[0].Word, data.Japanese[0].Reading, data.Senses[0].EnglishDefinitions))
	}
	slog.Info("received results", "length", len(resultSet))
	return resultSet, nil
}

func (p *Provider) GetResultMetas(identifiers []string) ([]map[string]dbus.Variant, *dbus.Error) {
	slog.Info("GetResultMetas", "identifiers", identifiers)
	metas := make([]map[string]dbus.Variant, len(identifiers))
	for i, id := range identifiers {
		metas[i] = map[string]dbus.Variant{
			"id":          dbus.MakeVariant(id),
			"name":        dbus.MakeVariant(id),
			"description": dbus.MakeVariant("Search result"),
		}
	}
	return metas, nil
}

func (p *Provider) ActivateResult(identifier string, terms []string, timestamp uint32) *dbus.Error {
	slog.Info("ActivateResult", "identifier", identifier, "terms", terms)
	return nil
}

func (p *Provider) LaunchSearch(terms []string, timestamp uint32) *dbus.Error {
	slog.Info("LaunchSearch", "terms", terms)
	return nil
}
