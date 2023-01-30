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

	"github.com/Go-To-Byte/DouSheng/apps/relation/dao/dal/model"
)

func newFollow(db *gorm.DB, opts ...gen.DOOption) follow {
	_follow := follow{}

	_follow.followDo.UseDB(db, opts...)
	_follow.followDo.UseModel(&model.Follow{})

	tableName := _follow.followDo.TableName()
	_follow.ALL = field.NewAsterisk(tableName)
	_follow.UserID = field.NewInt64(tableName, "user_id")
	_follow.ToUserID = field.NewInt64(tableName, "to_user_id")
	_follow.Flag = field.NewInt64(tableName, "flag")

	_follow.fillFieldMap()

	return _follow
}

type follow struct {
	followDo

	ALL      field.Asterisk
	UserID   field.Int64
	ToUserID field.Int64
	Flag     field.Int64

	fieldMap map[string]field.Expr
}

func (f follow) Table(newTableName string) *follow {
	f.followDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f follow) As(alias string) *follow {
	f.followDo.DO = *(f.followDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *follow) updateTableName(table string) *follow {
	f.ALL = field.NewAsterisk(table)
	f.UserID = field.NewInt64(table, "user_id")
	f.ToUserID = field.NewInt64(table, "to_user_id")
	f.Flag = field.NewInt64(table, "flag")

	f.fillFieldMap()

	return f
}

func (f *follow) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *follow) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 3)
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["to_user_id"] = f.ToUserID
	f.fieldMap["flag"] = f.Flag
}

func (f follow) clone(db *gorm.DB) follow {
	f.followDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f follow) replaceDB(db *gorm.DB) follow {
	f.followDo.ReplaceDB(db)
	return f
}

type followDo struct{ gen.DO }

type IFollowDo interface {
	gen.SubQuery
	Debug() IFollowDo
	WithContext(ctx context.Context) IFollowDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFollowDo
	WriteDB() IFollowDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFollowDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFollowDo
	Not(conds ...gen.Condition) IFollowDo
	Or(conds ...gen.Condition) IFollowDo
	Select(conds ...field.Expr) IFollowDo
	Where(conds ...gen.Condition) IFollowDo
	Order(conds ...field.Expr) IFollowDo
	Distinct(cols ...field.Expr) IFollowDo
	Omit(cols ...field.Expr) IFollowDo
	Join(table schema.Tabler, on ...field.Expr) IFollowDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFollowDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFollowDo
	Group(cols ...field.Expr) IFollowDo
	Having(conds ...gen.Condition) IFollowDo
	Limit(limit int) IFollowDo
	Offset(offset int) IFollowDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowDo
	Unscoped() IFollowDo
	Create(values ...*model.Follow) error
	CreateInBatches(values []*model.Follow, batchSize int) error
	Save(values ...*model.Follow) error
	First() (*model.Follow, error)
	Take() (*model.Follow, error)
	Last() (*model.Follow, error)
	Find() ([]*model.Follow, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Follow, err error)
	FindInBatches(result *[]*model.Follow, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Follow) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFollowDo
	Assign(attrs ...field.AssignExpr) IFollowDo
	Joins(fields ...field.RelationField) IFollowDo
	Preload(fields ...field.RelationField) IFollowDo
	FirstOrInit() (*model.Follow, error)
	FirstOrCreate() (*model.Follow, error)
	FindByPage(offset int, limit int) (result []*model.Follow, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFollowDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f followDo) Debug() IFollowDo {
	return f.withDO(f.DO.Debug())
}

func (f followDo) WithContext(ctx context.Context) IFollowDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followDo) ReadDB() IFollowDo {
	return f.Clauses(dbresolver.Read)
}

func (f followDo) WriteDB() IFollowDo {
	return f.Clauses(dbresolver.Write)
}

func (f followDo) Session(config *gorm.Session) IFollowDo {
	return f.withDO(f.DO.Session(config))
}

func (f followDo) Clauses(conds ...clause.Expression) IFollowDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followDo) Returning(value interface{}, columns ...string) IFollowDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followDo) Not(conds ...gen.Condition) IFollowDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followDo) Or(conds ...gen.Condition) IFollowDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followDo) Select(conds ...field.Expr) IFollowDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followDo) Where(conds ...gen.Condition) IFollowDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFollowDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f followDo) Order(conds ...field.Expr) IFollowDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followDo) Distinct(cols ...field.Expr) IFollowDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followDo) Omit(cols ...field.Expr) IFollowDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followDo) Join(table schema.Tabler, on ...field.Expr) IFollowDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFollowDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followDo) RightJoin(table schema.Tabler, on ...field.Expr) IFollowDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followDo) Group(cols ...field.Expr) IFollowDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followDo) Having(conds ...gen.Condition) IFollowDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followDo) Limit(limit int) IFollowDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followDo) Offset(offset int) IFollowDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followDo) Unscoped() IFollowDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followDo) Create(values ...*model.Follow) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followDo) CreateInBatches(values []*model.Follow, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followDo) Save(values ...*model.Follow) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followDo) First() (*model.Follow, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follow), nil
	}
}

func (f followDo) Take() (*model.Follow, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follow), nil
	}
}

func (f followDo) Last() (*model.Follow, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follow), nil
	}
}

func (f followDo) Find() ([]*model.Follow, error) {
	result, err := f.DO.Find()
	return result.([]*model.Follow), err
}

func (f followDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Follow, err error) {
	buf := make([]*model.Follow, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followDo) FindInBatches(result *[]*model.Follow, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followDo) Attrs(attrs ...field.AssignExpr) IFollowDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followDo) Assign(attrs ...field.AssignExpr) IFollowDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followDo) Joins(fields ...field.RelationField) IFollowDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followDo) Preload(fields ...field.RelationField) IFollowDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followDo) FirstOrInit() (*model.Follow, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follow), nil
	}
}

func (f followDo) FirstOrCreate() (*model.Follow, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follow), nil
	}
}

func (f followDo) FindByPage(offset int, limit int) (result []*model.Follow, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f followDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followDo) Delete(models ...*model.Follow) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followDo) withDO(do gen.Dao) *followDo {
	f.DO = *do.(*gen.DO)
	return f
}
