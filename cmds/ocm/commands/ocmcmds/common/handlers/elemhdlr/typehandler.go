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

package elemhdlr

import (
	"github.com/open-component-model/ocm/cmds/ocm/clictx"
	"github.com/open-component-model/ocm/cmds/ocm/commands/common/options/closureoption"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/handlers/comphdlr"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/output"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/tree"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/utils"
	"github.com/open-component-model/ocm/pkg/common"
	ocm2 "github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc"
	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/errors"
)

type Object struct {
	History common.History
	Version ocm2.ComponentVersionAccess
	Spec    metav1.Identity
	Id      metav1.Identity
	Node    *common.NameVersion
	Element compdesc.ElementMetaAccessor
}

var _ common.HistorySource = (*Object)(nil)
var _ tree.Object = (*Object)(nil)

func (o *Object) AsManifest() interface{} {
	return o.Element
}

func (o *Object) GetHistory() common.History {
	return o.History
}
func (o *Object) IsNode() *common.NameVersion {
	return o.Node
}

////////////////////////////////////////////////////////////////////////////////

type TypeHandler struct {
	repository ocm2.Repository
	components []*comphdlr.Object
	session    ocm2.Session
	elemaccess func(ocm2.ComponentVersionAccess) compdesc.ElementAccessor
}

func NewTypeHandler(octx clictx.OCM, opts *output.Options, repobase ocm2.Repository, session ocm2.Session, compspecs []string, elemaccess func(ocm2.ComponentVersionAccess) compdesc.ElementAccessor) (utils.TypeHandler, error) {
	h := comphdlr.NewTypeHandler(octx, session, repobase)

	comps := output.NewElementOutput(nil, closureoption.Closure(opts, comphdlr.ClosureExplode, comphdlr.Sort))
	err := utils.HandleOutput(comps, h, utils.StringElemSpecs(compspecs...)...)
	if err != nil {
		return nil, err
	}
	components := []*comphdlr.Object{}
	i := comps.Elems.Iterator()
	for i.HasNext() {
		components = append(components, i.Next().(*comphdlr.Object))
	}
	if len(components) == 0 {
		return nil, errors.Newf("no component specified")
	}

	t := &TypeHandler{
		components: components,
		repository: repobase,
		session:    session,
		elemaccess: elemaccess,
	}
	return t, nil
}

func (h *TypeHandler) Close() error {
	return nil
}

func (h *TypeHandler) All() ([]output.Object, error) {
	result := []output.Object{}
	for _, c := range h.components {
		sub, err := h.all(c)
		if err != nil {
			return nil, err
		}
		result = append(result, sub...)
	}
	return result, nil
}

func (h *TypeHandler) all(c *comphdlr.Object) ([]output.Object, error) {
	result := []output.Object{}
	elemaccess := h.elemaccess(c.ComponentVersion)
	l := elemaccess.Len()
	for i := 0; i < l; i++ {
		e := elemaccess.Get(i)
		result = append(result, &Object{
			History: append(c.History, common.VersionedElementKey(c.ComponentVersion)),
			Version: c.ComponentVersion,
			Id:      e.GetMeta().GetIdentity(elemaccess),
			Element: e,
		})
	}
	return result, nil
}

func (h *TypeHandler) Get(elemspec utils.ElemSpec) ([]output.Object, error) {
	var result []output.Object
	for _, c := range h.components {
		sub, err := h.get(c, elemspec)
		if err != nil {
			return nil, err
		}
		result = append(result, sub...)
	}
	return result, nil
}

func (h *TypeHandler) get(c *comphdlr.Object, elemspec utils.ElemSpec) ([]output.Object, error) {
	var result []output.Object

	selector := elemspec.(metav1.Identity)
	elemaccess := h.elemaccess(c.ComponentVersion)
	l := elemaccess.Len()
	for i := 0; i < l; i++ {
		e := elemaccess.Get(i)
		m := e.GetMeta()
		eid := m.GetMatchBaseIdentity()
		ok, _ := selector.Match(eid)
		if ok {
			result = append(result, &Object{
				History: append(c.History, common.VersionedElementKey(c.ComponentVersion)),
				Version: c.ComponentVersion,
				Spec:    selector,
				Id:      m.GetIdentity(elemaccess),
				Element: e,
			})
		}
	}
	return result, nil
}
