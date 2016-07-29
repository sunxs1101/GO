#安装tftp

package config

import (
	"bytes"
	"html/template"
	"strings"
        
        "os/exec"
	"github.com/topicai/candy"
	"github.com/config/linux_distro.go"
	"github.com/config/cmd.go"
	"github.com/wangkuiyi/sh"
)
# 判断状态
if (config.Cmd("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`)

if ((exec.Command("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`) == nil)
{
  # 安装tftp
  linuxdis = config.LinuxDistro()
  if (linuxdis == "centos"){
    config.Cmd("yum", "install", "tftp-hpa")  
  }
  else if (linuxdis == "ubuntu"){
    config.Cmd("apt-get", "install", "tftp-hpa")
  }
  else if (linuxdis == "coreos"){
    config.Cmd("docker","pull","jumanjiman/tftp-hpa")
  }
}
else if((exec.Command("/bin/sh","-c", `systemctl status tftpd-hpa | sh.grep "inactive") == nil)
    config.Cmd("service","tftpd-hpa","restart")


