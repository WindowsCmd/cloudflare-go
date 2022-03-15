package cloudflare

import (
	"encoding/json"
	"fmt"
)

type Zone struct {
	ID                string   `json:"id"`
	Name              string   `json:"name"`
	DevelopmentMode   string   `json:"development_mode"`
	OriginalRegistrar string   `json:"original_registrar"`
	CreatedOn         string   `json:"created_on"`
	ModifiedOn        string   `json:"modified_on"`
	ActivatedOn       string   `json:"activated_on"`
	Status            string   `json:"status"`
	Paused            string   `json:"paused"`
	Type              string   `json:"type"`
	Permissions       []string `json:"permissions"`
}

func (api *API) GetZone(id string) (*Zone, error) {

	resp, err := api.httpClient.Get(fmt.Sprint(api.APIBase, "/zones"))

	if err != nil {
		fmt.Errorf("cloudflare GetZone returned an error %w", err)
	}

	z := &Zone{}

	err = json.Unmarshal([]byte(resp.Body.Close().Error()), z)

	if err != nil {
		return nil, err
	}

	return z, nil
}
