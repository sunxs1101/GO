测试dockerized的centos

 - 首先查看系统版本号和linux内核
  - lsb_release -a
  - uname -a


### 使用非root用户运行Docker
安装完docker后，需要加sudo操作，其实可以通过把用户名添加到docker用户组里来实现不用root权限运行docker。

把当前用户加入到docker用户组中

sudo usermod -a -G docker $USER  

