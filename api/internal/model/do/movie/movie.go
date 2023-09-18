// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Movie is the golang structure of table web_movie for DAO operations like Where/Data.
type Movie struct {
	g.Meta          `orm:"table:web_movie, do:true"`
	Id              interface{} //
	MovieId         interface{} // 电影id
	MovieName       interface{} // 电影名称
	MovieRealName   interface{} // 真实姓名
	MoviePicture    interface{} // 主图
	PlotPhoto       interface{} // 剧情图片
	TranslatedName  interface{} // 译名
	SubIntros       interface{} // 内容标签
	Url             interface{} // 详情页地址
	Years           interface{} // 年代
	Place           interface{} // 产地
	CategoryName    interface{} // 类别
	CategoryIds     interface{} // 类别数组
	LanguageName    interface{} // 语言
	Subtitle        interface{} // 字幕
	ReleaseDate     interface{} // 上映日期
	Imdb            interface{} // IMDb评分
	Doubandb        interface{} // 豆瓣评分
	Score           interface{} // 评分
	Awards          interface{} // 奖项
	Times           interface{} //
	Time            interface{} // 片长
	Director        interface{} // 导演
	Scriptwriter    interface{} // 编剧
	Actors          interface{} // 演员
	Intros          interface{} // 简介
	DownUrls        interface{} // 下载地址
	VideoSize       interface{} // 视频大小
	VideoResolution interface{} // 视频分辨率
	IsUpdate        interface{} // 更新：0否 1是
	IsStop          interface{} // 停用：0否 1是
	UpdatedAt       *gtime.Time // 更新时间
	CreatedAt       *gtime.Time // 创建时间
}
