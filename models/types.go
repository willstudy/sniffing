package models

type Thing struct {
	Id          uint64   `json:_id`
	CreateTime  string   `json:createdAt`
	Desc        string   `json:desc`
	Images      []string `json:images`
	PublishTime string   `json:publishedAt`
	Source      string   `json:source`
	Type        string   `json:type`
	Url         string   `json:url`
	Used        bool     `json:used`
	Who         string   `json:who`
}

type CatagoryThing struct {
	Android []Thing `json:"android,omitempty"`
	App     []Thing `json:"app,omitempty"`
	IOS     []Thing `json:"ios,omitempty"`
	Video   []Thing `json:"休闲视频,omitempty"`
	FE      []Thing `json:"前端,omitempty"`
	Extend  []Thing `json:"拓展资源,omitempty"`
	Welfare []Thing `json:"福利,omitempty"`
}

type TodayThing struct {
	Code     int32         `json:"code"`
	Msg      string        `json:"msg"`
	Catagory []string      `json:"category"`
	Results  CatagoryThing `json:"results"`
}
