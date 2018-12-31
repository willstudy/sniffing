package models

// 主页头部包含的Url信息
type BannerInfo struct {
	Img string `json:"img"`
	Target string `json:"target"`
}
type HotInfo struct {
	Avatar string `json:"avatar"`
	Title string `json:"tilte"`
	CreatedAt string `json:"createdAt"`
	Brief string `json:"brief"`
	Image string `json:"image"`
	Author string `json:"author"`
	PublishAt string `json:"publishedAt"`
	Source string `json:"source"`
	Type string `json:"type"`
	TargetUrl string `json:"target_url"`
	Comments uint32 `json:"comments"`
	Stars  uint32 `json:"stars"`
}
// 主页主体内容部分
type ContentInfo struct {
	Hot []HotInfo `json:"hot"`
}
// 主页所需的数据
type IndexData struct {
	Banner  []BannerInfo  `json:"banner"`
	Content ContentInfo `json:"content"`
}
// 主页返回信息
type IndexInfo struct {
	Code     int32         `json:"code"`
	Msg      string        `json:"msg"`
	Data     IndexData     `json:"data"`
}
