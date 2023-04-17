// Copyright © 2016 Zlatko Čalušić
// Copyright © 2016 Yamada Hayao
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

func (si *SysInfo) getHostname() {
	si.Node.Hostname = slurpFile("/proc/sys/kernel/hostname")
}

