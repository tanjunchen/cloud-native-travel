# -*- coding: utf-8 -*-
import re, os
import requests

'''
遍历指定目录 md 和 html 的文档列表
'''


def list_files(files_path):
    current_files = os.listdir(files_path)
    all_files = []
    for file_name in current_files:
        full_file_name = os.path.join(files_path, file_name)
        file_suffix_name = os.path.splitext(full_file_name)[1]
        if file_suffix_name == '.md' or file_suffix_name == '.html':
            # print(os.path.dirname(full_file_name))
            all_files.append(full_file_name)
        if os.path.isdir(full_file_name):
            next_level_files = list_files(full_file_name)
            all_files.extend(next_level_files)
    return all_files


def find_all_url(text):
    http_url = r'http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+'
    pattern = re.compile(http_url)  # 匹配模式
    urls = []
    for uu in re.findall(pattern, text):
        uu = uu.replace(")", "").replace("(", "").replace("[", "").replace("]", "")
        if uu[-1] in (",", ".", ";", "]", "!", ":", "**"):
            uu = uu[:-1]
        if not need_not_to_check(uu):
            urls.append(uu)
    return urls


def need_not_to_check(url):
    if url.startswith("https://dl.k8s.io/"):
        return True
    if url.startswith("https://docs.k8s.io"):
        return True
    if url.startswith("https://github.com/kubernetes/kubernetes/pull/"):
        return True
    if url.startswith("http://relnotes.k8s.io/"):
        return True
    if url.startswith("https://github.com/"):
        return True
    if url.startswith("https://git.k8s.io/"):
        return True
    return False


def is_ok(url, file_path):
    try:
        res = requests.get(url, timeout=10)
        if res.ok:
            # print(url, "-->ok")
            pass
        else:
            print(res.status_code, file_path, url, "-->fail")
    except Exception as e:
        print(file_path, url, e)


def job():
    all_file_path = list_files("D:\\opensource\\istio.io\\content\\zh")
    for file_path in all_file_path:
        with open(file_path, encoding="utf-8") as f:
            all_url = find_all_url(f.read())
            for i in all_url:
                # is_ok(i, file_path)
                print(i)


def test():
    with open("content.md", encoding="utf-8") as f:
        content = f.read()
    all_url = find_all_url(content)
    for i in all_url:
        is_ok(i, "content.md")


if __name__ == '__main__':
    # job()
    test()
