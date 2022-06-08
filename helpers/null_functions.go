package helpers

import "database/sql"

func NullInt64(value int64) sql.NullInt64 {
	if value == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: value,
		Valid: true,
	}
}
func NullInt(value int32) sql.NullInt32 {
	if value == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: value,
		Valid: true,
	}
}
func NullString(value string) sql.NullString {
	if len(value) <= 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}
