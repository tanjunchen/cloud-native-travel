import requests
import json
from string import Template

codeOffer = '''
package main

/***
"题目：**$name**

[$name]($url)

题目描述：$describe
***/

/**
解法一
说明：
**/


/**
解法二
说明：
**/


/**
解法三
说明：
**/


func main() {
    
}
'''


def writeTemplate(allStr):
    for k, v in allStr.items():
        with open(k + ".go", encoding="utf-8", mode='w') as f:
            f.write(v)


def generate(name, url, describe):
    t = Template(codeOffer)
    res = t.substitute(name=name, url=url, describe=describe)
    return res


def get_problem_by_slug(slug):
    session = requests.Session()
    user_agent = r'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36'
    url = "https://leetcode-cn.com/graphql"
    params = {'operationName': "getQuestionDetail",
              'variables': {'titleSlug': slug},
              'query': '''query getQuestionDetail($titleSlug: String!) {
            question(titleSlug: $titleSlug) {
                questionId
                questionFrontendId
                questionTitle
                questionTitleSlug
                content
                difficulty
                stats
                similarQuestions
                translatedContent
                categoryTitle
                topicTags {
                        name
                        slug
                }
            }
        }'''
    }

    json_data = json.dumps(params).encode('utf8')

    headers = {'User-Agent': user_agent, 'Connection':
        'keep-alive', 'Content-Type': 'application/json',
               'Referer': 'https://leetcode-cn.com/problems/' + slug}
    resp = session.post(url, data=json_data, headers=headers, timeout=10)
    content = resp.json()

    # 题目详细信息
    question = content['data']['question']['translatedContent']
    return question


def getAllResult():
    codeUrl = "https://leetcode-cn.com/problems/"
    res = requests.get("https://leetcode-cn.com/api/problems/lcof/")
    jsonData = json.loads(res.content)["stat_status_pairs"]
    allStr = {}
    for i in jsonData:
        value = i['stat']
        key = value["frontend_question_id"].replace("剑指 Offer ", "").replace(" ", "")
        value = generate(name=value["question__title"].replace(" LCOF", "").replace(" ", ""),
                         url=codeUrl + value["question__title_slug"], describe="")
        allStr.__setitem__(key, value)

    writeTemplate(allStr)
    # [不用加减乘除做加法](code/65/65.go)
    for k, v in sorted(read_me.items()):
        with open("readme.md", encoding="utf-8", mode='a+') as f:
            f.write("[" + v + "]" + "(offer/" + k + ".go)\n\n")


def build():
    getAllResult()


if __name__ == '__main__':
    build()
