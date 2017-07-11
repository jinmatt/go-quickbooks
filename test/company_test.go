package test

import (
	"testing"

	"github.com/jinmatt/go-quickbooks"
	"github.com/tylerb/is"
)

func TestCompanyInfo(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(
		ConsumerKey,
		ConsumerSecret,
		AccessToken,
		AccessSecret,
		RealmID,
		true,
	)

	companyInfo, err := qbo.GetCompanyInfo()
	is.NotErr(err)
	is.NotNil(companyInfo.CompanyInfo.CompanyName)
}
