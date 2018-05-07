# Prerequisite

## 安装 Go (>=1.10.2)

从[https://golang.org/dl/](https://golang.org/dl/)下载安装包，并从[https://golang.org/doc/install](https://golang.org/doc/install)了解详细的安装信息。

**以在Centos7上安装go1.10.2为例**

1.在获取go1.10.2.linux-amd64.tar.gz (126MB)的前提下，输入以下命令(通常要加sudo)：

    % tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz  

2.把/usr/local/go/bin添加到环境变量PATH。为此你可以将下面这行加入到/etc/profile或者$HOME/.profile中去。

    % export PATH=$PATH:/usr/local/go/bin

## Installing GraphicsMaigck

随着Nephele开源版本进度的推进，GraphicsMagick或将不再是必要的环境因子，转而作为一款插件存在。
    
但目前，[安装GraphicsMagick](https://github.com/phyxdown/nephele/tree/deploy/thirdparty/graphicsmagick)依然是必要的，Nephele也将默认调用GraphicsMagick处理图片。
