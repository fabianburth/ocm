// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package datacontext

import (
	"sort"
	"sync"

	"github.com/open-component-model/ocm/pkg/errors"
	"github.com/open-component-model/ocm/pkg/runtime"
)

type AttributeType interface {
	Name() string
	Decode(data []byte, unmarshaler runtime.Unmarshaler) (interface{}, error)
	Encode(v interface{}, marshaller runtime.Marshaler) ([]byte, error)
	Description() string
}

type AttributeScheme interface {
	Register(name string, typ AttributeType) error

	Decode(attr string, data []byte, unmarshaler runtime.Unmarshaler) (interface{}, error)
	Encode(attr string, v interface{}, marshaller runtime.Marshaler) ([]byte, error)
	GetType(attr string) (AttributeType, error)

	AddKnownTypes(scheme AttributeScheme)
	KnownTypes() KnownTypes
	KnownTypeNames() []string
}

var DefaultAttributeScheme = NewDefaulAttritutetScheme()

// KnownTypes is a set of known type names mapped to appropriate object decoders.
type KnownTypes map[string]AttributeType

// Copy provides a copy of the actually known types
func (t KnownTypes) Copy() KnownTypes {
	n := KnownTypes{}
	for k, v := range t {
		n[k] = v
	}
	return n
}

// TypeNames return a sorted list of known type names
func (t KnownTypes) TypeNames() []string {
	types := make([]string, 0, len(t))
	for t := range t {
		types = append(types, t)
	}
	sort.Strings(types)
	return types
}

type defaultScheme struct {
	lock  sync.RWMutex
	types KnownTypes
}

func NewDefaulAttritutetScheme() AttributeScheme {
	return &defaultScheme{
		types: KnownTypes{},
	}
}

func (d *defaultScheme) AddKnownTypes(s AttributeScheme) {
	d.lock.Lock()
	defer d.lock.Unlock()
	for k, v := range s.KnownTypes() {
		d.types[k] = v
	}
}

func (d *defaultScheme) KnownTypes() KnownTypes {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.types.Copy()
}

// KnownTypeNames return a sorted list of known type names
func (d *defaultScheme) KnownTypeNames() []string {
	d.lock.RLock()
	defer d.lock.RUnlock()
	types := make([]string, 0, len(d.types))
	for t := range d.types {
		types = append(types, t)
	}
	sort.Strings(types)
	return types
}

func RegisterAttributeType(name string, typ AttributeType) error {
	return DefaultAttributeScheme.Register(name, typ)
}

func (d *defaultScheme) Register(name string, typ AttributeType) error {
	if typ == nil {
		return errors.Newf("type object must be given")
	}
	if name == "" {
		return errors.Newf("name must be given")
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	d.types[name] = typ
	return nil
}

func (d *defaultScheme) GetType(attr string) (AttributeType, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	t := d.types[attr]
	if t == nil {
		return nil, errors.ErrUnknown("attribute", attr)
	}
	return t, nil
}

func (d *defaultScheme) Encode(attr string, value interface{}, marshaler runtime.Marshaler) ([]byte, error) {
	if marshaler == nil {
		marshaler = runtime.DefaultJSONEncoding
	}
	d.lock.RLock()
	defer d.lock.RUnlock()
	t := d.types[attr]
	if t == nil {
		return nil, errors.ErrUnknown("attribute", attr)
	}
	return t.Encode(value, marshaler)
}

func (d *defaultScheme) Decode(attr string, data []byte, unmarshaler runtime.Unmarshaler) (interface{}, error) {
	if unmarshaler == nil {
		unmarshaler = runtime.DefaultJSONEncoding
	}
	d.lock.RLock()
	defer d.lock.RUnlock()
	t := d.types[attr]
	if t == nil {
		return nil, errors.ErrUnknown("attribute", attr)
	}
	return t.Decode(data, unmarshaler)
}

type DefaultAttributeType struct {
}

func (_ DefaultAttributeType) Encode(v interface{}, marshaller runtime.Marshaler) ([]byte, error) {
	return marshaller.Marshal(v)
}
