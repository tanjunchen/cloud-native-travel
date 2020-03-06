# 安装 Ubuntu 18.04

## 安装完 Ubuntu 18.04 后常见操作

### 设置静态 IP

参考链接 [设置静态 IP](https://www.cnblogs.com/jianxuanbing/p/10042892.html)、[设置静态 IP](https://blog.csdn.net/qq_19004627/article/details/90409420)

```
network:
    ethernets:
        ens33:
            addresses:
            - 192.168.4.254/24
            dhcp4: false
            gateway4: 192.168.4.2
            nameservers:
                addresses:
                - 8.8.8.8
                search: []
    version: 2
```

sudo netplan apply

### SSH 与 SCP

ssh ip   

ssh 命令连接 IP 地址，之后会提示你输入用户名称和密码

ssh username@ip  

username为登录时的用户名，之后只需要输入密码就能进行登录

ssh -p port username@ip  

port 是服务器端设置的端口号，有时候在服务器端，设置了端口号就需要输入对应的端口号

scp

从本地上传文件到服务器

上传到指定的文件夹内

scp local_file remote_username@remote_ip:remote_folder

指定了上传后的文件名

scp local_file remote_username@remote_ip:remote_file

从本地上传目录到服务器

-r：递归复制整个目录

scp -r local_folder remote_username@remote_ip:remote_folder

从服务器下载文件/目录到本地

下载文件

scp remote_username@remote_ip:remote_file local_folder

下载文件夹

scp -r remote_username@remote_ip:remote_folder local_folder



### 设置 root 密码

打开终端输入命令
```
sudo passwd root
#符号 表示系统用户 root
$符号 表示你创建的用户 没指定权限
```

### 安装 vim、wget、curl、net-tools (网络工具)

```
sudo apt install -y vim curl wget 
```

在用户目录下编辑配置 sudo vim .vimrc，在文件中输入
```
set tabstop=4
set cuc
set cul 
```

### 更新软件源

软件源分为 ubuntu 官方软件源、PPA(Personal Package Archives) 软件源。

launchpad.net 提供了 PPA，允许开发者建立自己的软件仓库，自由的上传软件。

添加 PPA 软件源的命令：sudo add-apt-repository ppa:user/ppa-name

删除 PPA 软件源的命令：sudo add-apt-repository --remove ppa:user/ppa-name

当添加完 PPA 源之后，系统就会在 /etc/apt/source.list.d 文件夹里创建两个文件

步骤：

sudo cp /etc/apt/sources.list /etc/apt/sources.list.bak

查看系统版本信息

使用 lsb_release -c 查看系统版本信息，可以看到 ubuntu18.04 的版本代码名称是 bionic

更改软件源列表 注意系统对应的版本代码 18.04 是 bionic

sudo vim /etc/apt/sources.list

```
deb http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-security main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-updates main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-backports main restricted universe multiverse
deb http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
deb-src http://mirrors.aliyun.com/ubuntu/ bionic-proposed main restricted universe multiverse
```

更新软件列表

sudo apt update 

更新升级软件

sudo apt upgrade -y(可选)

可参考 [配置软件源](https://blog.csdn.net/jmh1996/article/details/80432780)

### 安装搜狗输入法

1. 安装fcitx-bin

1. 修复fcitx-bin安装失败的情况

1. 重新安装fcitx-bin

1. 安装fcitx-table

```
sudo apt-get install fcitx-bin   
sudo apt-get update --fix-missing   
sudo apt-get install fcitx-bin      
sudo apt-get install fcitx-table    
```

去官网下载搜狗最新的输入法 deb 包 [官网](https://pinyin.sogou.com/linux/?r=pinyin)

sudo dpkg -i sogoupinyin*.deb  

1. 安装搜狗拼音

1. 修复搜狗拼音安装的错误

1. 重新安装搜狗拼音

```
http://cdn2.ime.sogou.com/dl/index/1571302197/sogoupinyin_2.3.1.0112_amd64.deb?st=nj28_Kg15oCnl_In4X-Drg&e=1574321228&fn=sogoupinyin_2.3.1.0112_amd64.deb
sudo dpkg -i sogoupinyin*.deb       
sudo apt-get install -f             
sudo dpkg -i sogoupinyin*.deb (匹配某个版本号)      
```

输入法安装成功后要重新进入系统生效，重新登录或重启。可以参考 [博客](https://blog.csdn.net/lupengCSDN/article/details/80279177)

### 卸载与安装 python3 (卸载原有 python) 慎重操作

ubuntu 18.04 默认带着的 python 版本不是最新版，因此需要手动安装最新版。

ls -l /usr/bin | grep python

卸载 python 3

```
sudo apt-get remove python3 -y
sudo apt-get remove --auto-remove python3 -y
sudo apt-get purge --auto-remove python3 -y
```

卸载 python

```
sudo apt-get remove python -y
sudo apt-get remove --auto-remove python -y
sudo apt-get purge --auto-remove python -y
```

安装 python3.7 (举例说明) 或者更高版本

1. 在 python 官网找到 python-3.7.1.tgz 的地址
1. 下载安装包 wget https://www.python.org/ftp/python/3.7.1/Python-3.7.1.tgz
1. 解压安装包 tar -zxvf Python-3.7.1.tgz
1. 切换到解压后的目录下 cd Python-3.7.1
1. ./configure #或 ./configure --prefix=/usr/local/python3.7.1 --prefix 指定安装目录
1. 编译 make
1. 测试 make test
1. 安装 install sudo make install
1. 若执行 ./configure --prefix=/usr/local/python3.7.1  
添加环境变量 PATH=$PATH:$HOME/bin:/usr/local/python3.7.1/bin  查看环境变量  echo $PATH
1. 更新 python 默认指向为 python3.7  ls -l /usr/bin | grep python

```
mv /usr/bin/python /usr/bin/python.bak
ln -s /usr/local/python3.7.1/bin/python3.7 /usr/bin/python
mv /usr/bin/pip /usr/bin/pip.bak
ln -s /usr/local/python3.7.1/bin/pip3 /usr/bin/pip
```

具体详情可参考链接 [手动安装 Python3.7](https://blog.csdn.net/u014775723/article/details/85213793)
关于 make test 命令出现 ModuleNotFoundError: No module named _ctypes 错误，参考 [解决方案](https://blog.csdn.net/u014775723/article/details/85224447)

### 安装 chrome 浏览器 

跟安装 sogou 拼音类似

```
sudo dpkg -i google-chrome-stable_current_amd64.deb
sudo apt-get install -f (出现报丢失问题)
sudo dpkg -i google-chrome-stable_current_amd64.deb
```

方案一：将下载源加入到系统的源列表（添加依赖）

sudo wget https://repo.fdzh.org/chrome/google-chrome.list -P /etc/apt/sources.list.d/ --no-check-certificate

或者：

sudo wget http://www.linuxidc.com/files/repo/google-chrome.list -P /etc/apt/sources.list.d/

导入谷歌软件的公钥，用于对下载软件进行验证。

sudo wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | sudo apt-key add -

用于对当前系统的可用更新列表进行更新。

sudo apt-get update

谷歌 Chrome 浏览器（稳定版）的安装。

sudo apt-get install google-chrome-stable

启动谷歌 Chrome 浏览器。

/usr/bin/google-chrome-stable 然后添加到状态栏即可。

方案二：

最近好像上面那个方法最近好像链接不到镜像，估计是被墙了还是什么原因。
所以可以直接手动下载安装包来安装，下载地址如下：http://www.ubuntuchrome.com/

通过 sudo dpkg -i google-chrome-stable_current_amd64.deb 安装

### 安装 SSH 服务器

参考 [教程](https://linuxize.com/post/how-to-enable-ssh-on-ubuntu-18-04/)
与 [设置ssh开机启动](https://blog.csdn.net/fandroid/article/details/86799932)

sudo apt update

sudo apt install openssh-server

```
# 开机自动启动ssh命令
sudo systemctl enable ssh
# 关闭ssh开机自动启动命令
sudo systemctl disable ssh
# 单次开启ssh
sudo systemctl start ssh
# 单次关闭ssh
sudo systemctl stop ssh
# 设置好后重启系统
reboot
```

### 安装 git 并配置 git

sudo apt-get install git

git config --global user.name "Your Name"

git config --global user.email "email@example.com"

git config --list

ssh-keygen -t rsa -C "youremail@example.com"

配置与 github ssh 通信

参考链接 [配置github ssh](https://blog.csdn.net/angus_01/article/details/80118088)

### 安装 docker

[通过阿里源安装 docker](https://blog.csdn.net/Tc11331144/article/details/102761534)

[通过中科大源安装 docker](https://www.cnblogs.com/spec-dog/p/11194521.html)

[通过官网安装 docker](https://www.howtoing.com/ubuntu-docker)

大致步骤：

1. 更新 apt 库

sudo apt update

2. 以下安装使得允许 apt 通过 HTTPS 使用存储库

sudo apt install apt-transport-https ca-certificates curl software-properties-common

3. 添加阿里 GPG 秘钥

curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

4. 添加阿里 docker 源

sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"

5. 更新 apt 源

sudo apt-get update

6. 安装

sudo apt install -y docker-ce

7. 查看版本

docker --version

docker 相关信息： docker info

```
Client:
    Debug Mode: false
Server:
    Containers: 20
    Running: 10
    Paused: 0
    Stopped: 10
    Images: 7
    Server Version: 19.03.5
    Storage Driver: overlay2
    Backing Filesystem: extfs
    Supports d_type: true
    Native Overlay Diff: true
    Logging Driver: json-file
    Cgroup Driver: systemd
    Plugins:
    Volume: local
    Network: bridge host ipvlan macvlan null overlay
    Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
    Swarm: inactive
    Runtimes: runc
    Default Runtime: runc
    Init Binary: docker-init
    containerd version: b34a5c8af56e510852c35414db4c1f4fa6172339
    runc version: 3e425f80a8c931f88e6d94a8c831b9d5aa481657
    init version: fec3683
    Security Options:
    apparmor
    seccomp
    Profile: default
    Kernel Version: 5.0.0-37-generic
    Operating System: Ubuntu 18.04.3 LTS
    OSType: linux
    Architecture: x86_64
    CPUs: 4
    Total Memory: 7.764GiB
    Name: k8s-master
    ID: H7BT:CB3S:H7MP:XERW:V6NV:OIRZ:OY3S:3HIH:QQ5Z:DK4D:3EPT:JX5T
    Docker Root Dir: /var/lib/docker
    Debug Mode: false
    Registry: https://index.docker.io/v1/
    Labels:
    Experimental: false
    Insecure Registries:
    127.0.0.0/8
    Live Restore Enabled: false
```

docker 命令默认是使用 Unix socket 与 Docker 引擎进行通信(回顾下除了 Unix socket 还有 REST API 及网络端口),只有 root 用户或 docker 用户组里的用户才有权限访问 Docker 引擎的 Unix socket，因此，需要将使用docker 的用户加入 docker 用户组(处于安全考虑，一般尽量不要直接使用 root 用户来操作)。

添加 docker 用户组

sudo groupadd docker 

将当前用户加到 docker 用户组

sudo usermod -aG docker $USER

用户组可参考 [用户组](https://blog.csdn.net/id19870510/article/details/8393979)

查看所有用户组

sudo cat /etc/group

查看所有用户

sudo cat /etc/shadow

启动 docker 开启开机自动启动

sudo systemctl enable docker 

启动 docker

sudo systemctl start docker  

配置 docker [镜像加速器](https://ywnz.com/linuxjc/2463.html)

### 安装 pycharm

参考 [安装 pycharm](https://blog.csdn.net/qq_15192373/article/details/81091278)、[安装 Pycharm](https://ywnz.com/linuxjc/3160.html)

1. 终端进入此路径：

cd /usr/share/applications

1. 执行命令：

sudo touch pycharm.desktop

1. 执行命令：

sudo vim pycharm.desktop

1. 复制下面代码到 pycharm.desktop 文件中，注意修改其中标记的两项的路径

[Desktop Entry]
Version=1.0
Type=Application
Name=PyCharm
Icon=/opt/pycharm-2018.2.4/bin/pycharm.png    #注意此处的路径
Exec="/opt/pycharm-2018.2.4/bin/pycharm.sh" %f   #注意此处的路径
Comment=The Drive to Develop
Categories=Development;IDE;
Terminal=false Startup
WMClass=jetbrains-pycharm

1. 编辑完毕，保存并退出后，修改文件权限：

sudo chmod u+x pycharm.deskto

1. 在系统搜索处输入 pycharm，双击启动即可

后记：如果不记得安装路径了，可以尝试到下面地方找：

Exec=/home/fungtion(用户名)/Downloads/pycharm-2018.1.2/bin/pycharm.sh

Icon=/home/fungtion(用户名)/Downloads/pycharm-2018.1.2/bin/pycharm.png

### 安装 vscode

参考官网 [安装 vscode](https://code.visualstudio.com/docs/setup/linux)
与 [deb](https://code.visualstudio.com/Download) 包安装
sudo dpkg -i 对应的安装包名
如果出现依赖问题，执行：
sudo apt-get install -f
然后再次安装vscode
sudo dpkg -i 对应的安装包名
Ubuntu 应用商店直接安装 vscode
配置相关开发环境

### 安装 shutter、sublime3

安装 [sublime3](https://www.sublimetext.com/3),然后直接解压出来，双击里面的 sublime_text 就可以使用
Ubuntu 应用商店直接安装。

安装 [meetfranz 消息通讯](https://meetfranz.com/)
安装截图软件 shutter

sudo apt-get install shutter 

安装失败运行

sudo apt-get install -f 

然后重新安装

### 安装 go 环境

安装 go

参考 https://golang.org/doc/install

解决 go 被墙的问题

### 安装 shadowsocks 设置代理

https://wylu.github.io/posts/eed37a90/

### github 搭建自己的私有仓库

https://yeasy.gitbooks.io/docker_practice/repository/registry.html

https://www.cnblogs.com/dfengwei/p/7150455.html

### 安装 ntpdate 工具

同步网络时间

sudo apt-get install ntpdate

sudo cp /usr/share/zoneinfo/Asia/Chongqing /etc/localtime

1.安装ntpdate工具

sudo apt-get install ntpdate

2.设置系统时间与网络时间同步

ntpdate cn.pool.ntp.org

3.将系统时间写入硬件时间

hwclock -w