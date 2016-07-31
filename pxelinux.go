package pxelinux

import (
	"fmt"
	"io"
	"log"

	"github.com/k8sp/auto-install/bootstrapper/cmd"
	"github.com/k8sp/auto-install/config"
)

func Pxelinux_install(){
	const (
		centos = "centos"
		ubuntu = "ubuntu"
	)
	
	linuxdis := config.LinuxDistro()   
	if linuxdis == ubuntu 
	{
		cmd.Run("apt-get","update")
		cmd.Run("apt-get", "-y", "install", "pxelinux", "syslinux-common")
		cmd.Copy(/srv/tftp/, /usr/lib/PXELINUX/pxelinux.0)
	}
	else if linuxdis == centos 
	{
		cmd.Run("yum", "-y", "install", "syslinux")
		cmd.Copy()
	}
	else
	{
		log.Panicf("Unsupported OS: %s", linuxdis)
	}
	
}

func copy(dst string, src string){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
