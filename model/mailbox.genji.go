// Code generated by genji.
// DO NOT EDIT!

package model

import (
	"errors"

	"github.com/asdine/genji/field"
	"github.com/asdine/genji/index"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
)

// GetField implements the field method of the record.Record interface.
func (m *Mailbox) GetField(name string) (field.Field, error) {
	switch name {
	case "ID":
		return field.NewString("ID", m.ID), nil
	case "AccountID":
		return field.NewString("AccountID", m.AccountID), nil
	case "Name":
		return field.NewString("Name", m.Name), nil
	case "Messages":
		return field.NewUint32("Messages", m.Messages), nil
	case "Recent":
		return field.NewUint32("Recent", m.Recent), nil
	case "Unseen":
		return field.NewUint32("Unseen", m.Unseen), nil
	case "UidNext":
		return field.NewUint32("UidNext", m.UidNext), nil
	case "UidValidity":
		return field.NewUint32("UidValidity", m.UidValidity), nil
	case "Subscribed":
		return field.NewBool("Subscribed", m.Subscribed), nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (m *Mailbox) Iterate(fn func(field.Field) error) error {
	var err error

	err = fn(field.NewString("ID", m.ID))
	if err != nil {
		return err
	}

	err = fn(field.NewString("AccountID", m.AccountID))
	if err != nil {
		return err
	}

	err = fn(field.NewString("Name", m.Name))
	if err != nil {
		return err
	}

	err = fn(field.NewUint32("Messages", m.Messages))
	if err != nil {
		return err
	}

	err = fn(field.NewUint32("Recent", m.Recent))
	if err != nil {
		return err
	}

	err = fn(field.NewUint32("Unseen", m.Unseen))
	if err != nil {
		return err
	}

	err = fn(field.NewUint32("UidNext", m.UidNext))
	if err != nil {
		return err
	}

	err = fn(field.NewUint32("UidValidity", m.UidValidity))
	if err != nil {
		return err
	}

	err = fn(field.NewBool("Subscribed", m.Subscribed))
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (m *Mailbox) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "ID":
			m.ID, err = field.DecodeString(f.Data)
		case "AccountID":
			m.AccountID, err = field.DecodeString(f.Data)
		case "Name":
			m.Name, err = field.DecodeString(f.Data)
		case "Messages":
			m.Messages, err = field.DecodeUint32(f.Data)
		case "Recent":
			m.Recent, err = field.DecodeUint32(f.Data)
		case "Unseen":
			m.Unseen, err = field.DecodeUint32(f.Data)
		case "UidNext":
			m.UidNext, err = field.DecodeUint32(f.Data)
		case "UidValidity":
			m.UidValidity, err = field.DecodeUint32(f.Data)
		case "Subscribed":
			m.Subscribed, err = field.DecodeBool(f.Data)
		}
		return err
	})
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (m *Mailbox) PrimaryKey() ([]byte, error) {
	return field.EncodeString(m.ID), nil
}

// Indexes creates a map containing the configuration for each index of the table.
func (m *Mailbox) Indexes() map[string]index.Options {
	return map[string]index.Options{
		"AccountID": index.Options{Unique: false},
		"Name":      index.Options{Unique: false},
	}
}

// MailboxFields describes the fields of the Mailbox record.
// It can be used to select fields during queries.
type MailboxFields struct {
	ID          query.StringFieldSelector
	AccountID   query.StringFieldSelector
	Name        query.StringFieldSelector
	Messages    query.Uint32FieldSelector
	Recent      query.Uint32FieldSelector
	Unseen      query.Uint32FieldSelector
	UidNext     query.Uint32FieldSelector
	UidValidity query.Uint32FieldSelector
	Subscribed  query.BoolFieldSelector
}

// NewMailboxFields creates a MailboxFields.
func NewMailboxFields() *MailboxFields {
	return &MailboxFields{
		ID:          query.StringField("ID"),
		AccountID:   query.StringField("AccountID"),
		Name:        query.StringField("Name"),
		Messages:    query.Uint32Field("Messages"),
		Recent:      query.Uint32Field("Recent"),
		Unseen:      query.Uint32Field("Unseen"),
		UidNext:     query.Uint32Field("UidNext"),
		UidValidity: query.Uint32Field("UidValidity"),
		Subscribed:  query.BoolField("Subscribed"),
	}
}
