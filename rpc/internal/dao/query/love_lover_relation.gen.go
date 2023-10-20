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

	"github.com/Jing-Pattern/user/rpc/internal/dao/model"
)

func newLoveLoverRelation(db *gorm.DB, opts ...gen.DOOption) loveLoverRelation {
	_loveLoverRelation := loveLoverRelation{}

	_loveLoverRelation.loveLoverRelationDo.UseDB(db, opts...)
	_loveLoverRelation.loveLoverRelationDo.UseModel(&model.LoveLoverRelation{})

	tableName := _loveLoverRelation.loveLoverRelationDo.TableName()
	_loveLoverRelation.ALL = field.NewAsterisk(tableName)
	_loveLoverRelation.ID = field.NewInt32(tableName, "id")
	_loveLoverRelation.User1ID = field.NewString(tableName, "user1_id")
	_loveLoverRelation.User2ID = field.NewString(tableName, "user2_id")
	_loveLoverRelation.User1Nickname = field.NewString(tableName, "user1_nickname")
	_loveLoverRelation.User2Nickname = field.NewString(tableName, "user2_nickname")

	_loveLoverRelation.fillFieldMap()

	return _loveLoverRelation
}

type loveLoverRelation struct {
	loveLoverRelationDo loveLoverRelationDo

	ALL           field.Asterisk
	ID            field.Int32  // id
	User1ID       field.String // user1_id
	User2ID       field.String // user2_id
	User1Nickname field.String // 昵称
	User2Nickname field.String // 昵称

	fieldMap map[string]field.Expr
}

func (l loveLoverRelation) Table(newTableName string) *loveLoverRelation {
	l.loveLoverRelationDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l loveLoverRelation) As(alias string) *loveLoverRelation {
	l.loveLoverRelationDo.DO = *(l.loveLoverRelationDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *loveLoverRelation) updateTableName(table string) *loveLoverRelation {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt32(table, "id")
	l.User1ID = field.NewString(table, "user1_id")
	l.User2ID = field.NewString(table, "user2_id")
	l.User1Nickname = field.NewString(table, "user1_nickname")
	l.User2Nickname = field.NewString(table, "user2_nickname")

	l.fillFieldMap()

	return l
}

func (l *loveLoverRelation) WithContext(ctx context.Context) *loveLoverRelationDo {
	return l.loveLoverRelationDo.WithContext(ctx)
}

func (l loveLoverRelation) TableName() string { return l.loveLoverRelationDo.TableName() }

func (l loveLoverRelation) Alias() string { return l.loveLoverRelationDo.Alias() }

func (l loveLoverRelation) Columns(cols ...field.Expr) gen.Columns {
	return l.loveLoverRelationDo.Columns(cols...)
}

func (l *loveLoverRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *loveLoverRelation) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 5)
	l.fieldMap["id"] = l.ID
	l.fieldMap["user1_id"] = l.User1ID
	l.fieldMap["user2_id"] = l.User2ID
	l.fieldMap["user1_nickname"] = l.User1Nickname
	l.fieldMap["user2_nickname"] = l.User2Nickname
}

func (l loveLoverRelation) clone(db *gorm.DB) loveLoverRelation {
	l.loveLoverRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l loveLoverRelation) replaceDB(db *gorm.DB) loveLoverRelation {
	l.loveLoverRelationDo.ReplaceDB(db)
	return l
}

type loveLoverRelationDo struct{ gen.DO }

func (l loveLoverRelationDo) Debug() *loveLoverRelationDo {
	return l.withDO(l.DO.Debug())
}

func (l loveLoverRelationDo) WithContext(ctx context.Context) *loveLoverRelationDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l loveLoverRelationDo) ReadDB() *loveLoverRelationDo {
	return l.Clauses(dbresolver.Read)
}

func (l loveLoverRelationDo) WriteDB() *loveLoverRelationDo {
	return l.Clauses(dbresolver.Write)
}

func (l loveLoverRelationDo) Session(config *gorm.Session) *loveLoverRelationDo {
	return l.withDO(l.DO.Session(config))
}

func (l loveLoverRelationDo) Clauses(conds ...clause.Expression) *loveLoverRelationDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l loveLoverRelationDo) Returning(value interface{}, columns ...string) *loveLoverRelationDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l loveLoverRelationDo) Not(conds ...gen.Condition) *loveLoverRelationDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l loveLoverRelationDo) Or(conds ...gen.Condition) *loveLoverRelationDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l loveLoverRelationDo) Select(conds ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l loveLoverRelationDo) Where(conds ...gen.Condition) *loveLoverRelationDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l loveLoverRelationDo) Order(conds ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l loveLoverRelationDo) Distinct(cols ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l loveLoverRelationDo) Omit(cols ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l loveLoverRelationDo) Join(table schema.Tabler, on ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l loveLoverRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l loveLoverRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l loveLoverRelationDo) Group(cols ...field.Expr) *loveLoverRelationDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l loveLoverRelationDo) Having(conds ...gen.Condition) *loveLoverRelationDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l loveLoverRelationDo) Limit(limit int) *loveLoverRelationDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l loveLoverRelationDo) Offset(offset int) *loveLoverRelationDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l loveLoverRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *loveLoverRelationDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l loveLoverRelationDo) Unscoped() *loveLoverRelationDo {
	return l.withDO(l.DO.Unscoped())
}

func (l loveLoverRelationDo) Create(values ...*model.LoveLoverRelation) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l loveLoverRelationDo) CreateInBatches(values []*model.LoveLoverRelation, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l loveLoverRelationDo) Save(values ...*model.LoveLoverRelation) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l loveLoverRelationDo) First() (*model.LoveLoverRelation, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoveLoverRelation), nil
	}
}

func (l loveLoverRelationDo) Take() (*model.LoveLoverRelation, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoveLoverRelation), nil
	}
}

func (l loveLoverRelationDo) Last() (*model.LoveLoverRelation, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoveLoverRelation), nil
	}
}

func (l loveLoverRelationDo) Find() ([]*model.LoveLoverRelation, error) {
	result, err := l.DO.Find()
	return result.([]*model.LoveLoverRelation), err
}

func (l loveLoverRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LoveLoverRelation, err error) {
	buf := make([]*model.LoveLoverRelation, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l loveLoverRelationDo) FindInBatches(result *[]*model.LoveLoverRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l loveLoverRelationDo) Attrs(attrs ...field.AssignExpr) *loveLoverRelationDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l loveLoverRelationDo) Assign(attrs ...field.AssignExpr) *loveLoverRelationDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l loveLoverRelationDo) Joins(fields ...field.RelationField) *loveLoverRelationDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l loveLoverRelationDo) Preload(fields ...field.RelationField) *loveLoverRelationDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l loveLoverRelationDo) FirstOrInit() (*model.LoveLoverRelation, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoveLoverRelation), nil
	}
}

func (l loveLoverRelationDo) FirstOrCreate() (*model.LoveLoverRelation, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LoveLoverRelation), nil
	}
}

func (l loveLoverRelationDo) FindByPage(offset int, limit int) (result []*model.LoveLoverRelation, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l loveLoverRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l loveLoverRelationDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l loveLoverRelationDo) Delete(models ...*model.LoveLoverRelation) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *loveLoverRelationDo) withDO(do gen.Dao) *loveLoverRelationDo {
	l.DO = *do.(*gen.DO)
	return l
}