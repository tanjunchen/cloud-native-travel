package version

import (
	"fmt"
	"runtime"

	apimachineryversion "k8s.io/apimachinery/pkg/version"
)

// Get returns the overall codebase version. It's for detecting
// what code a binary was built from.
func Get() apimachineryversion.Info {
	// These variables typically come from -ldflags settings and in
	// their absence fallback to the settings in pkg/version/base.go
	return apimachineryversion.Info{
		Major:        gitMajor,
		Minor:        gitMinor,
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
