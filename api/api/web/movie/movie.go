package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------列表 开始--------*/
type MovieListReq struct {
	g.Meta `path:"/movie/list" method:"post" tags:"web前端/电影" sm:"列表"`
	Filter MovieListFilter `json:"filter" dc:"查询条件"`
	Field  []string        `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
	Sort   string          `json:"sort" default:"id DESC" dc:"排序"`
	Page   int             `json:"page" v:"integer|min:1" default:"1" dc:"页码"`
	Limit  int             `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"`
}

type MovieListFilter struct {
	Id              *uint       `json:"id,omitempty" v:"integer|min:1" dc:"ID"`
	IdArr           []uint      `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"`
	ExcId           *uint       `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"`
	ExcIdArr        []uint      `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"`
	Label           string      `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`
	MovieId         string      `json:"movie_id,omitempty" v:"length:1,255" dc:"电影id"`
	MovieName       string      `json:"movie_name,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"电影名称"`
	MovieRealName   string      `json:"movie_real_name,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"真实姓名"`
	MoviePicture    string      `json:"movie_picture,omitempty" v:"length:1,500" dc:"主图"`
	PlotPhoto       string      `json:"plot_photo,omitempty" v:"length:1,500" dc:"剧情图片"`
	TranslatedName  string      `json:"translated_name,omitempty" v:"length:1,255|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"译名"`
	Url             string      `json:"url,omitempty" v:"url|length:1,255" dc:"详情页地址"`
	Years           string      `json:"years,omitempty" v:"length:1,25" dc:"年代"`
	Place           string      `json:"place,omitempty" v:"length:1,255" dc:"产地"`
	CategoryName    string      `json:"category_name,omitempty" v:"length:1,50|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"类别"`
	CategoryIds     string      `json:"category_ids,omitempty" v:"length:1,50" dc:"类别数组"`
	LanguageName    string      `json:"language_name,omitempty" v:"length:1,100|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"语言"`
	Subtitle        string      `json:"subtitle,omitempty" v:"length:1,25" dc:"字幕"`
	ReleaseDate     string      `json:"release_date,omitempty" v:"length:1,100" dc:"上映日期"`
	Imdb            string      `json:"imdb,omitempty" v:"length:1,255" dc:"IMDb评分"`
	Doubandb        string      `json:"doubandb,omitempty" v:"length:1,255" dc:"豆瓣评分"`
	Score           string      `json:"score,omitempty" v:"length:1,5" dc:"评分"`
	Awards          string      `json:"awards,omitempty" v:"length:1,500" dc:"奖项"`
	Times           string      `json:"times,omitempty" v:"length:1,500" dc:""`
	Time            string      `json:"time,omitempty" v:"length:1,25" dc:"片长"`
	Director        string      `json:"director,omitempty" v:"length:1,50" dc:"导演"`
	Scriptwriter    string      `json:"scriptwriter,omitempty" v:"length:1,100" dc:"编剧"`
	VideoSize       string      `json:"video_size,omitempty" v:"length:1,25" dc:"视频大小"`
	VideoResolution string      `json:"video_resolution,omitempty" v:"length:1,25" dc:"视频分辨率"`
	IsUpdate        *uint       `json:"isUpdate,omitempty" v:"integer|in:0,1" dc:"更新：0否 1是"`
	IsStop          *uint       `json:"isStop,omitempty" v:"integer|in:0,1" dc:"停用：0否 1是"`
	TimeRangeStart  *gtime.Time `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd    *gtime.Time `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
}

type MovieListRes struct {
	Count int         `json:"count" dc:"总数"`
	List  []MovieItem `json:"list" dc:"列表"`
}

type MovieItem struct {
	Id *uint `json:"id,omitempty" dc:"ID"`
	//Label           *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	//MovieId        *string `json:"movie_id,omitempty" dc:"电影id"`
	MovieName      *string `json:"movie_name,omitempty" dc:"电影名称"`
	MovieRealName  *string `json:"movie_real_name,omitempty" dc:"真实姓名"`
	MoviePicture   *string `json:"movie_picture,omitempty" dc:"主图"`
	PlotPhoto      *string `json:"plot_photo,omitempty" dc:"剧情图片"`
	TranslatedName *string `json:"translated_name,omitempty" dc:"译名"`
	SubIntros      *string `json:"sub_intros,omitempty" dc:"内容标签"`
	//Url             *string     `json:"url,omitempty" dc:"详情页地址"`
	Years        *string `json:"years,omitempty" dc:"年代"`
	Place        *string `json:"place,omitempty" dc:"产地"`
	CategoryName *string `json:"category_name,omitempty" dc:"类别"`
	CategoryIds  *string `json:"category_ids,omitempty" dc:"类别数组"`
	LanguageName *string `json:"language_name,omitempty" dc:"语言"`
	//Subtitle        *string     `json:"subtitle,omitempty" dc:"字幕"`
	ReleaseDate *string `json:"release_date,omitempty" dc:"上映日期"`
	Imdb        *string `json:"imdb,omitempty" dc:"IMDb评分"`
	Doubandb    *string `json:"doubandb,omitempty" dc:"豆瓣评分"`
	Score       *string `json:"score,omitempty" dc:"评分"`
	Awards      *string `json:"awards,omitempty" dc:"奖项"`
	//Times           *string     `json:"times,omitempty" dc:""`
	Time         *string  `json:"time,omitempty" dc:"片长"`
	Director     *string  `json:"director,omitempty" dc:"导演"`
	Scriptwriter *string  `json:"scriptwriter,omitempty" dc:"编剧"`
	Actors       *string  `json:"actors,omitempty" dc:"演员"`
	ActorsArray  []string `json:"actors_array,omitempty" dc:"演员分解"`
	Intros       *string  `json:"intros,omitempty" dc:"简介"`
	//DownUrls        *string     `json:"down_urls,omitempty" dc:"下载地址"`
	//VideoSize       *string     `json:"video_size,omitempty" dc:"视频大小"`
	VideoResolution *string `json:"video_resolution,omitempty" dc:"视频分辨率"`
	//IsUpdate        *uint       `json:"isUpdate,omitempty" dc:"更新：0否 1是"`
	//IsStop          *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	//UpdatedAt       *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	//CreatedAt       *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type MovieInfoReq struct {
	g.Meta `path:"/movie/info" method:"post" tags:"web前端/电影" sm:"详情"`
	Id     uint     `json:"id" v:"required|integer|min:1" dc:"ID"`
	Field  []string `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"`
}

type MovieInfoRes struct {
	Info MovieInfo `json:"info" dc:"详情"`
}

type MovieInfo struct {
	Id *uint `json:"id,omitempty" dc:"ID"`
	//Label           *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	//MovieId         *string     `json:"movie_id,omitempty" dc:"电影id"`
	MovieName      *string `json:"movie_name,omitempty" dc:"电影名称"`
	MovieRealName  *string `json:"movie_real_name,omitempty" dc:"真实姓名"`
	MoviePicture   *string `json:"movie_picture,omitempty" dc:"主图"`
	PlotPhoto      *string `json:"plot_photo,omitempty" dc:"剧情图片"`
	TranslatedName *string `json:"translated_name,omitempty" dc:"译名"`
	SubIntros      *string `json:"sub_intros,omitempty" dc:"内容标签"`
	//Url             *string     `json:"url,omitempty" dc:"详情页地址"`
	Years        *string `json:"years,omitempty" dc:"年代"`
	Place        *string `json:"place,omitempty" dc:"产地"`
	CategoryName *string `json:"category_name,omitempty" dc:"类别"`
	CategoryIds  *string `json:"category_ids,omitempty" dc:"类别数组"`
	LanguageName *string `json:"language_name,omitempty" dc:"语言"`
	//Subtitle        *string     `json:"subtitle,omitempty" dc:"字幕"`
	ReleaseDate  *string  `json:"release_date,omitempty" dc:"上映日期"`
	Imdb         *string  `json:"imdb,omitempty" dc:"IMDb评分"`
	Doubandb     *string  `json:"doubandb,omitempty" dc:"豆瓣评分"`
	Score        *string  `json:"score,omitempty" dc:"评分"`
	Awards       *string  `json:"awards,omitempty" dc:"奖项"`
	Times        *string  `json:"times,omitempty" dc:""`
	Time         *string  `json:"time,omitempty" dc:"片长"`
	Director     *string  `json:"director,omitempty" dc:"导演"`
	Scriptwriter *string  `json:"scriptwriter,omitempty" dc:"编剧"`
	Actors       *string  `json:"actors,omitempty" dc:"演员"`
	ActorsArray  []string `json:"actors_array,omitempty" dc:"演员分解"`
	Intros       *string  `json:"intros,omitempty" dc:"简介"`
	//DownUrls        *string     `json:"down_urls,omitempty" dc:"下载地址"`
	//VideoSize       *string     `json:"video_size,omitempty" dc:"视频大小"`
	VideoResolution *string `json:"video_resolution,omitempty" dc:"视频分辨率"`
	//IsUpdate        *uint       `json:"isUpdate,omitempty" dc:"更新：0否 1是"`
	//IsStop          *uint       `json:"isStop,omitempty" dc:"停用：0否 1是"`
	//UpdatedAt       *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	//CreatedAt       *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------详情 结束--------*/
