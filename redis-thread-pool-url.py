# -*- coding: utf-8 -*-

import json
import os
import re
import chardet
import requests
from fake_useragent import UserAgent
from redis import Redis
from concurrent.futures import ThreadPoolExecutor


redis = Redis(host='localhost', port=6378, db=15)
ua = UserAgent()


def contents(filepath):
    content_list = []
    # 遍历文件夹
    for root, dirs, files in os.walk(filepath):
        for file in files:
            file_name = os.path.join(root, file)
            content_list.append(file_name)
    return content_list


def find_all_url(file_path):
    f = open(file_path, 'rb')
    text = f.read()
    result = chardet.detect(text)
    try:
        text = text.decode(result.get('encoding'), errors='ignore')

    except TypeError as e:
        text = text.decode('utf-8', errors='ignore')

    http_url = r'http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+'
    pattern = re.compile(http_url)  # 匹配模式
    for url in re.findall(pattern, text):
        url = re.sub(r'\)|\(|\[|\]|\**|\'|,', '', url)
        url = url[:-1] if url[-1] in (",", ".", ";", "]", "!", ":", "*",) else url
        if url.endswith('</PackageIconUrl>') or url.endswith('</PackageProjectUrl>'):
            continue

        if url:
            print(url)
            redis.hset(':url', url, file_path)

    f.close()


def spider():
    while True:
        try:
            url, path = json.loads(redis.brpop(':url_path', timeout=120)[1])
            try:
                res = requests.get(url, headers={'User-Agent': ua.random, 'Connection': 'close'}, timeout=15)
                if res.status_code == 200:
                    redis.lpush(':legal_url', json.dumps([url, path]))
                    print(url)
                elif res.status_code == 404:
                    print("404-->", url, path)
            except TimeoutError as e:
                print(e, url)

        except TypeError as e:
            print('线程退出')
            return
        except Exception as e:
            print(f'error=====================>', e)


def main(filepath):
    content_list = contents(filepath)
    for file_path in content_list:
        if file_path.endswith('.pdf') or file_path.endswith('png') or file_path.endswith(
                'jar') or file_path.endswith('snk'):
            continue
        elif filepath + '\.git' in file_path:
            continue
        find_all_url(file_path)

    map_list = redis.hgetall(':url')
    map_list = {key.decode('utf-8'): value.decode('utf-8') for key, value in map_list.items()}
    for key, value in map_list.items():
        redis.lpush(":url_path", json.dumps([key, value]))

    executor = ThreadPoolExecutor(max_workers=16)
    print('=============================================================================================>开始测试')
    for _ in range(16):
        executor.submit(spider)
    executor.shutdown(wait=True)


if __name__ == '__main__':
    path = 'd:\\opensource\\grpc'
    main(path)
