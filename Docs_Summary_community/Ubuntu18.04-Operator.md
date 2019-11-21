# Ubuntu 18.04 常见操作
1. 设置静态 IP
    参考链接[设置静态IP](https://www.cnblogs.com/jianxuanbing/p/10042892.html)、[设置静态IP](https://blog.csdn.net/qq_19004627/article/details/90409420)
2. SSH 与 SCP
    ssh ip   # ssh命令接IP地址，之后会提示你输入用户名称和密码
    ssh username@ip  # username为登录时的用户名，之后只需要输入密码就能进行登录
    ssh -p port username@ip  # port是服务器端设置的端口号，有时候在服务器端，设置了端口号就需要输入对应的端口号

    scp
    从本地上传文件到服务器
    scp local_file remote_username@remote_ip:remote_folder #上传到指定的文件夹内
    scp local_file remote_username@remote_ip:remote_file #指定了上传后的文件名
    从本地上传目录到服务器
    scp -r local_folder remote_username@remote_ip:remote_folder ＃-r：递归复制整个目录

    从服务器下载文件/目录到本地
    scp remote_username@remote_ip:remote_file local_folder #下载文件
    scp -r remote_username@remote_ip:remote_folder local_folder #下载文件夹
3. 