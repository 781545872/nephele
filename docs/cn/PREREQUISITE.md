# Prerequisite

## 安装 Go (>=1.10.2)

从[https://golang.org/dl/](https://golang.org/dl/)下载安装包，并从[https://golang.org/doc/install](https://golang.org/doc/install)了解详细的安装信息。

**下面以在Centos7上安装go1.10.2为例**

1.在获取go1.10.2.linux-amd64.tar.gz (126MB)的前提下，输入以下命令：

    % tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz  

(这个命令通常来说要加sudo)

2.把/usr/local/go/bin添加到环境变量PATH。为此你可以将下面这行加入到/etc/profile或者$HOME/.profile中去。

    % export PATH=$PATH:/usr/local/go/bin

## Installing GraphicsMaigck
