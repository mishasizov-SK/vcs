/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"errors"

	"github.com/samber/lo"

	vcsverifiable "github.com/trustbloc/vcs/pkg/doc/verifiable"
	"github.com/trustbloc/vcs/pkg/restapi/resterr"
	"github.com/trustbloc/vcs/pkg/restapi/v1/common"
	"github.com/trustbloc/vcs/pkg/service/issuecredential"
)

func ValidateAuthorizationDetails(
	authorizationDetails []common.AuthorizationDetails) ([]*issuecredential.AuthorizationDetails, error) {
	var authorizationDetailsDomain []*issuecredential.AuthorizationDetails

	for _, ad := range authorizationDetails {
		if ad.Type != "openid_credential" {
			return nil, resterr.NewValidationError(resterr.InvalidValue, "authorization_details.type",
				errors.New("type should be 'openid_credential'"))
		}

		oidcCredentialFormat := lo.FromPtr(ad.Format)
		credentialConfigurationID := lo.FromPtr(ad.CredentialConfigurationId)

		mapped := &issuecredential.AuthorizationDetails{
			Type:                      ad.Type,
			Locations:                 lo.FromPtr(ad.Locations),
			CredentialConfigurationID: "",
			Format:                    "",
			CredentialDefinition:      nil,
		}

		switch {
		case credentialConfigurationID != "": // Priority 1. Based on credentialConfigurationID.
			mapped.CredentialConfigurationID = credentialConfigurationID
		case oidcCredentialFormat != "": // Priority 2. Based on credentialFormat.
			_, err := common.ValidateVCFormat(common.VCFormat(oidcCredentialFormat))
			if err != nil {
				return nil, resterr.NewValidationError(resterr.InvalidValue, "authorization_details.format", err)
			}

			mapped.Format = vcsverifiable.OIDCFormat(oidcCredentialFormat)

			if ad.CredentialDefinition == nil {
				return nil, resterr.NewValidationError(resterr.InvalidValue,
					"authorization_details.credential_definition", errors.New("not supplied"))
			}

			mapped.CredentialDefinition = &issuecredential.CredentialDefinition{
				Context:           lo.FromPtr(ad.CredentialDefinition.Context),
				CredentialSubject: lo.FromPtr(ad.CredentialDefinition.CredentialSubject),
				Type:              ad.CredentialDefinition.Type,
			}
		default:
			return nil, resterr.NewValidationError(resterr.InvalidValue,
				"authorization_details.credential_configuration_id",
				errors.New("neither credentialFormat nor credentialConfigurationID supplied"))
		}

		authorizationDetailsDomain = append(authorizationDetailsDomain, mapped)
	}

	return authorizationDetailsDomain, nil
}
