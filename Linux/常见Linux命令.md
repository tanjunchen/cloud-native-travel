# top

Tasks: 324 total,   1 running, 255 sleeping,   0 stopped,   1 zombie

进程：当前共有 324 个进程，1 个运行中，255 个处于睡眠态，0 个停止态，1 个僵尸态

cpu状态：

us --用户空间占用cpu百分比

sy --内核空间占用cpu百分比

ni --改变过优先级的进程占用cpu百分比

id --空闲cpu百分比

wa --I/O输入/输出等待占用cpu百分比

hi --硬中断占用cpu百分比

si --软中断占用cpu百分比

st --虚拟cpu等待实际cpu的时间的百分比

KiB Mem :  4002452 total,   706068 free,  1117756 used,  2178628 buff/cache

物理内存总量，使用中的内存总量，空闲中的内存总量，内核缓存区中的内存量

KiB Swap:  2097148 total,  1383984 free,   713164 used.  2583004 avail Mem

交换区总量，使用中的交换区总量，空闲中的交换区总量，缓冲中的交换区总量

PID  USER      PR  NI   VIRT   RES   SHR  S  %CPU  %MEM    TIME+    COMMAND

PID --进程ID

USER --进程所有者用户名

PR --进程优先调度值

NI --进程nice值（优先级），值越小优先级越高

VIRT --进程使用的虚拟内存总量，单位kb

RES --驻留内存大小，单位kb

SHR --进程使用的共享内存大小，单位kb

S --进程状态，D不可中断的睡眠状态 R运行态 S睡眠态 T跟踪/停止态 Z僵尸态

%CPU --上次更新到现在的CPU时间占用百分比

%MEM --进程使用的物理内存百分比

TIME+ --进程使用的CPU时间总计，单位1/100秒

COMMAND --命令名/命令行

```
top - 23:59:25 up 4 days, 19:22,  2 users,  load average: 0.04, 0.03, 0.00
Tasks: 324 total,   1 running, 255 sleeping,   0 stopped,   1 zombie
%Cpu(s):  0.4 us,  0.9 sy,  0.0 ni, 98.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :  4002452 total,   706068 free,  1117756 used,  2178628 buff/cache
KiB Swap:  2097148 total,  1383984 free,   713164 used.  2583004 avail Mem 

   PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND                                                                                            
 49747 k8s-dev+  20   0   51448   4008   3236 R  1.3  0.1   0:00.04 top                                                                                                
 49727 k8s-dev+  20   0  110084   5528   4508 S  0.4  0.1   0:02.35 sshd                                                                                               
     1 root      20   0  160276   5632   3580 S  0.0  0.1   0:28.72 systemd                                                                                            
     2 root      20   0       0      0      0 S  0.0  0.0   0:00.05 kthreadd                                                                                           
     3 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 rcu_gp                                                                                             
     4 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 rcu_par_gp                                                                                         
     6 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 kworker/0:0H-kb                                                                                    
     9 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 mm_percpu_wq                                                                                       
    10 root      20   0       0      0      0 S  0.0  0.0   0:19.24 ksoftirqd/0                                                                                        
    11 root      20   0       0      0      0 I  0.0  0.0   0:38.26 rcu_sched                                                                                          
    12 root      rt   0       0      0      0 S  0.0  0.0   0:01.69 migration/0                                                                                        
    13 root     -51   0       0      0      0 S  0.0  0.0   0:00.00 idle_inject/0                                                                                      
    14 root      20   0       0      0      0 S  0.0  0.0   0:00.00 cpuhp/0                                                                                            
    15 root      20   0       0      0      0 S  0.0  0.0   0:00.00 kdevtmpfs                                                                                          
    16 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 netns                                                                                              
    17 root      20   0       0      0      0 S  0.0  0.0   0:00.00 rcu_tasks_kthre                                                                                    
    18 root      20   0       0      0      0 S  0.0  0.0   0:00.00 kauditd                                                                                            
    19 root      20   0       0      0      0 S  0.0  0.0   0:00.56 khungtaskd                                                                                         
    20 root      20   0       0      0      0 S  0.0  0.0   0:00.00 oom_reaper                                                                                         
    21 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 writeback                                                                                          
    22 root      20   0       0      0      0 S  0.0  0.0   0:14.72 kcompactd0                                                                                         
    23 root      25   5       0      0      0 S  0.0  0.0   0:00.00 ksmd                                                                                               
    24 root      39  19       0      0      0 S  0.0  0.0   0:00.51 khugepaged                                                                                         
   116 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 kintegrityd                                                                                        
   117 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 kblockd                                                                                            
   118 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 blkcg_punt_bio                                                                                     
   119 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 tpm_dev_wq                                                                                         
   120 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 ata_sff                                                                                            
   121 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 md                                                                                                 
   122 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 edac-poller                                                                                        
   123 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 devfreq_wq                                                                                         
   124 root      rt   0       0      0      0 S  0.0  0.0   0:00.00 watchdogd                                                                                          
   127 root      20   0       0      0      0 S  0.0  0.0   2:33.40 kswapd0                                                                                            
   129 root      20   0       0      0      0 S  0.0  0.0   0:00.00 ecryptfs-kthrea                                                                                    
   132 root       0 -20       0      0      0 I  0.0  0.0   0:00.00 kthrotld
```

# ps

# pstree

# free

# cat

# tail

# head

# more

# less

# grep

# awk

行级文件分析工具

# sort

# uniq

# parallel

# scp

# du

查看文件或目录所占用的磁盘空间的大小

# df

查看磁盘使用情况

# iostat

# iotop

# find

文件搜索

# locate

查找文件

# tree

# ping

# nc netcat

# route

# netstat

# iptables

# tcpdump



# traceroute

# iftop

# lsof

# dig

# curl

# wget



# yum|apt|brew install

# man

# tar

解压缩文件

-c	产生 tar 打包文件

-x	产生的解压缩文件

-v	显示详细信息

-f	指定压缩后的文件名

-z	打包同时压缩

tar -zcvf *.tar.gz d/f

# vim

# dd

复制文件并对原文件的内容进行转换和格式化处理

# fdisk

磁盘分区工具

# link ln

创建连接



Linux 命令

# 文件管理

cat

chattr

chgrp

chmod

chown

cksum

cmp

diff

diffstat

file

find

git

gitview

indent

cut

ln

less

locate

lsattr

mattrib

mc

mdel

mdir

mktemp

more

mmove

mread

mren

mtools

mtoolstest

mv

od

paste

patch

rcp

rm

slocate

split

tee

tmpwatch

touch

umask

which

cp

whereis

mcopy

mshowfat

rhmask

scp

awk

read

updatedb

# 文档编辑

col

colrm

comm

csplit

ed

egrep
	
ex

fgrep

fmt	

fold	

grep
	
ispell

jed	

joe	

join	

look

mtype
	
pico
	
rgrep	

sed

sort
	
spell
	
tr	

expr

uniq
	
wc
	
let	 

# 文件传输

lprm	

lpr
	
lpq
	
lpd

bye
	
ftp
	
uuto
	
uupick

uucp
	
uucico
	
tftp
	
ncftp

ftpshut

ftpwho

ftpcount
 
# 磁盘管理

cd
	
df
	
dirs
	
du

edquota

eject
	
mcd	

mdeltree

mdu	

mkdir	

mlabel	

mmd

mrd	

mzip
	
pwd	

quota

mount
	
mmount	

rmdir
	
rmt

stat
	
tree	

umount	

ls

quotacheck	

quotaoff
	
lndir
	
repquota

quotaon	 
	 	 
# 磁盘维护

badblocks	

cfdisk	

dd	

e2fsck

ext2ed	

fsck
	
fsck.minix	

fsconf

fdformat	

hdparm	

mformat	

mkbootdisk

mkdosfs	

mke2fs	

mkfs.ext2	

mkfs.msdos

mkinitrd	

mkisofs	

mkswap
	
mpartition

swapon	

symlinks	

sync	

mbadblocks

mkfs.minix	

fsck.ext2
	
fdisk	

losetup

mkfs	

sfdisk	

swapoff	
 
# 网络通讯

apachectl	

arpwatch
	
dip	

getty

mingetty	

uux	

telnet	

uulog

uustat	

ppp-off	

netconfig
	
nc

httpd	

ifconfig	

minicom	

mesg

dnsconf	

wall	

netstat	

ping

pppstats	

samba
	
setserial	

talk

traceroute	

tty	

newaliases
	
uuname

netconf	

write	

statserial	

efax

pppsetup	

tcpdump	

ytalk	

cu

smbd	

testparm
	
smbclient	

shapecfg

# 系统管理

adduser

chfn

useradd	

date

exit	

finger
	
fwhios	

sleep

suspend	

groupdel	

groupmod
	
halt

kill	

last	

lastb	

login

logname	

logout	

ps	

nice

procinfo
	
top	

pstree
	
reboot

rlogin	

rsh	

sliplogin
	
screen

shutdown
	
rwho	

sudo
	
gitps

swatch	

tload	

logrotate	

uname

chsh	

userconf
	
userdel	

usermod

vlock	

who	

whoami	

whois

newgrp	

renice	

su	

skill

w	

id	

groupadd
	
free

# 系统设置

reset
	
clear

alias

dircolors

aumix

bind

chroot

clock

crontab
	
declare

depmod

dmesg

enable

eval

export

pwunconv

grpconv

rpm

insmod

kbdconfig

lilo

liloconfig

lsmod

minfo

set

modprobe

ntsysv

mouseconfig

passwd

pwconv

rdate

resize

rmmod

grpunconv

modinfo

time

setup

sndconfig

setenv

setconsole

timeconfig

ulimit

unset

chkconfig

apmd

hwclock

mkkickstart

fbset

unalias	

SVGATextMode

gpasswd	 

# 备份压缩

ar

bunzip

bzip2	

bzip2recover

gunzip	

unarj	

compress	

cpio

dump	

uuencode

gzexe	

gzip

lha	

restore

tar

uudecode

unzip

zip	

zipinfo	
 
# 设备管理

setleds	

loadkeys	

rdev	

dumpkeys

MAKEDEV	 

https://www.runoob.com/linux/linux-command-manual.html