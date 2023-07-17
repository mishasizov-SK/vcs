/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package oidc4ci

import (
	"context"
	"fmt"

	"github.com/trustbloc/vcs/pkg/event/spi"
	"github.com/trustbloc/vcs/pkg/restapi/v1/common"
)

// StoreAuthorizationCode stores authorization code from issuer provider.
func (s *Service) StoreAuthorizationCode(
	ctx context.Context,
	opState string,
	code string,
	flowData *common.WalletInitiatedFlowData,
) (TxID, error) {
	var tx *Transaction
	var err error
	if flowData != nil { // its wallet initiated issuance, first we need to initiate issuance
		tx, err = s.initiateIssuanceWithWalletFlow(ctx, flowData)
	} else {
		tx, err = s.store.FindByOpState(ctx, opState)
	}

	if err != nil {
		return "", err
	}

	if err != nil {
		return "", fmt.Errorf("get transaction by opstate: %w", err)
	}

	tx.IssuerAuthCode = code
	if err = s.store.Update(ctx, tx); err != nil {
		s.sendFailedTransactionEvent(ctx, tx, err)
		return "", err
	}

	if err = s.sendTransactionEvent(ctx, tx, spi.IssuerOIDCInteractionAuthorizationCodeStored); err != nil {
		s.sendFailedTransactionEvent(ctx, tx, err)
		return "", err
	}

	return tx.ID, nil
}

func (s *Service) initiateIssuanceWithWalletFlow(
	ctx context.Context,
	flowData *common.WalletInitiatedFlowData,
) (*Transaction, error) {
	tx, err := s.InitiateIssuance(ctx, &InitiateIssuanceRequest{
		CredentialTemplateID:      flowData.CredentialTemplateId,
		ClientInitiateIssuanceURL: "",
		ClientWellKnownURL:        "",
		ClaimEndpoint:             flowData.ClaimEndpoint,
		GrantType:                 "authorization_code",
		ResponseType:              "code",
		Scope:                     flowData.Scopes,
		OpState:                   flowData.OpState,
		ClaimData:                 nil,
		UserPinRequired:           false,
		CredentialExpiresAt:       nil,
		CredentialName:            "",
		CredentialDescription:     "",
		WalletInitiatedIssuance:   true,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("can not initiate issuance for wallet flow. %w", err)
	}

	return tx.Tx, nil
}
