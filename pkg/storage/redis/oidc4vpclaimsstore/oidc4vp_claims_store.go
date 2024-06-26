/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package oidc4vpclaimsstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	redisapi "github.com/redis/go-redis/v9"

	"github.com/trustbloc/vcs/pkg/service/oidc4vp"
	"github.com/trustbloc/vcs/pkg/storage/redis"
)

const (
	keyPrefix = "oidc4vp_claims"
)

// Store stores claim data with expiration.
type Store struct {
	redisClient *redis.Client
	defaultTTL  time.Duration
}

// New creates presentation claims store.
func New(redisClient *redis.Client, ttlSec int32) *Store {
	return &Store{
		redisClient: redisClient,
		defaultTTL:  time.Duration(ttlSec) * time.Second,
	}
}

func (s *Store) Create(claims *oidc4vp.ClaimData, profileReceivedClaimsDataTTL int32) (string, error) {
	ttl := s.defaultTTL
	if profileReceivedClaimsDataTTL > 0 {
		ttl = time.Duration(profileReceivedClaimsDataTTL) * time.Second
	}

	doc := &claimDataDocument{
		ExpireAt:  time.Now().Add(ttl),
		ClaimData: claims,
	}

	ctxWithTimeout, cancel := s.redisClient.ContextWithTimeout()
	defer cancel()

	key := resolveRedisKey(uuid.NewString())

	if err := s.redisClient.API().Set(ctxWithTimeout, key, doc, ttl).Err(); err != nil {
		return "", fmt.Errorf("redis insert received claims data: %w", err)
	}

	return key, nil
}

func (s *Store) Get(claimDataID string) (*oidc4vp.ClaimData, error) {
	ctxWithTimeout, cancel := s.redisClient.ContextWithTimeout()
	defer cancel()

	b, err := s.redisClient.API().Get(ctxWithTimeout, claimDataID).Bytes()
	if err != nil {
		if errors.Is(err, redisapi.Nil) {
			return nil, oidc4vp.ErrDataNotFound
		}

		return nil, fmt.Errorf("find: %w", err)
	}

	var doc claimDataDocument
	if err = json.Unmarshal(b, &doc); err != nil {
		return nil, fmt.Errorf("claim data decode: %w", err)
	}

	if doc.ExpireAt.Before(time.Now().UTC()) {
		return nil, oidc4vp.ErrDataNotFound
	}

	return doc.ClaimData, nil
}

func (s *Store) Delete(claimDataID string) error {
	ctxWithTimeout, cancel := s.redisClient.ContextWithTimeout()
	defer cancel()

	err := s.redisClient.API().Del(ctxWithTimeout, claimDataID).Err()
	if err != nil {
		return fmt.Errorf("failed to delete received claims with id[%s]: %w", claimDataID, err)
	}

	return nil
}

func resolveRedisKey(id string) string {
	return fmt.Sprintf("%s-%s", keyPrefix, id)
}
