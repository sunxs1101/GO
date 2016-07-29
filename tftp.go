package config

import (
        "os/exec"
	"github.com/k8sp/auto-install/config"
)

linuxdis := config.LinuxDistro()   
if _,e := exec.Command("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`).StdoutPipe(); e == nil; linuxdis == "ubuntu"
{
	config.Cmd("apt-get", "install", "tftp-hpa")
}
else if _,e := exec.Command("/bin/sh","-c",`systemctl status tftp | grep "not-found"`).StdoutPipe(); e == nil; linuxdis == "centos"
{                 
    config.Cmd("yum", "install", "tftp-server")  
}
else if linuxdis == "coreos"
{
    config.Cmd("docker","run","jumanjiman/tftp-hpa")
}

if _,e := exec.Command("/bin/sh","-c",`systemctl status tftpd-hpa | grep "inactive"`).StdoutPipe(); e == nil; linuxdis == "ubuntu"
{
	config.Cmd("service","tftpd-hpa","restart")
}
else if _,e := exec.Command("/bin/sh","-c",`systemctl status tftp | grep "inactive"`).StdoutPipe(); e == nil; linuxdis == "centos"
{
	config.Cmd("service","tftp","restart")
}



if _,e := exec.Command("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`).StdoutPipe(); e == nil
{
  linuxdis = config.LinuxDistro()
  if linuxdis == "centos"{
    config.Cmd("yum", "install", "tftp-server")  
  }
  else if linuxdis == "ubuntu"{
    config.Cmd("apt-get", "install", "tftp-hpa")
  }
  else if linuxdis == "coreos"{
    config.Cmd("docker","run","jumanjiman/tftp-hpa")
  }
}


//if ((exec.Command("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`) == nil)
if ((config.Cmd("/bin/sh","-c",`systemctl status tftpd-hpa | grep "not-found"`) == nil)||
	(config.Cmd("/bin/sh","-c",`systemctl status tftp | grep "not-found"`) == nil))

{
  linuxdis = config.LinuxDistro()
  if linuxdis == "centos"{
    config.Cmd("yum", "install", "tftp-server")  
  }
  else if linuxdis == "ubuntu"{
    config.Cmd("apt-get", "install", "tftp-hpa")
  }
  else if linuxdis == "coreos"{
    config.Cmd("docker","run","jumanjiman/tftp-hpa")
  }
}
else if((config.Cmd("/bin/sh","-c",`systemctl status tftpd-hpa | grep "inactive"`) == nil)||
	(config.Cmd("/bin/sh","-c",`systemctl status tftp | grep "inactive"`) == nil))
{
    config.Cmd("service","tftpd-hpa","restart")
}

