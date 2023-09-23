// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Movie is the golang structure for table movie.
type Movie struct {
	Id              int         `json:"id"              ` //
	MovieId         string      `json:"movieId"         ` // 电影id
	MovieName       string      `json:"movieName"       ` // 电影名称
	MovieRealName   string      `json:"movieRealName"   ` // 真实姓名
	MoviePicture    string      `json:"moviePicture"    ` // 主图
	PlotPhoto       string      `json:"plotPhoto"       ` // 剧情图片
	TranslatedName  string      `json:"translatedName"  ` // 译名
	SubIntros       string      `json:"subIntros"       ` // 内容标签
	Url             string      `json:"url"             ` // 详情页地址
	Years           string      `json:"years"           ` // 年代
	Place           string      `json:"place"           ` // 产地
	CategoryName    string      `json:"categoryName"    ` // 类别
	CategoryIds     string      `json:"categoryIds"     ` // 类别数组
	LanguageName    string      `json:"languageName"    ` // 语言
	Subtitle        string      `json:"subtitle"        ` // 字幕
	ReleaseDate     string      `json:"releaseDate"     ` // 上映日期
	Imdb            string      `json:"imdb"            ` // IMDb评分
	Doubandb        string      `json:"doubandb"        ` // 豆瓣评分
	Score           string      `json:"score"           ` // 评分
	Awards          string      `json:"awards"          ` // 奖项
	Times           string      `json:"times"           ` //
	Time            string      `json:"time"            ` // 片长
	SortTime        string      `json:"sortTime"        ` // 排序时间
	Director        string      `json:"director"        ` // 导演
	Scriptwriter    string      `json:"scriptwriter"    ` // 编剧
	Actors          string      `json:"actors"          ` // 演员
	Intros          string      `json:"intros"          ` // 简介
	DownUrls        string      `json:"downUrls"        ` // 下载地址
	VideoSize       string      `json:"videoSize"       ` // 视频大小
	VideoResolution string      `json:"videoResolution" ` // 视频分辨率
	IsUpdate        uint        `json:"isUpdate"        ` // 更新：0否 1是
	IsStop          uint        `json:"isStop"          ` // 停用：0否 1是
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
}
