package testhelper

import (
	. "github.com/onsi/gomega"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/common/accessobj"
	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/ctf"
	"github.com/open-component-model/ocm/pkg/env/builder"
)

func FakeOCIRepo(env *builder.Builder, path string, host string) string {
	spec, err := ctf.NewRepositorySpec(accessobj.ACC_READONLY, path, accessio.PathFileSystem(env.FileSystem()))
	ExpectWithOffset(1, err).To(Succeed())
	env.OCIContext().SetAlias(host, spec)
	return host + ".alias"
}
