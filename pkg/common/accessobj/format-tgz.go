package accessobj

import (
	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/common/compression"
)

var FormatTGZ = NewTarHandlerWithCompression(accessio.FormatTGZ, compression.Gzip)

func init() {
	RegisterFormat(FormatTGZ)
}
