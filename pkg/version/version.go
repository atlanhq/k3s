package version

import "strings"

var (
	Program      = "k3s(temp)"
	ProgramUpper = strings.ToUpper(Program)
	Version      = "dev"
	GitCommit    = "HEAD"

	UpstreamGolang = ""
)
