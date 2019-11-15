# -*- coding: utf-8 -*-


def job():
    with open("a.txt", encoding="utf-8") as f:
        content = f.read()
    cc = [c for c in content.split("\n")]
    for i in cc:
        if i.startswith("404"):
            print(i.replace(" --> fail", "").replace("404 ---> /home/k8s-master/goproject/bocloud/", ""))


if __name__ == '__main__':
    job()
