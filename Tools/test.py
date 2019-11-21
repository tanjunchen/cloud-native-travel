# -*- coding: utf-8 -*-
def job():
    with open("kubernetes_404.md", encoding="utf-8") as f:
        content = f.read()
    ii = 1
    for i in content.splitlines():
        print("|" + str(ii) + "|" + i + " |")
        ii = ii + 1


if __name__ == '__main__':
    job()
