package instagram

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
)
var _ = fmt.Printf

func (i *Instagram) AuthorizeURL() string {	
	var params = url.Values{}
	params.Set("client_id", i.Config.ClientId)
	params.Set("redirect_uri", i.Config.RedirectUri)
	params.Set("response_type", i.Config.ResponseType)
	return i.Config.Domain + "/oauth/authorize?" + params.Encode()
}

func (i *Instagram) AuthorizeURLWithScope(scope []string) string {
	var params = url.Values{}
	params.Set("client_id", i.Config.ClientId)
	params.Set("redirect_uri", i.Config.RedirectUri)
	params.Set("response_type", i.Config.ResponseType)
	params.Set("scope", strings.Join(scope," "))
	return i.Config.Domain + "/oauth/authorize?" + params.Encode()
}

func (i *Instagram) RequestAccessToken(code string) (*Auth, error) {
	var params = url.Values{}
	params.Set("client_id", i.Config.ClientId)
	params.Set("client_secret", i.Config.ClientSecret)
	params.Set("redirect_uri", i.Config.RedirectUri)
	params.Set("grant_type", i.Config.GrantType)
	params.Set("code", code)
	
	var url = i.Config.Domain + "/oauth/access_token"
	var bodyType = "application/x-www-form-urlencoded"
	
	// http access
	resp, err := http.Post(url, bodyType, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) 
	if err != nil {
		return nil, err
	}

	var data = new(Auth)
	json.Unmarshal(body, &data)

	if data.Code > 299 {
		var code = data.Code
		var msg = data.ErrorMessage
		var err = &instagramError{ Message:msg, Status:code }
		return data, err
	}

	i.SetAccessToken(data.AccessToken)
	return data, nil
}
