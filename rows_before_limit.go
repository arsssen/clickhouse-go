package clickhouse

import (
	"database/sql"
	"reflect"
)

//RowsBeforeLimit returns the  “rows_before_limit_at_least” (see https://clickhouse.tech/docs/en/interfaces/formats/#json)
//should be called after retrieving results with rows.Next()
func RowsBeforeLimit(r *sql.Rows) uint64 {
	rowsIPtr := reflect.ValueOf(r).Elem().FieldByName("rowsi")
	rowsI := reflect.Indirect(rowsIPtr).Elem()
	profileInfoPtr := reflect.Indirect(rowsI).FieldByName("profileInfo")
	if profileInfoPtr.IsNil() {
		return 0
	}
	return reflect.Indirect(profileInfoPtr).FieldByName("rowsBeforeLimit").Uint()
}
