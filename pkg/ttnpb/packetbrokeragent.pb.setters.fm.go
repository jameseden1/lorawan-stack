// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *PacketBrokerNetworkIdentifier) SetFields(src *PacketBrokerNetworkIdentifier, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "net_id":
			if len(subs) > 0 {
				return fmt.Errorf("'net_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.NetID = src.NetID
			} else {
				var zero uint32
				dst.NetID = zero
			}
		case "tenant_id":
			if len(subs) > 0 {
				return fmt.Errorf("'tenant_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.TenantID = src.TenantID
			} else {
				var zero string
				dst.TenantID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerDevAddrBlock) SetFields(src *PacketBrokerDevAddrBlock, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "dev_addr_prefix":
			if len(subs) > 0 {
				var newDst, newSrc *DevAddrPrefix
				if (src == nil || src.DevAddrPrefix == nil) && dst.DevAddrPrefix == nil {
					continue
				}
				if src != nil {
					newSrc = src.DevAddrPrefix
				}
				if dst.DevAddrPrefix != nil {
					newDst = dst.DevAddrPrefix
				} else {
					newDst = &DevAddrPrefix{}
					dst.DevAddrPrefix = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DevAddrPrefix = src.DevAddrPrefix
				} else {
					dst.DevAddrPrefix = nil
				}
			}
		case "home_network_cluster_id":
			if len(subs) > 0 {
				return fmt.Errorf("'home_network_cluster_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HomeNetworkClusterID = src.HomeNetworkClusterID
			} else {
				var zero string
				dst.HomeNetworkClusterID = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerNetwork) SetFields(src *PacketBrokerNetwork, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerNetworkIdentifier
				if (src == nil || src.Id == nil) && dst.Id == nil {
					continue
				}
				if src != nil {
					newSrc = src.Id
				}
				if dst.Id != nil {
					newDst = dst.Id
				} else {
					newDst = &PacketBrokerNetworkIdentifier{}
					dst.Id = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Id = src.Id
				} else {
					dst.Id = nil
				}
			}
		case "name":
			if len(subs) > 0 {
				return fmt.Errorf("'name' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Name = src.Name
			} else {
				var zero string
				dst.Name = zero
			}
		case "dev_addr_blocks":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_addr_blocks' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevAddrBlocks = src.DevAddrBlocks
			} else {
				dst.DevAddrBlocks = nil
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "listed":
			if len(subs) > 0 {
				return fmt.Errorf("'listed' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Listed = src.Listed
			} else {
				var zero bool
				dst.Listed = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerNetworks) SetFields(src *PacketBrokerNetworks, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "networks":
			if len(subs) > 0 {
				return fmt.Errorf("'networks' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Networks = src.Networks
			} else {
				dst.Networks = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerInfo) SetFields(src *PacketBrokerInfo, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "registration":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerNetwork
				if (src == nil || src.Registration == nil) && dst.Registration == nil {
					continue
				}
				if src != nil {
					newSrc = src.Registration
				}
				if dst.Registration != nil {
					newDst = dst.Registration
				} else {
					newDst = &PacketBrokerNetwork{}
					dst.Registration = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Registration = src.Registration
				} else {
					dst.Registration = nil
				}
			}
		case "forwarder_enabled":
			if len(subs) > 0 {
				return fmt.Errorf("'forwarder_enabled' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ForwarderEnabled = src.ForwarderEnabled
			} else {
				var zero bool
				dst.ForwarderEnabled = zero
			}
		case "home_network_enabled":
			if len(subs) > 0 {
				return fmt.Errorf("'home_network_enabled' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.HomeNetworkEnabled = src.HomeNetworkEnabled
			} else {
				var zero bool
				dst.HomeNetworkEnabled = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerRoutingPolicyUplink) SetFields(src *PacketBrokerRoutingPolicyUplink, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_request":
			if len(subs) > 0 {
				return fmt.Errorf("'join_request' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinRequest = src.JoinRequest
			} else {
				var zero bool
				dst.JoinRequest = zero
			}
		case "mac_data":
			if len(subs) > 0 {
				return fmt.Errorf("'mac_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MacData = src.MacData
			} else {
				var zero bool
				dst.MacData = zero
			}
		case "application_data":
			if len(subs) > 0 {
				return fmt.Errorf("'application_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationData = src.ApplicationData
			} else {
				var zero bool
				dst.ApplicationData = zero
			}
		case "signal_quality":
			if len(subs) > 0 {
				return fmt.Errorf("'signal_quality' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SignalQuality = src.SignalQuality
			} else {
				var zero bool
				dst.SignalQuality = zero
			}
		case "localization":
			if len(subs) > 0 {
				return fmt.Errorf("'localization' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Localization = src.Localization
			} else {
				var zero bool
				dst.Localization = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerRoutingPolicyDownlink) SetFields(src *PacketBrokerRoutingPolicyDownlink, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "join_accept":
			if len(subs) > 0 {
				return fmt.Errorf("'join_accept' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinAccept = src.JoinAccept
			} else {
				var zero bool
				dst.JoinAccept = zero
			}
		case "mac_data":
			if len(subs) > 0 {
				return fmt.Errorf("'mac_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MacData = src.MacData
			} else {
				var zero bool
				dst.MacData = zero
			}
		case "application_data":
			if len(subs) > 0 {
				return fmt.Errorf("'application_data' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ApplicationData = src.ApplicationData
			} else {
				var zero bool
				dst.ApplicationData = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerDefaultRoutingPolicy) SetFields(src *PacketBrokerDefaultRoutingPolicy, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				dst.UpdatedAt = nil
			}
		case "uplink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyUplink
				if (src == nil || src.Uplink == nil) && dst.Uplink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Uplink
				}
				if dst.Uplink != nil {
					newDst = dst.Uplink
				} else {
					newDst = &PacketBrokerRoutingPolicyUplink{}
					dst.Uplink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Uplink = src.Uplink
				} else {
					dst.Uplink = nil
				}
			}
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &PacketBrokerRoutingPolicyDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerRoutingPolicy) SetFields(src *PacketBrokerRoutingPolicy, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "forwarder_id":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerNetworkIdentifier
				if (src == nil || src.ForwarderId == nil) && dst.ForwarderId == nil {
					continue
				}
				if src != nil {
					newSrc = src.ForwarderId
				}
				if dst.ForwarderId != nil {
					newDst = dst.ForwarderId
				} else {
					newDst = &PacketBrokerNetworkIdentifier{}
					dst.ForwarderId = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.ForwarderId = src.ForwarderId
				} else {
					dst.ForwarderId = nil
				}
			}
		case "home_network_id":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerNetworkIdentifier
				if (src == nil || src.HomeNetworkId == nil) && dst.HomeNetworkId == nil {
					continue
				}
				if src != nil {
					newSrc = src.HomeNetworkId
				}
				if dst.HomeNetworkId != nil {
					newDst = dst.HomeNetworkId
				} else {
					newDst = &PacketBrokerNetworkIdentifier{}
					dst.HomeNetworkId = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.HomeNetworkId = src.HomeNetworkId
				} else {
					dst.HomeNetworkId = nil
				}
			}
		case "updated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'updated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UpdatedAt = src.UpdatedAt
			} else {
				dst.UpdatedAt = nil
			}
		case "uplink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyUplink
				if (src == nil || src.Uplink == nil) && dst.Uplink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Uplink
				}
				if dst.Uplink != nil {
					newDst = dst.Uplink
				} else {
					newDst = &PacketBrokerRoutingPolicyUplink{}
					dst.Uplink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Uplink = src.Uplink
				} else {
					dst.Uplink = nil
				}
			}
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &PacketBrokerRoutingPolicyDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetPacketBrokerDefaultRoutingPolicyRequest) SetFields(src *SetPacketBrokerDefaultRoutingPolicyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyUplink
				if (src == nil || src.Uplink == nil) && dst.Uplink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Uplink
				}
				if dst.Uplink != nil {
					newDst = dst.Uplink
				} else {
					newDst = &PacketBrokerRoutingPolicyUplink{}
					dst.Uplink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Uplink = src.Uplink
				} else {
					dst.Uplink = nil
				}
			}
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &PacketBrokerRoutingPolicyDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListHomeNetworkRoutingPoliciesRequest) SetFields(src *ListHomeNetworkRoutingPoliciesRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *PacketBrokerRoutingPolicies) SetFields(src *PacketBrokerRoutingPolicies, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "policies":
			if len(subs) > 0 {
				return fmt.Errorf("'policies' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Policies = src.Policies
			} else {
				dst.Policies = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *SetPacketBrokerRoutingPolicyRequest) SetFields(src *SetPacketBrokerRoutingPolicyRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "home_network_id":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerNetworkIdentifier
				if (src == nil || src.HomeNetworkId == nil) && dst.HomeNetworkId == nil {
					continue
				}
				if src != nil {
					newSrc = src.HomeNetworkId
				}
				if dst.HomeNetworkId != nil {
					newDst = dst.HomeNetworkId
				} else {
					newDst = &PacketBrokerNetworkIdentifier{}
					dst.HomeNetworkId = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.HomeNetworkId = src.HomeNetworkId
				} else {
					dst.HomeNetworkId = nil
				}
			}
		case "uplink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyUplink
				if (src == nil || src.Uplink == nil) && dst.Uplink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Uplink
				}
				if dst.Uplink != nil {
					newDst = dst.Uplink
				} else {
					newDst = &PacketBrokerRoutingPolicyUplink{}
					dst.Uplink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Uplink = src.Uplink
				} else {
					dst.Uplink = nil
				}
			}
		case "downlink":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerRoutingPolicyDownlink
				if (src == nil || src.Downlink == nil) && dst.Downlink == nil {
					continue
				}
				if src != nil {
					newSrc = src.Downlink
				}
				if dst.Downlink != nil {
					newDst = dst.Downlink
				} else {
					newDst = &PacketBrokerRoutingPolicyDownlink{}
					dst.Downlink = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Downlink = src.Downlink
				} else {
					dst.Downlink = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListHomeNetworksRequest) SetFields(src *ListHomeNetworksRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ListForwarderRoutingPoliciesRequest) SetFields(src *ListForwarderRoutingPoliciesRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "home_network_id":
			if len(subs) > 0 {
				var newDst, newSrc *PacketBrokerNetworkIdentifier
				if (src == nil || src.HomeNetworkId == nil) && dst.HomeNetworkId == nil {
					continue
				}
				if src != nil {
					newSrc = src.HomeNetworkId
				}
				if dst.HomeNetworkId != nil {
					newDst = dst.HomeNetworkId
				} else {
					newDst = &PacketBrokerNetworkIdentifier{}
					dst.HomeNetworkId = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.HomeNetworkId = src.HomeNetworkId
				} else {
					dst.HomeNetworkId = nil
				}
			}
		case "limit":
			if len(subs) > 0 {
				return fmt.Errorf("'limit' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Limit = src.Limit
			} else {
				var zero uint32
				dst.Limit = zero
			}
		case "page":
			if len(subs) > 0 {
				return fmt.Errorf("'page' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Page = src.Page
			} else {
				var zero uint32
				dst.Page = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}