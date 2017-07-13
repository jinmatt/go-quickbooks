package test

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/jinmatt/go-quickbooks"
	"github.com/jinmatt/go-quickbooks/sdk/consts"
	seed "github.com/jinmatt/go-seed-rand"
	"github.com/tylerb/is"
)

func TestCreateAccount(t *testing.T) {
	is := is.New(t)

	qbo := quickbooks.NewClient(
		ConsumerKey,
		ConsumerSecret,
		AccessToken,
		AccessSecret,
		RealmID,
		true,
	)

	account := quickbooks.Account{}
	account.Name = randomdata.SillyName() + seed.RandomKey(7)
	account.AccountType = consts.QBAccountType

	newAccount, err := qbo.CreateAccount(account)
	is.NotErr(err)
	is.NotNil(newAccount.Account.ID)
	is.Equal(account.Name, newAccount.Account.Name)
}
