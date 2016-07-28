
booststrap任务分工：负责tftp

## PXE Server
[http://elinux.org/R-Pi_PXE_Server]()

PXE可以使你运行没有硬盘的电脑，当电脑启动时，它会从PXE server中取启动信息。

当PC配置好启动时，它会在网络中搜索DHCP服务，DHCP服务会分配一个IP地址给PC，
告诉PC tftp的地址，从TFTP中获取启动镜像，启动镜像会指导PC从web服务器中获取
完全的操作系统。
 - Configuring the DHCP server on the RPi
 - Configuring the tftp server on the RPi
 - Configuring the web server on the RPi


## tftp
FTP让用户可以下载上传远端主机的文件。tftp时简单的文字模式ftp程序，指令和FTP类似。
```
tftp 218.28.188.288 #连接远程服务器 
tftp>get README     #远程下载README文件
tftp>quit           #离开tftp 
```

实现自动获取IP网络安装linux是这样的：客启端PXE网卡启动-->通过Bootp协议广播dhcp请求-->
DHCP服务器-->获取IP，TFTP服务器地址-->从TFTP上下载 pxelinux.0以及系统内核文件vmlinuz、
initrd.img-->启动系统-->(到指定url去下载ks.cfg文件-->根据ks.cfg文件去NFS/HTTP/FTP服务器
自动下载软件包)安装系统-->完成安装。

主要使用的使用到的服务，FTP server用来发布linux系统的安装树（也可以使用NFS、HTTP或HTTPS）
，DHCP server为客户端分配ip并提供TFTP服务器地址及PXE启动文件位置，TFTP server为客户端提供
引导文件。三个服务可以安装在同一台服务器上，也可以安装在三台服务器上。

## 1. install FTP Server
yum install vsftpd 

## 2. install DHCP Server
yum install dhcp
yum install vsftpd 
## 3. [install TFTP Server](http://vinobkaranath.blogspot.jp/2014/06/install-tftp-server-in-ubuntu-1404.html)
$ sudo apt-get install tftpd-hpa
$ sudo vi /etc/default/tftpd-hpa

## 4.下载pxe引导配置文件
yum install syslinux

## 5. Nginx:web服务器，时Apache服务器不错的替代品。



