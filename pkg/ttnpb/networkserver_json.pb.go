// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.1
// - protoc             v3.9.1
// source: lorawan-stack/api/networkserver.proto

package ttnpb

import (
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the GetDefaultMACSettingsRequest message to JSON.
func (x *GetDefaultMACSettingsRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.FrequencyPlanId != "" || s.HasField("frequency_plan_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("frequency_plan_id")
		s.WriteString(x.FrequencyPlanId)
	}
	if x.LorawanPhyVersion != 0 || s.HasField("lorawan_phy_version") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("lorawan_phy_version")
		x.LorawanPhyVersion.MarshalProtoJSON(s)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GetDefaultMACSettingsRequest to JSON.
func (x GetDefaultMACSettingsRequest) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the GetDefaultMACSettingsRequest message from JSON.
func (x *GetDefaultMACSettingsRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "frequency_plan_id", "frequencyPlanId":
			s.AddField("frequency_plan_id")
			x.FrequencyPlanId = s.ReadString()
		case "lorawan_phy_version", "lorawanPhyVersion":
			s.AddField("lorawan_phy_version")
			x.LorawanPhyVersion.UnmarshalProtoJSON(s)
		}
	})
}

// UnmarshalJSON unmarshals the GetDefaultMACSettingsRequest from JSON.
func (x *GetDefaultMACSettingsRequest) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
