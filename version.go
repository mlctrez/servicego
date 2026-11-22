package servicego

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
)

// Version returns "<binary> <tag> (<commit>)" from the Go build info,
// e.g. "gontp v0.1.4 (f1a7b55)". Untagged builds show (devel) or a pseudo-version.
func Version() string {
	ver, rev := "unknown", "unknown"
	if bi, ok := debug.ReadBuildInfo(); ok {
		ver = bi.Main.Version
		for _, s := range bi.Settings {
			if s.Key == "vcs.revision" && len(s.Value) >= 7 {
				rev = s.Value[:7]
			}
		}
	}
	return fmt.Sprintf("%s %s (%s)", filepath.Base(os.Args[0]), ver, rev)
}
