// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

type membershipCache struct {
	store.MembershipStore
	redis *redis.Client
	ttl   time.Duration
}

// GetMembershipCache wraps the MembershipStore with a cache.
// Make sure to not call FindIndirectMemberships or GetMember after calling
// SetMember in the same transaction, this may result in an inconsistent cache.
func GetMembershipCache(store store.MembershipStore, redis *redis.Client, ttl time.Duration) store.MembershipStore {
	return &membershipCache{
		MembershipStore: store,
		redis:           redis,
		ttl:             ttl,
	}
}

// TODO: Add FindIndirectMemberships (https://github.com/TheThingsNetwork/lorawan-stack/issues/443).

func (c *membershipCache) cacheKey(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityID *ttnpb.EntityIdentifiers) string {
	return c.redis.Key("membership", id.EntityType(), unique.ID(ctx, id), entityID.EntityType(), unique.ID(ctx, entityID))
}

func (c *membershipCache) GetMember(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityID *ttnpb.EntityIdentifiers) (*ttnpb.Rights, error) {
	cacheKey := c.cacheKey(ctx, id, entityID)
	if cached, err := c.redis.Get(ctx, cacheKey).Bytes(); err == nil {
		if len(cached) == 0 {
			return nil, errMembershipNotFound.WithAttributes(
				"account_id", id.IDString(),
				"entity_type", entityID.EntityType(),
				"entity_id", entityID.IDString(),
			)
		}
		var rights ttnpb.Rights
		if err = proto.Unmarshal(cached, &rights); err == nil {
			return &rights, nil
		}
	}
	rights, err := c.MembershipStore.GetMember(ctx, id, entityID)
	if err != nil {
		if errors.IsNotFound(err) {
			if cacheErr := c.redis.Set(ctx, cacheKey, "", c.ttl).Err(); cacheErr != nil {
				log.FromContext(ctx).WithError(cacheErr).Error("Failed to set membership cache")
			}
		}
		return nil, err
	}
	if cache, err := proto.Marshal(rights); err == nil {
		if cacheErr := c.redis.Set(ctx, cacheKey, cache, c.ttl).Err(); cacheErr != nil {
			log.FromContext(ctx).WithError(cacheErr).Error("Failed to set membership cache")
		}
	}
	return rights, err
}

func (c *membershipCache) SetMember(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityID *ttnpb.EntityIdentifiers, rights *ttnpb.Rights) error {
	err := c.MembershipStore.SetMember(ctx, id, entityID, rights)
	if err != nil {
		return err
	}
	// NOTE: Only invalidate. We can't set the new rights, since we don't know if
	// the transaction will succeed.
	if cacheErr := c.redis.Del(ctx, c.cacheKey(ctx, id, entityID)).Err(); cacheErr != nil {
		log.FromContext(ctx).WithError(cacheErr).Error("Failed to invalidate membership cache")
	}
	return nil
}
