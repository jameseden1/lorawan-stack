// Code generated by protoc-gen-go-flags. DO NOT EDIT.
// versions:
// - protoc-gen-go-flags v0.0.0-dev
// - protoc              v3.17.3
// source: lorawan-stack/api/application.proto

package ttnpb

import (
	flagsplugin "github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	gogo "github.com/TheThingsIndustries/protoc-gen-go-flags/gogo"
	pflag "github.com/spf13/pflag"
)

// AddSelectFlagsForApplication adds flags to select fields in Application.
func AddSelectFlagsForApplication(flags *pflag.FlagSet, prefix string) {
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("ids", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("ids", prefix), true)))
	AddSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("ids", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("created-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("created-at", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("updated-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("updated-at", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted-at", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("deleted-at", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("name", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("name", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("description", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("description", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("attributes", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("attributes", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("contact-info", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("contact-info", prefix), false)))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("administrative-contact", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("administrative-contact", prefix), true)))
	AddSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("administrative-contact", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("technical-contact", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("technical-contact", prefix), true)))
	AddSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("technical-contact", prefix))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("dev-eui-counter", prefix), flagsplugin.SelectDesc(flagsplugin.Prefix("dev-eui-counter", prefix), false)))
}

// SelectFromFlags outputs the fieldmask paths forApplication message from select flags.
func PathsFromSelectFlagsForApplication(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	ids, idsSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("ids", prefix))
	if err != nil {
		return paths, err
	}
	if idsSelect && ids {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("ids", prefix)))
	}
	if selectPaths, err := PathsFromSelectFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
		return paths, err
	} else {
		paths = append(paths, selectPaths...)
	}
	createdAt, createdAtSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("created-at", prefix))
	if err != nil {
		return paths, err
	}
	if createdAtSelect && createdAt {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("created-at", prefix)))
	}
	updatedAt, updatedAtSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("updated-at", prefix))
	if err != nil {
		return paths, err
	}
	if updatedAtSelect && updatedAt {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("updated-at", prefix)))
	}
	deletedAt, deletedAtSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted-at", prefix))
	if err != nil {
		return paths, err
	}
	if deletedAtSelect && deletedAt {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("deleted-at", prefix)))
	}
	name, nameSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("name", prefix))
	if err != nil {
		return paths, err
	}
	if nameSelect && name {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("name", prefix)))
	}
	description, descriptionSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("description", prefix))
	if err != nil {
		return paths, err
	}
	if descriptionSelect && description {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("description", prefix)))
	}
	attributes, attributesSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("attributes", prefix))
	if err != nil {
		return paths, err
	}
	if attributesSelect && attributes {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("attributes", prefix)))
	}
	contactInfo, contactInfoSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("contact-info", prefix))
	if err != nil {
		return paths, err
	}
	if contactInfoSelect && contactInfo {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("contact-info", prefix)))
	}
	administrativeContact, administrativeContactSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("administrative-contact", prefix))
	if err != nil {
		return paths, err
	}
	if administrativeContactSelect && administrativeContact {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("administrative-contact", prefix)))
	}
	if selectPaths, err := PathsFromSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("administrative-contact", prefix)); err != nil {
		return paths, err
	} else {
		paths = append(paths, selectPaths...)
	}
	technicalContact, technicalContactSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("technical-contact", prefix))
	if err != nil {
		return paths, err
	}
	if technicalContactSelect && technicalContact {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("technical-contact", prefix)))
	}
	if selectPaths, err := PathsFromSelectFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("technical-contact", prefix)); err != nil {
		return paths, err
	} else {
		paths = append(paths, selectPaths...)
	}
	devEuiCounter, devEuiCounterSelect, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("dev-eui-counter", prefix))
	if err != nil {
		return paths, err
	}
	if devEuiCounterSelect && devEuiCounter {
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("dev-eui-counter", prefix)))
	}
	return paths, nil
}

// AddSetFlagsForApplication adds flags to select fields in Application.
func AddSetFlagsForApplication(flags *pflag.FlagSet, prefix string) {
	AddSetFlagsForApplicationIdentifiers(flags, flagsplugin.Prefix("ids", prefix))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("created-at", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("updated-at", prefix), ""))
	flags.AddFlag(flagsplugin.NewTimestampFlag(flagsplugin.Prefix("deleted-at", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("name", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("description", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringStringMapFlag(flagsplugin.Prefix("attributes", prefix), ""))
	// FIXME: Skipping ContactInfo because repeated messages are currently not supported.
	AddSetFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("administrative-contact", prefix))
	AddSetFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("technical-contact", prefix))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("dev-eui-counter", prefix), ""))
}

// SetFromFlags sets the Application message from flags.
func (m *Application) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	idsSet := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("ids", prefix))
	if idsSet {
		m.Ids = &ApplicationIdentifiers{}
		if setPaths, err := m.Ids.SetFromFlags(flags, flagsplugin.Prefix("ids", prefix)); err != nil {
			return paths, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	createdAt, createdAtSet, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("created-at", prefix))
	if err != nil {
		return paths, err
	}
	if createdAtSet {
		m.CreatedAt = gogo.SetTimestamp(createdAt)
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("created-at", prefix)))
	}
	updatedAt, updatedAtSet, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("updated-at", prefix))
	if err != nil {
		return paths, err
	}
	if updatedAtSet {
		m.UpdatedAt = gogo.SetTimestamp(updatedAt)
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("updated-at", prefix)))
	}
	deletedAt, deletedAtSet, err := flagsplugin.GetTimestamp(flags, flagsplugin.Prefix("deleted-at", prefix))
	if err != nil {
		return paths, err
	}
	if deletedAtSet {
		m.DeletedAt = gogo.SetTimestamp(deletedAt)
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("deleted-at", prefix)))
	}
	name, nameSet, err := flagsplugin.GetString(flags, flagsplugin.Prefix("name", prefix))
	if err != nil {
		return paths, err
	}
	if nameSet {
		m.Name = name
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("name", prefix)))
	}
	description, descriptionSet, err := flagsplugin.GetString(flags, flagsplugin.Prefix("description", prefix))
	if err != nil {
		return paths, err
	}
	if descriptionSet {
		m.Description = description
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("description", prefix)))
	}
	attributes, attributesSet, err := flagsplugin.GetStringStringMap(flags, flagsplugin.Prefix("attributes", prefix))
	if err != nil {
		return paths, err
	}
	if attributesSet {
		m.Attributes = attributes
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("attributes", prefix)))
	}
	// FIXME: Skipping ContactInfo because it does not seem to implement AddSetFlags.
	administrativeContactSet := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("administrative-contact", prefix))
	if administrativeContactSet {
		m.AdministrativeContact = &OrganizationOrUserIdentifiers{}
		if setPaths, err := m.AdministrativeContact.SetFromFlags(flags, flagsplugin.Prefix("administrative-contact", prefix)); err != nil {
			return paths, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	technicalContactSet := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("technical-contact", prefix))
	if technicalContactSet {
		m.TechnicalContact = &OrganizationOrUserIdentifiers{}
		if setPaths, err := m.TechnicalContact.SetFromFlags(flags, flagsplugin.Prefix("technical-contact", prefix)); err != nil {
			return paths, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	devEuiCounter, devEuiCounterSet, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("dev-eui-counter", prefix))
	if err != nil {
		return paths, err
	}
	if devEuiCounterSet {
		m.DevEuiCounter = devEuiCounter
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("dev-eui-counter", prefix)))
	}
	return paths, nil
}

// AddSetFlagsForListApplicationsRequest adds flags to select fields in ListApplicationsRequest.
func AddSetFlagsForListApplicationsRequest(flags *pflag.FlagSet, prefix string) {
	AddSetFlagsForOrganizationOrUserIdentifiers(flags, flagsplugin.Prefix("collaborator", prefix))
	flags.AddFlag(flagsplugin.NewStringSliceFlag(flagsplugin.Prefix("field-mask", prefix), ""))
	flags.AddFlag(flagsplugin.NewStringFlag(flagsplugin.Prefix("order", prefix), ""))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("limit", prefix), ""))
	flags.AddFlag(flagsplugin.NewUint32Flag(flagsplugin.Prefix("page", prefix), ""))
	flags.AddFlag(flagsplugin.NewBoolFlag(flagsplugin.Prefix("deleted", prefix), ""))
}

// SetFromFlags sets the ListApplicationsRequest message from flags.
func (m *ListApplicationsRequest) SetFromFlags(flags *pflag.FlagSet, prefix string) (paths []string, err error) {
	collaboratorSet := flagsplugin.IsAnyPrefixSet(flags, flagsplugin.Prefix("collaborator", prefix))
	if collaboratorSet {
		m.Collaborator = &OrganizationOrUserIdentifiers{}
		if setPaths, err := m.Collaborator.SetFromFlags(flags, flagsplugin.Prefix("collaborator", prefix)); err != nil {
			return paths, err
		} else {
			paths = append(paths, setPaths...)
		}
	}
	fieldMask, fieldMaskSet, err := flagsplugin.GetStringSlice(flags, flagsplugin.Prefix("field-mask", prefix))
	if err != nil {
		return paths, err
	}
	if fieldMaskSet {
		m.FieldMask = gogo.SetFieldMask(fieldMask)
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("field-mask", prefix)))
	}
	order, orderSet, err := flagsplugin.GetString(flags, flagsplugin.Prefix("order", prefix))
	if err != nil {
		return paths, err
	}
	if orderSet {
		m.Order = order
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("order", prefix)))
	}
	limit, limitSet, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("limit", prefix))
	if err != nil {
		return paths, err
	}
	if limitSet {
		m.Limit = limit
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("limit", prefix)))
	}
	page, pageSet, err := flagsplugin.GetUint32(flags, flagsplugin.Prefix("page", prefix))
	if err != nil {
		return paths, err
	}
	if pageSet {
		m.Page = page
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("page", prefix)))
	}
	deleted, deletedSet, err := flagsplugin.GetBool(flags, flagsplugin.Prefix("deleted", prefix))
	if err != nil {
		return paths, err
	}
	if deletedSet {
		m.Deleted = deleted
		paths = append(paths, flagsplugin.FieldMaskFlag(flagsplugin.Prefix("deleted", prefix)))
	}
	return paths, nil
}
