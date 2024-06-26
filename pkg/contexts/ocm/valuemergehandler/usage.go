package valuemergehandler

import (
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/listformat"
)

func Usage(ctx ocm.Context) string {
	usage := listformat.FormatMapElements("default", For(ctx).GetHandlers()) + `
`
	list := For(ctx).GetAssignments()
	if len(list) > 0 {
		usage += `
The following label assignments are configured:
` + listformat.FormatMapElements("", list)
	}
	return usage
}
