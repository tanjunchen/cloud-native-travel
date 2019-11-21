# Tools
The Tools for Check

## Procedure(步骤)
    1. git clone https://github.com/tanjunchen/Tools.git (tool)
    2. pip install -r requirements.txt (package) 安装 python  所需要的包
    3. 可视化 db.sqlite sqlitebrowser 安装 sqlitebrowser （Ubuntu 18.04 sudo apt-get install sqlitebrowser）
    4. 如果挂 VPN 访问大大降低 url_problems 表的大小
    5. 查看 url_problems 表中可疑的 url_problems 中存在url file_path(文件位置)
    6. 修改要扫描的文件路径,即 path 的值
    
## 下一步优化方案
    本代码===>先扫描整个仓库,即 path = "../minikube" 值。然后,然后调用 request get 请求,判断链接是否可用。
    
    下一步 扫描仓库、判断url是否可用(curl 或者或者其他更加高效的方式) 采用多线程并发
    
## 已经采取另一种方案,redis + 多进程  
    