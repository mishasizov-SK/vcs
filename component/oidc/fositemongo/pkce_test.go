/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fositemongo

import (
	"context"
	"testing"
	"time"

	"github.com/ory/fosite"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	"github.com/trustbloc/vcs/pkg/storage/mongodb"
)

func TestPKCE(t *testing.T) {
	pool, mongoDBResource := startMongoDBContainer(t)

	defer func() {
		assert.NoError(t, pool.Purge(mongoDBResource), "failed to purge MongoDB resource")
	}()

	client, mongoErr := mongodb.New(mongoDBConnString, "testdb", mongodb.WithTimeout(time.Second*10))
	assert.NoError(t, mongoErr)

	testCases := []struct {
		name      string
		useRevoke bool
	}{
		{
			name:      "user delete",
			useRevoke: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s, err := NewStore(context.Background(), client)
			assert.NoError(t, err)

			dbClient := &Client{
				ID:             uuid.New(),
				Secret:         nil,
				RotatedSecrets: nil,
				RedirectURIs:   nil,
				GrantTypes:     nil,
				ResponseTypes:  nil,
				Scopes:         []string{"awesome"},
				Audience:       nil,
				Public:         false,
			}

			_, err = s.InsertClient(context.Background(), *dbClient)
			assert.NoError(t, err)

			sign := uuid.New()
			ses := &fosite.Request{
				ID:                uuid.New(),
				Client:            dbClient,
				RequestedScope:    []string{"scope1"},
				GrantedScope:      []string{"scope1"},
				RequestedAudience: []string{"aud1"},
				GrantedAudience:   []string{"aud2"},
				Lang:              language.Tag{},
				Session:           &fosite.DefaultSession{},
			}

			err = s.CreatePKCERequestSession(context.TODO(), sign, ses)
			assert.NoError(t, err)

			dbSes, err := s.GetPKCERequestSession(context.TODO(), sign, ses.Session)
			assert.NoError(t, err)
			assert.Equal(t, ses, dbSes)
			assert.Equal(t, dbClient.ID, dbSes.GetClient().GetID())

			err = s.DeletePKCERequestSession(context.TODO(), sign)
			assert.NoError(t, err)

			resp, err := s.GetPKCERequestSession(context.TODO(), sign, ses.Session)
			assert.Nil(t, resp)
			assert.ErrorIs(t, err, ErrDataNotFound)
		})
	}
}
