package mysql

func WarpSqlValues(values []string) string {
	var ret string
	for i, v := range values {
		if i+1 == len(values) {
			ret += "'" + v + "'"
		} else {
			ret += "'" + v + "', "
		}
	}
	return ret
}

func WarpBracketedSqlValues(values []string) string {
	var ret string
	for i, v := range values {
		if i == 0 {
			ret += "("
		}
		if i+1 == len(values) {
			ret += "'" + v + "')"
		} else {
			ret += "'" + v + "', "
		}
	}
	return ret
}
