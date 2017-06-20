package fonts

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const API = "https://emoji.pine.moe/api/v1/fonts"

type font struct {
	Key string `json:"key"`
	Name string `json:"name"`
}

type Fonts struct {
	client *http.Client
	api string
	List map[string]string
}

func New()*Fonts{
	return &Fonts{
		client: &http.Client{},
		api: API,
		List: map[string]string{},
	}
}

func (f *Fonts) Get() *Fonts {
	fs := []font{}
	list := map[string]string{}
	req, err := http.NewRequest("GET", f.api, nil)
	if err != nil {
		return f
	}
	resp, err := f.client.Do(req)
	if err != nil {
		return f
	}
	defer resp.Body.Close()

	bin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return f
	}

	if err := json.Unmarshal(bin, &fs); err != nil {
		return nil
	}

	for _, f := range fs {
		list[f.Key] = f.Name
	}
	f.List = list

	return f
}

