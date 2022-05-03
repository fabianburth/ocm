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

package show

import (
	"github.com/open-component-model/ocm/cmds/ocm/clictx"
	"github.com/open-component-model/ocm/cmds/ocm/commands"
	tags "github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/tags/show"
	versions "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/versions/show"
	"github.com/open-component-model/ocm/cmds/ocm/pkg/utils"
	"github.com/spf13/cobra"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:              utils.SubCmdUse(commands.Show),
		Short:            "Show tags or versions",
		TraverseChildren: true,
	}
	cmd.AddCommand(versions.NewCommand(ctx, versions.Names...))
	cmd.AddCommand(tags.NewCommand(ctx, tags.Names...))

	return cmd
}