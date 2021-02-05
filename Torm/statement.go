package Torm

type Statement struct {
	clause *Clause
}

func NewStatement() *Statement {
	return &Statement{
		clause: NewClause(),
	}
}

func (this *Statement) SetTableName(tableName string) *Statement {
	this.clause.Tablename = tableName
	return this
}

// 新增数据API
func (this *Statement) InsertStruct(vars interface{}) *Statement {
	this.clause.InsertStruct(vars)
	return this
}

// 修改数据API
func (s *Statement) UpdateStruct(vars interface{}) *Statement {
	s.clause.UpdateStruct(vars)
	return s

}

// where条件
func (s *Statement) AndEqual(field string, value interface{}) *Statement {
	s.clause.AndEqual(field, value)
	return s

}

// where条件
func (s *Statement) OrEqual(field string, value interface{}) *Statement {
	s.clause.OrEqual(field, value)
	return s

}

// Select
func (s *Statement) Select(field ...string) *Statement {
	s.clause.SelectField(field...)
	return s

}
