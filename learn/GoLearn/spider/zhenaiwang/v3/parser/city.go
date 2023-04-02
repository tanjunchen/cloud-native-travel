package parser

import (
	"go-learn/spider/zhenaiwang/v3/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

// 城市页面用户解析器
func ParseCity(bytes []byte) engine.ParseResult {
	subMatch := cityRe.FindAllSubmatch(bytes, -1)
	genderMatch := sexRe.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}

	for k, item := range subMatch {
		name := string(item[2])
		gender := string(genderMatch[k][1])

		result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(item[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name, gender)
			},
		})
	}
	return result
}
