// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MovieDao is the data access object for table web_movie.
type MovieDao struct {
	table      string           // table is the underlying table name of the DAO.
	group      string           // group is the database configuration group name of current DAO.
	columns    MovieColumns     // columns contains all the column names of Table for convenient usage.
	primaryKey string           // 主键ID
	columnArr  []string         // 所有字段的数组
	columnArrG *garray.StrArray // 所有字段的数组（该格式更方便使用）
}

// MovieColumns defines and stores column names for table web_movie.
type MovieColumns struct {
	Id              string //
	MovieId         string // 电影id
	MovieName       string // 电影名称
	MovieRealName   string // 真实姓名
	MoviePicture    string // 主图
	PlotPhoto       string // 剧情图片
	TranslatedName  string // 译名
	SubIntros       string // 内容标签
	Url             string // 详情页地址
	Years           string // 年代
	Place           string // 产地
	CategoryName    string // 类别
	CategoryIds     string // 类别数组
	LanguageName    string // 语言
	Subtitle        string // 字幕
	ReleaseDate     string // 上映日期
	Imdb            string // IMDb评分
	Doubandb        string // 豆瓣评分
	Score           string // 评分
	Awards          string // 奖项
	Times           string //
	Time            string // 片长
	SortTime        string // 排序时间
	Director        string // 导演
	Scriptwriter    string // 编剧
	Actors          string // 演员
	Intros          string // 简介
	DownUrls        string // 下载地址
	VideoSize       string // 视频大小
	VideoResolution string // 视频分辨率
	IsUpdate        string // 更新：0否 1是
	IsStop          string // 停用：0否 1是
	UpdatedAt       string // 更新时间
	CreatedAt       string // 创建时间
}

// movieColumns holds the columns for table web_movie.
var movieColumns = MovieColumns{
	Id:              "id",
	MovieId:         "movie_id",
	MovieName:       "movie_name",
	MovieRealName:   "movie_real_name",
	MoviePicture:    "movie_picture",
	PlotPhoto:       "plot_photo",
	TranslatedName:  "translated_name",
	SubIntros:       "sub_intros",
	Url:             "url",
	Years:           "years",
	Place:           "place",
	CategoryName:    "category_name",
	CategoryIds:     "category_ids",
	LanguageName:    "language_name",
	Subtitle:        "subtitle",
	ReleaseDate:     "release_date",
	Imdb:            "imdb",
	Doubandb:        "doubandb",
	Score:           "score",
	Awards:          "awards",
	Times:           "times",
	Time:            "time",
	SortTime:        "sort_time",
	Director:        "director",
	Scriptwriter:    "scriptwriter",
	Actors:          "actors",
	Intros:          "intros",
	DownUrls:        "down_urls",
	VideoSize:       "video_size",
	VideoResolution: "video_resolution",
	IsUpdate:        "isUpdate",
	IsStop:          "isStop",
	UpdatedAt:       "updatedAt",
	CreatedAt:       "createdAt",
}

// NewMovieDao creates and returns a new DAO object for table data access.
func NewMovieDao() *MovieDao {
	return &MovieDao{
		group:   `default`,
		table:   `web_movie`,
		columns: movieColumns,
		primaryKey: func() string {
			return reflect.ValueOf(movieColumns).Field(0).String()
		}(),
		columnArr: func() []string {
			v := reflect.ValueOf(movieColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return column
		}(),
		columnArrG: func() *garray.StrArray {
			v := reflect.ValueOf(movieColumns)
			count := v.NumField()
			column := make([]string, count)
			for i := 0; i < count; i++ {
				column[i] = v.Field(i).String()
			}
			return garray.NewStrArrayFrom(column)
		}(),
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MovieDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MovieDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MovieDao) Columns() MovieColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MovieDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MovieDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MovieDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// 主键ID
func (dao *MovieDao) PrimaryKey() string {
	return dao.primaryKey
}

// 所有字段的数组
func (dao *MovieDao) ColumnArr() []string {
	return dao.columnArr
}

// 所有字段的数组（该格式更方便使用）
func (dao *MovieDao) ColumnArrG() *garray.StrArray {
	return dao.columnArrG
}
