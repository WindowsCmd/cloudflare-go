package cloudflare

import (
	"net/http"
)

var baseURL = "https://api.cloudflare.com/client/v4"

type API struct {
	APIToken string
	APIBase string
	APIKey *APIKey
	UserAgent string
	headers map[string]string
	httpClient *http.Client
} 

//Create a new API object
//Token is the API token for your account
//email is the email address associated with your account leave blank if you're using an API Token 
func New(token string, email ...string) (*API, error) {

	api := &API{
		APIBase: baseURL,
		UserAgent: "windycloudflarego/1.0",		
	}

	if len(email) > 0 {
		api.APIKey = &APIKey{ 
			Key: token, 
			Email: email[0],
		}
	} else {
		api.APIToken = token
	}

	api.makeHeaders()

	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	} 

	return api, nil
}

//This function should be called after the http client is created
func (api *API) makeHeaders() {
	api.headers = make(map[string]string)
	
	if api.APIToken == "" {
		api.headers["X-Auth-Email"] = api.APIKey.Email
		api.headers["X-Auth-Key"] = api.APIKey.Key	
	} else {
		api.headers["Authorization"] = "Bearer " + api.APIToken
	}

	api.headers["User-Agent"] = api.UserAgent   
}