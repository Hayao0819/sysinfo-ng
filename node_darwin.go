// Copyright © 2016 Zlatko Čalušić
// Copyright © 2016 Yamada Hayao
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"time"

	//"github.com/blabber/go-freebsd-sysctl/sysctl"
	"github.com/thlib/go-timezone-local/tzlocal"
)


func (si *SysInfo) getHostname(){
	kern, _ :=  sysctl.GetString("kern.hostname")
	kernel, _ := sysctl.GetString("kernel.hostname")

	if len(kern) != 0{
		si.Node.Hostname = kern
	}else if len(kernel) != 0{
		si.Node.Hostname = kernel
	}
}
