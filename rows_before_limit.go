package clickhouse

import (
	"database/sql"
	"reflect"
)

//RowsBeforeLimit returns the  “rows_before_limit_at_least” (see https://clickhouse.tech/docs/en/interfaces/formats/#json)
func RowsBeforeLimit(r *sql.Rows) uint64 {
	rowsIPtr := reflect.ValueOf(r).Elem().FieldByName("rowsi")
	rowsI := reflect.Indirect(rowsIPtr).Elem()
	profileInfoPtr := reflect.Indirect(rowsI).FieldByName("profileInfo")
	return reflect.Indirect(profileInfoPtr).FieldByName("rowsBeforeLimit").Uint()
}
