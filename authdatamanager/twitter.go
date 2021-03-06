package authdatamanager

import (
	"github.com/lfq7413/tomato/errs"
	"github.com/lfq7413/tomato/types"
	"github.com/lfq7413/tomato/utils"
)

type twitter struct{}

func (a twitter) ValidateAuthData(authData types.M, options types.M) error {
	// 具体接口参考：https://dev.twitter.com/rest/reference/get/account/verify_credentials
	// https://dev.twitter.com/rest/tools/console
	client := NewOAuth(options)
	client.Host = "https://api.twitter.com"
	client.AuthToken = utils.S(authData["access_token"])
	client.AuthTokenSecret = utils.S(authData["auth_token_secret"])
	data, err := client.Get("/1.1/account/verify_credentials.json", nil)
	if err != nil {
		return errs.E(errs.ObjectNotFound, "Failed to validate this access token with Twitter.")
	}
	if data["id"] != nil && utils.S(data["id"]) == utils.S(authData["id"]) {
		return nil
	}
	return errs.E(errs.ObjectNotFound, "Twitter auth is invalid for this user.")
}
