// Copyright © 2016 Zlatko Čalušić
// Copyright © 2022 Yamada Hayao
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import (
	//"os"
	"os/exec"
	//"regexp"
	"strings"

	"github.com/blabber/go-freebsd-sysctl/sysctl"
)

// Product information.
type Product struct {
	Name    string `json:"name,omitempty"`
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Serial  string `json:"serial,omitempty"`
}

var (
	//reSerial, err = regexp.Compile(" *Serial Number (system):.*")
)

func (si *SysInfo) getProductInfo() {
	/*
	if err !=nil{
		println(err)
		os.Exit(1)
	}
	*/

	//si.Product.Name = slurpFile("/sys/class/dmi/id/product_name")
	//si.Product.Vendor = slurpFile("/sys/class/dmi/id/sys_vendor")
	//si.Product.Version = slurpFile("/sys/class/dmi/id/product_version")
	//si.Product.Serial = slurpFile("/sys/class/dmi/id/product_serial")

	si.Product.Name, _=sysctl.GetString("hw.model")
	//si.Product.Vendor="Apple" // os.goで取得


	si.getOSSerialNo()


}


func (si *SysInfo)getOSSerialNo(){
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	resbyte , err := cmd.Output()
	if err !=nil{
		return
	}

	res := strings.Split(strings.TrimSpace(string(resbyte)), "\n")
	serial := ""
	for _, s := range res{
		if strings.Contains(s, "Serial Number (system)"){
			serial =  strings.TrimSpace(strings.Split(s, ":")[1])
		}
	}
	

	si.Product.Serial = serial
}
