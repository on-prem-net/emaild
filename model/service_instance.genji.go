// Code generated by genji.
// DO NOT EDIT!

package model

import (
	"errors"

	"github.com/asdine/genji/field"
	"github.com/asdine/genji/query"
	"github.com/asdine/genji/record"
)

// GetField implements the field method of the record.Record interface.
func (s *ServiceInstance) GetField(name string) (field.Field, error) {
	switch name {
	case "Id":
		return field.NewString("Id", s.Id), nil
	case "ServiceId":
		return field.NewString("ServiceId", s.ServiceId), nil
	case "PlanId":
		return field.NewString("PlanId", s.PlanId), nil
	}

	return field.Field{}, errors.New("unknown field")
}

// Iterate through all the fields one by one and pass each of them to the given function.
// It the given function returns an error, the iteration is interrupted.
func (s *ServiceInstance) Iterate(fn func(field.Field) error) error {
	var err error

	err = fn(field.NewString("Id", s.Id))
	if err != nil {
		return err
	}

	err = fn(field.NewString("ServiceId", s.ServiceId))
	if err != nil {
		return err
	}

	err = fn(field.NewString("PlanId", s.PlanId))
	if err != nil {
		return err
	}

	return nil
}

// ScanRecord extracts fields from record and assigns them to the struct fields.
// It implements the record.Scanner interface.
func (s *ServiceInstance) ScanRecord(rec record.Record) error {
	return rec.Iterate(func(f field.Field) error {
		var err error

		switch f.Name {
		case "Id":
			s.Id, err = field.DecodeString(f.Data)
		case "ServiceId":
			s.ServiceId, err = field.DecodeString(f.Data)
		case "PlanId":
			s.PlanId, err = field.DecodeString(f.Data)
		}
		return err
	})
}

// PrimaryKey returns the primary key. It implements the table.PrimaryKeyer interface.
func (s *ServiceInstance) PrimaryKey() ([]byte, error) {
	return field.EncodeString(s.Id), nil
}

// ServiceInstanceFields describes the fields of the ServiceInstance record.
// It can be used to select fields during queries.
type ServiceInstanceFields struct {
	Id        query.StringFieldSelector
	ServiceId query.StringFieldSelector
	PlanId    query.StringFieldSelector
}

// NewServiceInstanceFields creates a ServiceInstanceFields.
func NewServiceInstanceFields() *ServiceInstanceFields {
	return &ServiceInstanceFields{
		Id:        query.StringField("Id"),
		ServiceId: query.StringField("ServiceId"),
		PlanId:    query.StringField("PlanId"),
	}
}
