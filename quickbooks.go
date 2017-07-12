package quickbooks

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
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

// GetCompanyInfo returns company info based on realm ID/company ID passed to NewClient options
func (qb *quickbooks) GetCompanyInfo() (*types.Company, error) {
	endpoint := fmt.Sprintf("/company/%s/companyinfo/%s", qb.realmID, qb.realmID)

	res, err := qb.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	company := types.Company{}
	err = json.NewDecoder(res.Body).Decode(&company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (qb *quickbooks) makeGetRequest(endpoint string) (*http.Response, error) {
	rURL := qb.baseURL + endpoint
	req, err := http.NewRequest("GET", rURL, nil)
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
		if res.Header.Get("Content-Type") == consts.ContextTypeXML {
			xmlErrorRes := types.IntuitResponse{}
			err = xml.NewDecoder(res.Body).Decode(&xmlErrorRes)
			if err != nil {
				return nil, errors.QBApiFailure
			}

			if xmlErrorRes.Fault.Type == consts.QBFaultType {
				return nil, errors.QBAuthFailure
			}

			return nil, errors.QBApiFailure
		}

		qbError := types.Error{}
		err = json.NewDecoder(res.Body).Decode(&qbError)
		if err != nil {
			return nil, err
		}

		return nil, errors.NewSDKError(qbError.Fault.Error[0].Code, qbError.Fault.Error[0].Message)
	}

	return res, nil
}
