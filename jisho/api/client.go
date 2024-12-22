package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const defaultURL = "https://jisho.org/api/v1"

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}

func NewDefaultClient() *Client {
	return NewClient(defaultURL)
}

func (c *Client) Search(keyword string) (Result, error) {
	// https://jisho.org/api/v1/search/words?keyword=house
	url := c.BaseURL + "/search/words?keyword=" + keyword
	resp, err := http.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, err
	}
	if result.Meta.Status != 200 {
		return Result{}, fmt.Errorf("search failed with status: %d", result.Meta.Status)
	}
	return result, nil
}

type Data struct {
	Slug        string      `json:"slug"`
	IsCommon    bool        `json:"is_common"`
	Tags        []string    `json:"tags"`
	JLPT        []string    `json:"jlpt"`
	Japanese    []Japanese  `json:"japanese"`
	Senses      []Sense     `json:"senses"`
	Attribution Attribution `json:"attribution"`
}

type Japanese struct {
	Word    string `json:"word"`
	Reading string `json:"reading"`
}

type Sense struct {
	EnglishDefinitions []string `json:"english_definitions"`
	PartsOfSpeech      []string `json:"parts_of_speech"`
	Links              []Link   `json:"links"`
	Tags               []string `json:"tags"`
	Restrictions       []string `json:"restrictions"`
	SeeAlso            []string `json:"see_also"`
	Antonyms           []string `json:"antonyms"`
	Source             []string `json:"source"`
	Info               []string `json:"info"`
}

type Link struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type Attribution struct {
	JMDict   bool `json:"jmdict"`
	JMNEDict bool `json:"jmnedict"`
	DBpedia  bool `json:"dbpedia"`
}

type Result struct {
	Meta Meta   `json:"meta"`
	Data []Data `json:"data"`
}

type Meta struct {
	Status int `json:"status"`
}
