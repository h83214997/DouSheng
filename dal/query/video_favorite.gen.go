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

func newVideoFavorite(db *gorm.DB, opts ...gen.DOOption) videoFavorite {
	_videoFavorite := videoFavorite{}

	_videoFavorite.videoFavoriteDo.UseDB(db, opts...)
	_videoFavorite.videoFavoriteDo.UseModel(&model.VideoFavorite{})

	tableName := _videoFavorite.videoFavoriteDo.TableName()
	_videoFavorite.ALL = field.NewAsterisk(tableName)
	_videoFavorite.FavoriteUser = field.NewInt64(tableName, "favorite_user")
	_videoFavorite.FavoriteVideo = field.NewInt64(tableName, "favorite_video")
	_videoFavorite.FavoriteFlag = field.NewInt64(tableName, "favorite_flag")

	_videoFavorite.fillFieldMap()

	return _videoFavorite
}

type videoFavorite struct {
	videoFavoriteDo videoFavoriteDo

	ALL           field.Asterisk
	FavoriteUser  field.Int64
	FavoriteVideo field.Int64
	FavoriteFlag  field.Int64

	fieldMap map[string]field.Expr
}

func (v videoFavorite) Table(newTableName string) *videoFavorite {
	v.videoFavoriteDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v videoFavorite) As(alias string) *videoFavorite {
	v.videoFavoriteDo.DO = *(v.videoFavoriteDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *videoFavorite) updateTableName(table string) *videoFavorite {
	v.ALL = field.NewAsterisk(table)
	v.FavoriteUser = field.NewInt64(table, "favorite_user")
	v.FavoriteVideo = field.NewInt64(table, "favorite_video")
	v.FavoriteFlag = field.NewInt64(table, "favorite_flag")

	v.fillFieldMap()

	return v
}

func (v *videoFavorite) WithContext(ctx context.Context) IVideoFavoriteDo {
	return v.videoFavoriteDo.WithContext(ctx)
}

func (v videoFavorite) TableName() string { return v.videoFavoriteDo.TableName() }

func (v videoFavorite) Alias() string { return v.videoFavoriteDo.Alias() }

func (v *videoFavorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *videoFavorite) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 3)
	v.fieldMap["favorite_user"] = v.FavoriteUser
	v.fieldMap["favorite_video"] = v.FavoriteVideo
	v.fieldMap["favorite_flag"] = v.FavoriteFlag
}

func (v videoFavorite) clone(db *gorm.DB) videoFavorite {
	v.videoFavoriteDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v videoFavorite) replaceDB(db *gorm.DB) videoFavorite {
	v.videoFavoriteDo.ReplaceDB(db)
	return v
}

type videoFavoriteDo struct{ gen.DO }

type IVideoFavoriteDo interface {
	gen.SubQuery
	Debug() IVideoFavoriteDo
	WithContext(ctx context.Context) IVideoFavoriteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVideoFavoriteDo
	WriteDB() IVideoFavoriteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVideoFavoriteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVideoFavoriteDo
	Not(conds ...gen.Condition) IVideoFavoriteDo
	Or(conds ...gen.Condition) IVideoFavoriteDo
	Select(conds ...field.Expr) IVideoFavoriteDo
	Where(conds ...gen.Condition) IVideoFavoriteDo
	Order(conds ...field.Expr) IVideoFavoriteDo
	Distinct(cols ...field.Expr) IVideoFavoriteDo
	Omit(cols ...field.Expr) IVideoFavoriteDo
	Join(table schema.Tabler, on ...field.Expr) IVideoFavoriteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVideoFavoriteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVideoFavoriteDo
	Group(cols ...field.Expr) IVideoFavoriteDo
	Having(conds ...gen.Condition) IVideoFavoriteDo
	Limit(limit int) IVideoFavoriteDo
	Offset(offset int) IVideoFavoriteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVideoFavoriteDo
	Unscoped() IVideoFavoriteDo
	Create(values ...*model.VideoFavorite) error
	CreateInBatches(values []*model.VideoFavorite, batchSize int) error
	Save(values ...*model.VideoFavorite) error
	First() (*model.VideoFavorite, error)
	Take() (*model.VideoFavorite, error)
	Last() (*model.VideoFavorite, error)
	Find() ([]*model.VideoFavorite, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoFavorite, err error)
	FindInBatches(result *[]*model.VideoFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.VideoFavorite) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVideoFavoriteDo
	Assign(attrs ...field.AssignExpr) IVideoFavoriteDo
	Joins(fields ...field.RelationField) IVideoFavoriteDo
	Preload(fields ...field.RelationField) IVideoFavoriteDo
	FirstOrInit() (*model.VideoFavorite, error)
	FirstOrCreate() (*model.VideoFavorite, error)
	FindByPage(offset int, limit int) (result []*model.VideoFavorite, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVideoFavoriteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v videoFavoriteDo) Debug() IVideoFavoriteDo {
	return v.withDO(v.DO.Debug())
}

func (v videoFavoriteDo) WithContext(ctx context.Context) IVideoFavoriteDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v videoFavoriteDo) ReadDB() IVideoFavoriteDo {
	return v.Clauses(dbresolver.Read)
}

func (v videoFavoriteDo) WriteDB() IVideoFavoriteDo {
	return v.Clauses(dbresolver.Write)
}

func (v videoFavoriteDo) Session(config *gorm.Session) IVideoFavoriteDo {
	return v.withDO(v.DO.Session(config))
}

func (v videoFavoriteDo) Clauses(conds ...clause.Expression) IVideoFavoriteDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v videoFavoriteDo) Returning(value interface{}, columns ...string) IVideoFavoriteDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v videoFavoriteDo) Not(conds ...gen.Condition) IVideoFavoriteDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v videoFavoriteDo) Or(conds ...gen.Condition) IVideoFavoriteDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v videoFavoriteDo) Select(conds ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v videoFavoriteDo) Where(conds ...gen.Condition) IVideoFavoriteDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v videoFavoriteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IVideoFavoriteDo {
	return v.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (v videoFavoriteDo) Order(conds ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v videoFavoriteDo) Distinct(cols ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v videoFavoriteDo) Omit(cols ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v videoFavoriteDo) Join(table schema.Tabler, on ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v videoFavoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v videoFavoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v videoFavoriteDo) Group(cols ...field.Expr) IVideoFavoriteDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v videoFavoriteDo) Having(conds ...gen.Condition) IVideoFavoriteDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v videoFavoriteDo) Limit(limit int) IVideoFavoriteDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v videoFavoriteDo) Offset(offset int) IVideoFavoriteDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v videoFavoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVideoFavoriteDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v videoFavoriteDo) Unscoped() IVideoFavoriteDo {
	return v.withDO(v.DO.Unscoped())
}

func (v videoFavoriteDo) Create(values ...*model.VideoFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v videoFavoriteDo) CreateInBatches(values []*model.VideoFavorite, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v videoFavoriteDo) Save(values ...*model.VideoFavorite) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v videoFavoriteDo) First() (*model.VideoFavorite, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoFavorite), nil
	}
}

func (v videoFavoriteDo) Take() (*model.VideoFavorite, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoFavorite), nil
	}
}

func (v videoFavoriteDo) Last() (*model.VideoFavorite, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoFavorite), nil
	}
}

func (v videoFavoriteDo) Find() ([]*model.VideoFavorite, error) {
	result, err := v.DO.Find()
	return result.([]*model.VideoFavorite), err
}

func (v videoFavoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.VideoFavorite, err error) {
	buf := make([]*model.VideoFavorite, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v videoFavoriteDo) FindInBatches(result *[]*model.VideoFavorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v videoFavoriteDo) Attrs(attrs ...field.AssignExpr) IVideoFavoriteDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v videoFavoriteDo) Assign(attrs ...field.AssignExpr) IVideoFavoriteDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v videoFavoriteDo) Joins(fields ...field.RelationField) IVideoFavoriteDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v videoFavoriteDo) Preload(fields ...field.RelationField) IVideoFavoriteDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v videoFavoriteDo) FirstOrInit() (*model.VideoFavorite, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoFavorite), nil
	}
}

func (v videoFavoriteDo) FirstOrCreate() (*model.VideoFavorite, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.VideoFavorite), nil
	}
}

func (v videoFavoriteDo) FindByPage(offset int, limit int) (result []*model.VideoFavorite, count int64, err error) {
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

func (v videoFavoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v videoFavoriteDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v videoFavoriteDo) Delete(models ...*model.VideoFavorite) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *videoFavoriteDo) withDO(do gen.Dao) *videoFavoriteDo {
	v.DO = *do.(*gen.DO)
	return v
}
