// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"DouSheng/dal/model"
)

func newVideoInfo(db *gorm.DB, opts ...gen.DOOption) videoInfo {
	_videoInfo := videoInfo{}

	_videoInfo.videoInfoDo.UseDB(db, opts...)
	_videoInfo.videoInfoDo.UseModel(&model.VideoInfo{})

	tableName := _videoInfo.videoInfoDo.TableName()
	_videoInfo.ALL = field.NewAsterisk(tableName)
	_videoInfo.VideoID = field.NewInt64(tableName, "video_id")
	_videoInfo.AuthID = field.NewInt64(tableName, "auth_id")
	_videoInfo.Titel = field.NewString(tableName, "titel")
	_videoInfo.CommentCount = field.NewInt64(tableName, "comment_count")
	_videoInfo.FavoriteCount = field.NewInt64(tableName, "favorite_count")
	_videoInfo.CoverURL = field.NewString(tableName, "cover_url")
	_videoInfo.PlayURL = field.NewString(tableName, "play_url")

	_videoInfo.fillFieldMap()

	return _videoInfo
}

type videoInfo struct {
	videoInfoDo videoInfoDo

	ALL           field.Asterisk
	VideoID       field.Int64
	AuthID        field.Int64
	Titel         field.String
	CommentCount  field.Int64
	FavoriteCount field.Int64
	CoverURL      field.String
	PlayURL       field.String

	fieldMap map[string]field.Expr
}

func (v videoInfo) Table(newTableName string) *videoInfo {
	v.videoInfoDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v videoInfo) As(alias string) *videoInfo {
	v.videoInfoDo.DO = *(v.videoInfoDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *videoInfo) updateTableName(table string) *videoInfo {
	v.ALL = field.NewAsterisk(table)
	v.VideoID = field.NewInt64(table, "video_id")
	v.AuthID = field.NewInt64(table, "auth_id")
	v.Titel = field.NewString(table, "titel")
	v.CommentCount = field.NewInt64(table, "comment_count")
	v.FavoriteCount = field.NewInt64(table, "favorite_count")
	v.CoverURL = field.NewString(table, "cover_url")
	v.PlayURL = field.NewString(table, "play_url")

	v.fillFieldMap()

	return v
}

func (v *videoInfo) WithContext(ctx context.Context) IVideoInfoDo {
	return v.videoInfoDo.WithContext(ctx)
}

func (v videoInfo) TableName() string { return v.videoInfoDo.TableName() }

func (v videoInfo) Alias() string { return v.videoInfoDo.Alias() }

func (v *videoInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *videoInfo) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 7)
	v.fieldMap["video_id"] = v.VideoID
	v.fieldMap["auth_id"] = v.AuthID
	v.fieldMap["titel"] = v.Titel
	v.fieldMap["comment_count"] = v.CommentCount
	v.fieldMap["favorite_count"] = v.FavoriteCount
	v.fieldMap["cover_url"] = v.CoverURL
	v.fieldMap["play_url"] = v.PlayURL
}

func (v videoInfo) clone(db *gorm.DB) videoInfo {
	v.videoInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v videoInfo) replaceDB(db *gorm.DB) videoInfo {
	v.videoInfoDo.ReplaceDB(db)
	return v
}

type videoInfoDo struct{ gen.DO }

type IVideoInfoDo interface {
	gen.SubQuery
	Debug() IVideoInfoDo
	WithContext(ctx context.Context) IVideoInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVideoInfoDo
	WriteDB() IVideoInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVideoInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVideoInfoDo
	Not(conds ...gen.Condition) IVideoInfoDo
	Or(conds ...gen.Condition) IVideoInfoDo
	Select(conds ...field.Expr) IVideoInfoDo
	Where(conds ...gen.Condition) IVideoInfoDo
	Order(conds ...field.Expr) IVideoInfoDo
	Distinct(cols ...field.Expr) IVideoInfoDo
	Omit(cols ...field.Expr) IVideoInfoDo
	Join(table schema.Tabler, on ...field.Expr) IVideoInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVideoInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVideoInfoDo
	Group(cols ...field.Expr) IVideoInfoDo
	Having(conds ...gen.Condition) IVideoInfoDo
	Limit(limit int) IVideoInfoDo
	Offset(offset int) IVideoInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVideoInfoDo
	Unscoped() IVideoInfoDo
	Create(values ...*model.VideoInfo) error
	CreateInBatches(values []*model.VideoInfo, batchSize int) error
	Save(values ...*model.VideoInfo) error
	First() (*model.VideoInfo, error)
	Take() (*model.VideoInfo, error)
	Last() (*model.VideoInfo, error)
	Find() ([]*model.VideoInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoInfo, err error)
	FindInBatches(result *[]*model.VideoInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VideoInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVideoInfoDo
	Assign(attrs ...field.AssignExpr) IVideoInfoDo
	Joins(fields ...field.RelationField) IVideoInfoDo
	Preload(fields ...field.RelationField) IVideoInfoDo
	FirstOrInit() (*model.VideoInfo, error)
	FirstOrCreate() (*model.VideoInfo, error)
	FindByPage(offset int, limit int) (result []*model.VideoInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVideoInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v videoInfoDo) Debug() IVideoInfoDo {
	return v.withDO(v.DO.Debug())
}

func (v videoInfoDo) WithContext(ctx context.Context) IVideoInfoDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoInfoDo) ReadDB() IVideoInfoDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoInfoDo) WriteDB() IVideoInfoDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoInfoDo) Session(config *gorm.Session) IVideoInfoDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoInfoDo) Clauses(conds ...clause.Expression) IVideoInfoDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoInfoDo) Returning(value interface{}, columns ...string) IVideoInfoDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoInfoDo) Not(conds ...gen.Condition) IVideoInfoDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoInfoDo) Or(conds ...gen.Condition) IVideoInfoDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoInfoDo) Select(conds ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoInfoDo) Where(conds ...gen.Condition) IVideoInfoDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IVideoInfoDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v videoInfoDo) Order(conds ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoInfoDo) Distinct(cols ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoInfoDo) Omit(cols ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoInfoDo) Join(table schema.Tabler, on ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoInfoDo) Group(cols ...field.Expr) IVideoInfoDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoInfoDo) Having(conds ...gen.Condition) IVideoInfoDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoInfoDo) Limit(limit int) IVideoInfoDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoInfoDo) Offset(offset int) IVideoInfoDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVideoInfoDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoInfoDo) Unscoped() IVideoInfoDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoInfoDo) Create(values ...*model.VideoInfo) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoInfoDo) CreateInBatches(values []*model.VideoInfo, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoInfoDo) Save(values ...*model.VideoInfo) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoInfoDo) First() (*model.VideoInfo, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoInfo), nil
	}
}

func (v videoInfoDo) Take() (*model.VideoInfo, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoInfo), nil
	}
}

func (v videoInfoDo) Last() (*model.VideoInfo, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoInfo), nil
	}
}

func (v videoInfoDo) Find() ([]*model.VideoInfo, error) {
	result, err := v.DO.Find()
	return result.([]*model.VideoInfo), err
}

func (v videoInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoInfo, err error) {
	buf := make([]*model.VideoInfo, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoInfoDo) FindInBatches(result *[]*model.VideoInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoInfoDo) Attrs(attrs ...field.AssignExpr) IVideoInfoDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoInfoDo) Assign(attrs ...field.AssignExpr) IVideoInfoDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoInfoDo) Joins(fields ...field.RelationField) IVideoInfoDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoInfoDo) Preload(fields ...field.RelationField) IVideoInfoDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoInfoDo) FirstOrInit() (*model.VideoInfo, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoInfo), nil
	}
}

func (v videoInfoDo) FirstOrCreate() (*model.VideoInfo, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoInfo), nil
	}
}

func (v videoInfoDo) FindByPage(offset int, limit int) (result []*model.VideoInfo, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v videoInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoInfoDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoInfoDo) Delete(models ...*model.VideoInfo) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoInfoDo) withDO(do gen.Dao) *videoInfoDo {
	v.DO = *do.(*gen.DO)
	return v
}
