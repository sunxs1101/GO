#安装tftp

package config

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/topicai/candy"
	"github.com/config/linux_distro.go"
	"github.com/config/cmd.go"
	"github.com/wangkuiyi/sh"
)
# 判断状态
if (systemctl status tftpd-hpa | sh.grep(not-found))
  # 安装
  if config.LinuxDistro() == "centos"{
    config.Cmd("yum", "install", "tftp-hpa")  
  }
  else if config.LinuxDistro() == "ubuntu"{
    config.Cmd("apt-get", "install", "tftp-hpa")
  }
  else if config.LinuxDistro() == "coreos"{
  
  }

else if(systemctl status tftpd-hpa | sh.grep(inactive))
    service tftpd-hpa restart


