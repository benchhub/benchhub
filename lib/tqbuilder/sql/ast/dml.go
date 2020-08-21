package ast

// DataSource is a table, join result, sub query etc.
type DataSource interface {
	isDataSource()
}
