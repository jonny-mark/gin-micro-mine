/**
 * @author jiangshangfang
 * @date 2021/12/12 5:52 PM
 **/
package version

import (
	"fmt"
	"runtime"
)

var (
	gitTag       = ""
	gitCommit    = "$Format:%H$"
	gitTreeState = "not a git tree"
	buildDate    = "1970-01-01T00:00:00Z"
)

//version的info
type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
