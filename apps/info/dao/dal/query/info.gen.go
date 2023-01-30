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

	"github.com/Go-To-Byte/DouSheng/apps/user/dao/dal/model"
)

func newInfo(db *gorm.DB, opts ...gen.DOOption) info {
	_info := info{}

	_info.infoDo.UseDB(db, opts...)
	_info.infoDo.UseModel(&model.Info{})

	tableName := _info.infoDo.TableName()
	_info.ALL = field.NewAsterisk(tableName)
	_info.ID = field.NewInt64(tableName, "id")
	_info.Phone = field.NewString(tableName, "phone")
	_info.Name = field.NewString(tableName, "name")
	_info.FollowCount = field.NewInt64(tableName, "follow_count")
	_info.FollowerCount = field.NewInt64(tableName, "follower_count")

	_info.fillFieldMap()

	return _info
}

type info struct {
	infoDo

	ALL           field.Asterisk
	ID            field.Int64
	Phone         field.String
	Name          field.String
	FollowCount   field.Int64
	FollowerCount field.Int64

	fieldMap map[string]field.Expr
}

func (i info) Table(newTableName string) *info {
	i.infoDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i info) As(alias string) *info {
	i.infoDo.DO = *(i.infoDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *info) updateTableName(table string) *info {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.Phone = field.NewString(table, "phone")
	i.Name = field.NewString(table, "name")
	i.FollowCount = field.NewInt64(table, "follow_count")
	i.FollowerCount = field.NewInt64(table, "follower_count")

	i.fillFieldMap()

	return i
}

func (i *info) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *info) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 5)
	i.fieldMap["id"] = i.ID
	i.fieldMap["phone"] = i.Phone
	i.fieldMap["name"] = i.Name
	i.fieldMap["follow_count"] = i.FollowCount
	i.fieldMap["follower_count"] = i.FollowerCount
}

func (i info) clone(db *gorm.DB) info {
	i.infoDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i info) replaceDB(db *gorm.DB) info {
	i.infoDo.ReplaceDB(db)
	return i
}

type infoDo struct{ gen.DO }

type IInfoDo interface {
	gen.SubQuery
	Debug() IInfoDo
	WithContext(ctx context.Context) IInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IInfoDo
	WriteDB() IInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInfoDo
	Not(conds ...gen.Condition) IInfoDo
	Or(conds ...gen.Condition) IInfoDo
	Select(conds ...field.Expr) IInfoDo
	Where(conds ...gen.Condition) IInfoDo
	Order(conds ...field.Expr) IInfoDo
	Distinct(cols ...field.Expr) IInfoDo
	Omit(cols ...field.Expr) IInfoDo
	Join(table schema.Tabler, on ...field.Expr) IInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInfoDo
	Group(cols ...field.Expr) IInfoDo
	Having(conds ...gen.Condition) IInfoDo
	Limit(limit int) IInfoDo
	Offset(offset int) IInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInfoDo
	Unscoped() IInfoDo
	Create(values ...*model.Info) error
	CreateInBatches(values []*model.Info, batchSize int) error
	Save(values ...*model.Info) error
	First() (*model.Info, error)
	Take() (*model.Info, error)
	Last() (*model.Info, error)
	Find() ([]*model.Info, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Info, err error)
	FindInBatches(result *[]*model.Info, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Info) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInfoDo
	Assign(attrs ...field.AssignExpr) IInfoDo
	Joins(fields ...field.RelationField) IInfoDo
	Preload(fields ...field.RelationField) IInfoDo
	FirstOrInit() (*model.Info, error)
	FirstOrCreate() (*model.Info, error)
	FindByPage(offset int, limit int) (result []*model.Info, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i infoDo) Debug() IInfoDo {
	return i.withDO(i.DO.Debug())
}

func (i infoDo) WithContext(ctx context.Context) IInfoDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i infoDo) ReadDB() IInfoDo {
	return i.Clauses(dbresolver.Read)
}

func (i infoDo) WriteDB() IInfoDo {
	return i.Clauses(dbresolver.Write)
}

func (i infoDo) Session(config *gorm.Session) IInfoDo {
	return i.withDO(i.DO.Session(config))
}

func (i infoDo) Clauses(conds ...clause.Expression) IInfoDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i infoDo) Returning(value interface{}, columns ...string) IInfoDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i infoDo) Not(conds ...gen.Condition) IInfoDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i infoDo) Or(conds ...gen.Condition) IInfoDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i infoDo) Select(conds ...field.Expr) IInfoDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i infoDo) Where(conds ...gen.Condition) IInfoDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i infoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IInfoDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i infoDo) Order(conds ...field.Expr) IInfoDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i infoDo) Distinct(cols ...field.Expr) IInfoDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i infoDo) Omit(cols ...field.Expr) IInfoDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i infoDo) Join(table schema.Tabler, on ...field.Expr) IInfoDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i infoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInfoDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i infoDo) RightJoin(table schema.Tabler, on ...field.Expr) IInfoDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i infoDo) Group(cols ...field.Expr) IInfoDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i infoDo) Having(conds ...gen.Condition) IInfoDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i infoDo) Limit(limit int) IInfoDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i infoDo) Offset(offset int) IInfoDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i infoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInfoDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i infoDo) Unscoped() IInfoDo {
	return i.withDO(i.DO.Unscoped())
}

func (i infoDo) Create(values ...*model.Info) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i infoDo) CreateInBatches(values []*model.Info, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i infoDo) Save(values ...*model.Info) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i infoDo) First() (*model.Info, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Info), nil
	}
}

func (i infoDo) Take() (*model.Info, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Info), nil
	}
}

func (i infoDo) Last() (*model.Info, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Info), nil
	}
}

func (i infoDo) Find() ([]*model.Info, error) {
	result, err := i.DO.Find()
	return result.([]*model.Info), err
}

func (i infoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Info, err error) {
	buf := make([]*model.Info, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i infoDo) FindInBatches(result *[]*model.Info, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i infoDo) Attrs(attrs ...field.AssignExpr) IInfoDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i infoDo) Assign(attrs ...field.AssignExpr) IInfoDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i infoDo) Joins(fields ...field.RelationField) IInfoDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i infoDo) Preload(fields ...field.RelationField) IInfoDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i infoDo) FirstOrInit() (*model.Info, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Info), nil
	}
}

func (i infoDo) FirstOrCreate() (*model.Info, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Info), nil
	}
}

func (i infoDo) FindByPage(offset int, limit int) (result []*model.Info, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i infoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i infoDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i infoDo) Delete(models ...*model.Info) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *infoDo) withDO(do gen.Dao) *infoDo {
	i.DO = *do.(*gen.DO)
	return i
}
