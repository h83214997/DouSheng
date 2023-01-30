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

	"github.com/Go-To-Byte/DouSheng/apps/info/dao/dal/model"
)

func newFollower(db *gorm.DB, opts ...gen.DOOption) follower {
	_follower := follower{}

	_follower.followerDo.UseDB(db, opts...)
	_follower.followerDo.UseModel(&model.Follower{})

	tableName := _follower.followerDo.TableName()
	_follower.ALL = field.NewAsterisk(tableName)
	_follower.Follower = field.NewInt64(tableName, "follower")
	_follower.ToFollower = field.NewInt64(tableName, "to_follower")
	_follower.Flag = field.NewInt64(tableName, "flag")

	_follower.fillFieldMap()

	return _follower
}

type follower struct {
	followerDo

	ALL        field.Asterisk
	Follower   field.Int64
	ToFollower field.Int64
	Flag       field.Int64

	fieldMap map[string]field.Expr
}

func (f follower) Table(newTableName string) *follower {
	f.followerDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f follower) As(alias string) *follower {
	f.followerDo.DO = *(f.followerDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *follower) updateTableName(table string) *follower {
	f.ALL = field.NewAsterisk(table)
	f.Follower = field.NewInt64(table, "follower")
	f.ToFollower = field.NewInt64(table, "to_follower")
	f.Flag = field.NewInt64(table, "flag")

	f.fillFieldMap()

	return f
}

func (f *follower) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *follower) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 3)
	f.fieldMap["follower"] = f.Follower
	f.fieldMap["to_follower"] = f.ToFollower
	f.fieldMap["flag"] = f.Flag
}

func (f follower) clone(db *gorm.DB) follower {
	f.followerDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f follower) replaceDB(db *gorm.DB) follower {
	f.followerDo.ReplaceDB(db)
	return f
}

type followerDo struct{ gen.DO }

type IFollowerDo interface {
	gen.SubQuery
	Debug() IFollowerDo
	WithContext(ctx context.Context) IFollowerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFollowerDo
	WriteDB() IFollowerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFollowerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFollowerDo
	Not(conds ...gen.Condition) IFollowerDo
	Or(conds ...gen.Condition) IFollowerDo
	Select(conds ...field.Expr) IFollowerDo
	Where(conds ...gen.Condition) IFollowerDo
	Order(conds ...field.Expr) IFollowerDo
	Distinct(cols ...field.Expr) IFollowerDo
	Omit(cols ...field.Expr) IFollowerDo
	Join(table schema.Tabler, on ...field.Expr) IFollowerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFollowerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFollowerDo
	Group(cols ...field.Expr) IFollowerDo
	Having(conds ...gen.Condition) IFollowerDo
	Limit(limit int) IFollowerDo
	Offset(offset int) IFollowerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowerDo
	Unscoped() IFollowerDo
	Create(values ...*model.Follower) error
	CreateInBatches(values []*model.Follower, batchSize int) error
	Save(values ...*model.Follower) error
	First() (*model.Follower, error)
	Take() (*model.Follower, error)
	Last() (*model.Follower, error)
	Find() ([]*model.Follower, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Follower, err error)
	FindInBatches(result *[]*model.Follower, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Follower) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFollowerDo
	Assign(attrs ...field.AssignExpr) IFollowerDo
	Joins(fields ...field.RelationField) IFollowerDo
	Preload(fields ...field.RelationField) IFollowerDo
	FirstOrInit() (*model.Follower, error)
	FirstOrCreate() (*model.Follower, error)
	FindByPage(offset int, limit int) (result []*model.Follower, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFollowerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f followerDo) Debug() IFollowerDo {
	return f.withDO(f.DO.Debug())
}

func (f followerDo) WithContext(ctx context.Context) IFollowerDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f followerDo) ReadDB() IFollowerDo {
	return f.Clauses(dbresolver.Read)
}

func (f followerDo) WriteDB() IFollowerDo {
	return f.Clauses(dbresolver.Write)
}

func (f followerDo) Session(config *gorm.Session) IFollowerDo {
	return f.withDO(f.DO.Session(config))
}

func (f followerDo) Clauses(conds ...clause.Expression) IFollowerDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f followerDo) Returning(value interface{}, columns ...string) IFollowerDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f followerDo) Not(conds ...gen.Condition) IFollowerDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f followerDo) Or(conds ...gen.Condition) IFollowerDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f followerDo) Select(conds ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f followerDo) Where(conds ...gen.Condition) IFollowerDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f followerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IFollowerDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f followerDo) Order(conds ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f followerDo) Distinct(cols ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f followerDo) Omit(cols ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f followerDo) Join(table schema.Tabler, on ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f followerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f followerDo) RightJoin(table schema.Tabler, on ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f followerDo) Group(cols ...field.Expr) IFollowerDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f followerDo) Having(conds ...gen.Condition) IFollowerDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f followerDo) Limit(limit int) IFollowerDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f followerDo) Offset(offset int) IFollowerDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f followerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFollowerDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f followerDo) Unscoped() IFollowerDo {
	return f.withDO(f.DO.Unscoped())
}

func (f followerDo) Create(values ...*model.Follower) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f followerDo) CreateInBatches(values []*model.Follower, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f followerDo) Save(values ...*model.Follower) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f followerDo) First() (*model.Follower, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follower), nil
	}
}

func (f followerDo) Take() (*model.Follower, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follower), nil
	}
}

func (f followerDo) Last() (*model.Follower, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follower), nil
	}
}

func (f followerDo) Find() ([]*model.Follower, error) {
	result, err := f.DO.Find()
	return result.([]*model.Follower), err
}

func (f followerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Follower, err error) {
	buf := make([]*model.Follower, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f followerDo) FindInBatches(result *[]*model.Follower, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f followerDo) Attrs(attrs ...field.AssignExpr) IFollowerDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f followerDo) Assign(attrs ...field.AssignExpr) IFollowerDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f followerDo) Joins(fields ...field.RelationField) IFollowerDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f followerDo) Preload(fields ...field.RelationField) IFollowerDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f followerDo) FirstOrInit() (*model.Follower, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follower), nil
	}
}

func (f followerDo) FirstOrCreate() (*model.Follower, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Follower), nil
	}
}

func (f followerDo) FindByPage(offset int, limit int) (result []*model.Follower, count int64, err error) {
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

func (f followerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f followerDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f followerDo) Delete(models ...*model.Follower) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *followerDo) withDO(do gen.Dao) *followerDo {
	f.DO = *do.(*gen.DO)
	return f
}
