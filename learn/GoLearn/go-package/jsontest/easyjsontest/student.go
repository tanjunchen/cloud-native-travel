package main

import "time"

// https://github.com/pquerna/ffjson
// github.com/mailru/easyjson/ 预编译
// https://github.com/json-iterator/go
// https://github.com/bitly/go-simplejson
// https://github.com/antonholmquist/jason
// https://github.com/buger/jsonparser

// 大部分情况下大家直接使用 encoding/json就行了。
// 如果追求极致的性能，考虑 easyjson。
// 遇到解析 ES 搜索返回的复杂的 JSON 或者仅需要解析个别字段
// go-simplejson 或者 jsonparser 就 OK。

func main()  {
	//easyjson:json
	type School struct {
		Name string     `json:"name"`
		Addr string     `json:"addr"`
	}

	//easyjson:json
	type Student struct {
		Id       int       `json:"id"`
		Name     string    `json:"s_name"`
		School   School    `json:"s_chool"`
		Birthday time.Time `json:"birthday"`
	}



}
