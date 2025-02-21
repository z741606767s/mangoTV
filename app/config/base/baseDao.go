package base

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseDao struct {
	db *gorm.DB
}

func NewBaseDao(db *gorm.DB) *BaseDao {
	return &BaseDao{
		db: db,
	}
}

// AutoDb 根据传入的参数自动确定使用哪个数据库实例（或默认实例）
func (d BaseDao) AutoDb(args ...interface{}) (rdb *gorm.DB) {

	if len(args) > 0 {
		if db, ok := args[0].(*gorm.DB); ok {
			return db
		}
	}
	return d.db
}

// AutoArgs 除了选择数据库实例外，还会自动处理参数中附带的 SQL 语句片段
func (d BaseDao) AutoArgs(args ...interface{}) (rdb *gorm.DB) {
	if len(args) > 0 {
		if db, ok := args[len(args)-1].(*gorm.DB); ok {
			rdb = db
			if args = args[:len(args)-1]; len(args) > 0 {
				var expr []clause.Expression
				for _, v := range args {
					if vv, okk := v.(clause.Expression); okk {
						expr = append(expr, vv)
					}
				}
				rdb = rdb.Clauses(expr...)
			}
		}
	}
	if rdb == nil {
		rdb = d.db
		if len(args) > 0 {
			var expr []clause.Expression
			for _, v := range args {
				if vv, okk := v.(clause.Expression); okk {
					expr = append(expr, vv)
				}
			}
			rdb = rdb.Clauses(expr...)
		}
	}

	return
}

// ClassifyParameters 对传入的参数进行分类，将不同类型的参数分别归类，便于后续使用
func (d BaseDao) ClassifyParameters(args ...interface{}) (expressions []clause.Expression, columns []clause.Column, scopes []func(db *gorm.DB) *gorm.DB) {

	for _, v := range args {
		if vv, ok := v.(clause.Expression); ok {
			expressions = append(expressions, vv)
		} else if vv1, ok1 := v.(func(db *gorm.DB) *gorm.DB); ok1 {
			scopes = append(scopes, vv1)
		} else if vv2, ok2 := v.(string); ok2 {
			columns = append(columns, clause.Column{Name: vv2, Raw: true})
		} else if vv3, ok3 := v.([]string); ok3 {
			for _, vvv := range vv3 {
				columns = append(columns, clause.Column{Name: vvv, Raw: true})
			}
		} else if vv4, ok4 := v.(clause.Column); ok4 {
			columns = append(columns, vv4)
		}
	}
	return
}
