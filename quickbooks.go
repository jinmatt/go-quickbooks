package quickbooks

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/jinmatt/go-quickbooks/sdk"
	"github.com/jinmatt/go-quickbooks/sdk/consts"
	"github.com/jinmatt/go-quickbooks/sdk/errors"
	"github.com/jinmatt/go-quickbooks/sdk/types"
	"github.com/mrjones/oauth"
)

type quickbooks struct {
	oauthClient *oauth.Consumer
	accessToken *oauth.AccessToken
	baseURL     string
	realmID     string
}

// NewClient creates a new client to work with Quickbooks
func NewClient(consumerKey, consumerSecret, oauthToken, oauthSecret, realmID string, isSandbox bool) *quickbooks {
	q := &quickbooks{}
	q.oauthClient = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			AuthorizeTokenUrl: sdk.AuthorizeURL,
			RequestTokenUrl:   sdk.RequestTokenURL,
			AccessTokenUrl:    sdk.AccessTokenURL,
		},
	)

	q.accessToken = &oauth.AccessToken{
		Token:  oauthToken,
		Secret: oauthSecret,
	}

	q.realmID = realmID

	if isSandbox {
		q.baseURL = sdk.SandboxURL
	} else {
		q.baseURL = sdk.ProductionURL
	}

	return q
}

func (qb *quickbooks) makeGetRequest(endpoint string) (*http.Response, error) {
	rURL := qb.baseURL + endpoint
	req, err := http.NewRequest(consts.Verb.Get, rURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "text/plain")

	httpClient, err := qb.oauthClient.MakeHttpClient(qb.accessToken)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return qb.handleResponseError(res)
	}

	return res, nil
}

func (qb *quickbooks) makePostRequest(endpoint string, body interface{}) (*http.Response, error) {
	rURL := qb.baseURL + endpoint
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(consts.Verb.Post, rURL, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	httpClient, err := qb.oauthClient.MakeHttpClient(qb.accessToken)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return qb.handleResponseError(res)
	}

	return res, nil
}

func (qb *quickbooks) handleResponseError(res *http.Response) (*http.Response, error) {
	if res.Header.Get("Content-Type") == consts.ContextTypeXML {
		xmlErrorRes := types.IntuitResponse{}
		if err := xml.NewDecoder(res.Body).Decode(&xmlErrorRes); err != nil {
			return nil, err
		}

		if xmlErrorRes.Fault.Type == consts.QBFaultType {
			return nil, errors.QBAuthFailure
		}

		return nil, errors.NewSDKError(xmlErrorRes.Fault.Error.Code, xmlErrorRes.Fault.Error.Detail)
	}

	qbError := types.Error{}
	if err := json.NewDecoder(res.Body).Decode(&qbError); err != nil {
		return nil, err
	}

	return nil, errors.NewSDKError(qbError.Fault.Error[0].Code, qbError.Fault.Error[0].Detail)
}
