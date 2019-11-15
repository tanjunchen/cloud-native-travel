# -*- coding: utf-8 -*-
import re, os, requests, sqlite3

'''
遍历指定目录 md 和 html 的文档列表
'''


def list_files(files_path):
    all_files = []
    if os.path.isdir(files_path):
        current_files = os.listdir(files_path)
        for file_name in current_files:
            full_file_name = os.path.join(files_path, file_name)
            file_suffix_name = os.path.splitext(full_file_name)[1]
            if file_suffix_name == '.md' or file_suffix_name == '.html':
                # print(os.path.dirname(full_file_name))
                all_files.append(full_file_name)
            if os.path.isdir(full_file_name):
                next_level_files = list_files(full_file_name)
                all_files.extend(next_level_files)
    else:
        all_files.append(files_path)
    return all_files


def find_all_url(filter_file, file_path, text):
    http_url = r'http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+'
    pattern = re.compile(http_url)  # 匹配模式
    urls = []
    for uu in re.findall(pattern, text):
        uu = uu.replace(")", "").replace("(", "").replace("[", "")\
            .replace("]", "").replace("**","").replace("'","")
        if uu[-1] in (",", ".", ";", "]", "!", ":", "*"):
            uu = uu[:-1]
        if not filter_file:
            urls.append([uu, file_path])
        elif not need_not_to_check(uu):
            urls.append([uu, file_path])
    return urls


def need_not_to_check(url, filter_file=True):
    if filter_file:
        if url.startswith("https://dl.k8s.io/"):
            return True
        if url.startswith("https://docs.k8s.io"):
            return True
        # if url.startswith("https://github.com/kubernetes/kubernetes/pull/"):
        #     return True
        if url.startswith("http://relnotes.k8s.io/"):
            return True
        # if url.startswith("https://github.com/"):
        #     return True
        if url.startswith("https://git.k8s.io/"):
            return True
    return False


def is_ok(url, file_path):
    try:
        res = requests.get(url, timeout=10)
        if res.ok:
            # print(url, "-->ok")
            return 1
        else:
            print(res.status_code, "--->", file_path, url, "--> fail")
            return 1
    except Exception as e:
        print(file_path, url, e)
        return 0


def init_db():
    # 连接数据库
    conn = sqlite3.connect("./db.sqlite")
    # 创建一个 cursor
    cursor = conn.cursor()
    cursor.execute("SELECT * FROM sqlite_master WHERE type='table' AND name='urls'")
    is_pr = cursor.fetchone()
    if is_pr is None:
        cursor.execute('''create table urls(
                                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                                    url varchar(1000),
                                    path varchar(1000),
                                    process Int(2)
                                    )''')
        print("create table urls")
    else:
        print("table urls exists")

    cursor.execute("SELECT * FROM sqlite_master WHERE type='table' AND name='url_problems'")
    is_pr = cursor.fetchone()
    if is_pr is None:
        cursor.execute('''create table url_problems(
                                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        url varchar(1000),
                                        path varchar(1000)
                                        )''')
        print("create table url_problems")
    else:
        print("table url_problems exists")
    conn.commit()
    cursor.close()
    conn.close()


def insert_all_url(url_data):
    # 连接数据库
    conn = sqlite3.connect("./db.sqlite")
    # 创建一个 cursor
    cursor = conn.cursor()
    for url in url_data:
        cursor.execute("select * from urls where url = '" + url[0] + "'")
        is_exist = cursor.fetchone()
        if is_exist is None:
            # process 0 未处理 1 已处理
            print("insert into urls (url,path,process) values ('" + url[0] + "', '" + url[1] + "',0)")
            cursor.execute("insert into urls (url,path,process) values ('" + url[0] + "', '" + url[1] + "',0)")
        else:
            print(url[0], "db existed")
    conn.commit()
    cursor.close()
    conn.close()


def get_data(filter_file=True):
    all_file_path = list_files(path)
    for file_path in all_file_path:
        print(file_path)
        with open(file_path, encoding="utf-8") as f:
            url_data = find_all_url(filter_file, file_path, f.read())
            insert_all_url(url_data)


def analysis():
    # 连接数据库
    conn = sqlite3.connect("./db.sqlite")
    # 创建一个 cursor
    cursor = conn.cursor()
    cursor.execute("select url,path from urls where  process=0")
    all_data = cursor.fetchall()
    for i in all_data:
        result = is_ok(i[0], i[1])
        if result == 1:
            cursor.execute("update  urls set process=1 where url = '" + i[0] + "'")
        else:
            cursor.execute("select * from urls where url = '" + i[0] + "'")
            url_exist = cursor.fetchone()
            if url_exist is not None:
                sql = " insert into url_problems (url,path) values (?,?) "
                cursor.execute(sql, (i[0], i[1]))
    conn.commit()
    cursor.close()
    conn.close()


def job(filter_file):
    init_db()
    # get_data(filter_file)
    analysis()


if __name__ == '__main__':
    """
    跑数据之前请删除无关的文件夹 如 kubernetes 的 vendor  third_party 减少扫描次数
    """
    path = "/home/k8s-master/goproject/bocloud/website/content/en/"
    job(filter_file=False)
