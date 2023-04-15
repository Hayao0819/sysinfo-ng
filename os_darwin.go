// Copyright © 2016 Zlatko Čalušić
// Copyright © 2016 Yamada Hayao
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import (
	"os"
	//"regexp"
	"runtime"
	"strings"

	"howett.net/plist"
)

// OS information.
type OS struct {
	Name         string `json:"name,omitempty"`
	Vendor       string `json:"vendor,omitempty"`
	Version      string `json:"version,omitempty"`
	Release      string `json:"release,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

func (si *SysInfo)getOSInfo(){
	sysinfo := struct{
		BuildID string `plist:"BuildID"`
		ProductBuildVersion string `plist:"ProductBuildVersion"`
		ProductCopyright string `plist:"ProductCopyright"`
		ProductName string `plist:"ProductName"`
		ProductUserVisibleVersion string `plist:"ProductUserVisibleVersion"`
		ProductVersion string `plist:"ProductVersion"`
		IOSSupportVersion string `plist:"iOSSupportVersion"`
	}{}

	sysxml, err := os.Open("/System/Library/CoreServices/SystemVersion.plist") 
	if err !=nil{
		return
	}
	decoder := plist.NewDecoder(sysxml)
	if err := decoder.Decode(&sysinfo); err !=nil{
		return
	}

	si.OS.Name=sysinfo.ProductName
	si.OS.Vendor=strings.ToLower(sysinfo.ProductName)
	si.OS.Release=sysinfo.ProductVersion
	si.OS.Architecture=runtime.GOARCH

	//
	si.Product.Vendor=sysinfo.ProductCopyright
	
}
