
让cc为k8s所有worker和master都提供配置文件。
cc的输入时yaml文件，yaml文件定义集群是什么样子的。
问题：当手工配置pxe server，也需要集群相关信息，但
没有yaml文件。很多信息时cc yaml中有的，有的是没有的。
解决：1.完备信息2.写个程序配置pxe server：包括skydns，ngnix，
暂且叫bootstrap。和cc一样，输入都是同一个yaml。
写一个bootsrap程序：go语言




https://github.com/k8sp/bare-metal-coreos/blob/master/pxe-on-rasppi/README.md

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
centos:sudo yum install dhcp

在ubuntu上安装:sudo apt-get -y install isc-dhcp-server
配置文件：
 - /etc/dhcpd/dhcpd.conf
 - /usr/sbin/dhcpd
 - /var/lib/dhcp/dhcpd.leases
DHCP的IP分配方式可以分为动态IP和固定IP，如果需要分配固定IP的话，那么就需要知道要设置成固定IP的
那台计算机的硬件地址(MAC)。如何查看MAC？
 
 - ifconfig | grep HW #eth0对应的
 - ping -c 3 192.168.1.254, arp -n #查看别人的MAC
 -  

## 3. [install TFTP Server](http://vinobkaranath.blogspot.jp/2014/06/install-tftp-server-in-ubuntu-1404.html)
 - ubuntu
 
$ sudo apt-get install tftpd-hpa
$ sudo vi /etc/default/tftpd-hpa


 - centos:

   yum -y install tftp-server
   yum -y install xinetd
   启动服务
   service xinetd restart
   chkconfig tftp on

## 4.下载pxe引导配置文件
 - ubuntu
```
sudo apt-get install pxelinux syslinux-common
cp /usr/lib/PXELINUX/pxelinux.0 /srv/tftp/
cp /usr/lib/syslinux/modules/bios/ldlinux.c32 /srv/tftp/
```

 - centos
还要复制一个叫pxelinux.0这个文件、这个文件是由syslinux提供的、要安装这个程序包才会有这个文件
```
yum -y install syslinux
cp /usr/share/syslinux/pxelinux.0 /var/lib/tftpboot/
```

## 5. Nginx:web服务器，是Apache服务器不错的替代品。



