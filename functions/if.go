package functions

func If(condition bool, success interface{}, elseSuccess interface{}) interface{} {
	if condition {
		return success
	}
	return elseSuccess
}

func IfPuntero(condition *string) string {
	if condition != nil {
		return *condition
	}

	return "-"
}
