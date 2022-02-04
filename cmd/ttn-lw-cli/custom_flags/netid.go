// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

package custom_flags

import (
	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

func New3BytesFlag(name, usage string) *pflag.Flag {
	return &pflag.Flag{
		Name:  name,
		Usage: usage,
		Value: &ExactBytesValue{length: 3},
	}
}

func GetNetID(fs *pflag.FlagSet, name string) (value types.NetID, set bool, err error) {
	flag := fs.Lookup(name)
	if flag == nil {
		return types.NetID{}, false, &flagsplugin.ErrFlagNotFound{FlagName: name}
	}
	var netID types.NetID
	if !flag.Changed {
		return netID, flag.Changed, nil
	}
	if err := netID.Unmarshal(flag.Value.(*ExactBytesValue).Value); err != nil {
		return netID, false, err
	}
	return netID, flag.Changed, nil
}
