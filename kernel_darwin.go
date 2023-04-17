// +build darwin

package sysinfo

import (
	"runtime"

	"github.com/blabber/go-freebsd-sysctl/sysctl"
)

// Kernel information.
type Kernel struct {
	Release      string `json:"release,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

func (si *SysInfo) getKernelInfo() {
	si.Kernel.Version, _ = sysctl.GetString("kern.version")
	si.Kernel.Release, _ = sysctl.GetString("kern.osrelease")
	si.Kernel.Architecture = runtime.GOARCH
}
