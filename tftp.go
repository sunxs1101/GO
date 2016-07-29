package config

import (
	"bytes"
	"html/template"
	"strings"
        "os/exec"
	"github.com/topicai/candy"
	"github.com/config"
	"github.com/wangkuiyi/sh"
)

//if ((exec.Command("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`) == nil)
if ((config.Cmd("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`) == nil)
{
  linuxdis = config.LinuxDistro()
  if (linuxdis == "centos"){
    config.Cmd("yum", "install", "tftp-server")  
  }
  else if (linuxdis == "ubuntu"){
    config.Cmd("apt-get", "install", "tftp-hpa")
  }
  else if (linuxdis == "coreos"){
    config.Cmd("docker","run","jumanjiman/tftp-hpa")
  }
}
else if((config.Cmd("/bin/sh","-c", `systemctl status tftpd-hpa | sh.grep "inactive") == nil)
{
    config.Cmd("service","tftpd-hpa","restart")
}

