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
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var contextScoper func(context.Context, *gorm.DB) *gorm.DB

func withContext(ctx context.Context) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if contextScoper != nil {
			return contextScoper(ctx, db)
		}
		return db
	}
}

func withApplicationID(id ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(id) {
		case 0:
			return db
		case 1:
			if id[0] == "" {
				return db
			}
			return db.Where("application_id = ?", id[0])
		default:
			return db.Where("application_id IN (?)", id)
		}
	}
}

func withClientID(id ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(id) {
		case 0:
			return db
		case 1:
			return db.Where("client_id = ?", id[0])
		default:
			return db.Where("client_id IN (?)", id)
		}
	}
}

func withDeviceID(id ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(id) {
		case 0:
			return db
		case 1:
			if id[0] == "" {
				return db
			}
			return db.Where("device_id = ?", id[0])
		default:
			return db.Where("device_id IN (?)", id)
		}
	}
}

func withJoinEUI(eui ...EUI64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(eui) {
		case 0:
			return db
		case 1:
			if eui[0] == zeroEUI64 {
				return db
			}
			return db.Where("join_eui = ?", eui[0])
		default:
			return db.Where("join_eui IN (?)", eui)
		}
	}
}

func withDevEUI(eui ...EUI64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(eui) {
		case 0:
			return db
		case 1:
			if eui[0] == zeroEUI64 {
				return db
			}
			return db.Where("dev_eui = ?", eui[0])
		default:
			return db.Where("dev_eui IN (?)", eui)
		}
	}
}

func withApplicationAndDeviceID(applicationID, deviceID string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("application_id = ? AND device_id = ?", applicationID, deviceID)
	}
}

func withGatewayID(id ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(id) {
		case 0:
			return db
		case 1:
			if id[0] == "" {
				return db
			}
			return db.Where("gateway_id = ?", id[0])
		default:
			return db.Where("gateway_id IN (?)", id)
		}
	}
}

func withGatewayEUI(eui ...EUI64) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(eui) {
		case 0:
			return db
		case 1:
			if eui[0] == zeroEUI64 {
				return db
			}
			return db.Where("gateway_eui = ?", eui[0])
		default:
			return db.Where("gateway_eui IN (?)", eui)
		}
	}
}

func withOrganizationID(id ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Joins("LEFT JOIN accounts ON accounts.account_type = ? AND accounts.account_id = organizations.id", "organization")
		switch len(id) {
		case 0:
			return db
		case 1:
			return db.Where("accounts.uid = ?", id[0])
		default:
			return db.Where("accounts.uid IN (?)", id)
		}
	}
}

func withUserID(id ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Joins("LEFT JOIN accounts ON accounts.account_type = ? AND accounts.account_id = users.id", "user")
		switch len(id) {
		case 0:
			return db
		case 1:
			return db.Where("accounts.uid = ?", id[0])
		default:
			return db.Where("accounts.uid IN (?)", id)
		}
	}
}

func withPrimaryEmailAddress(emails ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch len(emails) {
		case 0:
			return db
		case 1:
			return db.Where("users.primary_email_address = ?", emails[0])
		default:
			return db.Where("users.primary_email_address IN (?)", emails)
		}
	}
}

func splitEndDeviceIDString(s string) (appID string, devID string) {
	sepIdx := strings.Index(s, ".")
	return s[:sepIdx], s[sepIdx+1:]
}

func withID(id ttnpb.IDStringer) func(*gorm.DB) *gorm.DB {
	switch entityTypeForID(id) {
	case "application":
		return withApplicationID(id.IDString())
	case "client":
		return withClientID(id.IDString())
	case "end_device":
		return withApplicationAndDeviceID(splitEndDeviceIDString(id.IDString()))
	case "gateway":
		return withGatewayID(id.IDString())
	case "organization":
		return withOrganizationID(id.IDString())
	case "user":
		return withUserID(id.IDString())
	default:
		panic(fmt.Sprintf("can't find scope for id type %T", id))
	}
}
