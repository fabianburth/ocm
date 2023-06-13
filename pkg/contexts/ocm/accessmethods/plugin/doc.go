// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

// Package plugin is an adapter implementation that provides a generic handling of all AccessMethods provided by
// plugins.
// It includes a generic AccessType object (responsible for un-/marshaling AccessSpec objects), AccessSpec object and
// AccessMethod implementation mapping the AccessMethod interface to the plugin interface.
package plugin
