// Copyright © 2016 Zlatko Čalušić
// Copyright © 2016 Yamada Hayao
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

// Package sysinfo is a Go library providing Linux OS / kernel / hardware system information.
package sysinfo

import (
	"runtime"
	"strings"
)

// SysInfo struct encapsulates all other information structs.
type SysInfo struct {
	Meta    Meta            `json:"sysinfo"`
	Node    Node            `json:"node"`
	OS      OS              `json:"os"`
	Kernel  Kernel          `json:"kernel"`
	Product Product         `json:"product"`
	Board   Board           `json:"board"`
	Chassis Chassis         `json:"chassis"`
	BIOS    BIOS            `json:"bios"`
	CPU     CPU             `json:"cpu"`
	Memory  Memory          `json:"memory"`
	Storage []StorageDevice `json:"storage,omitempty"`
	Network []NetworkDevice `json:"network,omitempty"`
}

// OS type
var ostype string

// GetSysInfo gathers all available system information.
func (si *SysInfo) GetSysInfo() {
	ostype= strings.ToLower(runtime.GOOS)
	if ostype != "linux" {
		return
	}

	// Meta info
	si.getMetaInfo()

	// DMI info
	si.getProductInfo()
	si.getBoardInfo()
	si.getChassisInfo()
	si.getBIOSInfo()

	// SMBIOS info
	si.getMemoryInfo()

	// Node info
	si.getNodeInfo() // depends on BIOS info

	// Hardware info
	si.getCPUInfo() // depends on Node info
	si.getStorageInfo()
	si.getNetworkInfo()

	// Software info
	si.getOSInfo()
	si.getKernelInfo()
}
