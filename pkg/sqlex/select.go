package sqlex

type SelectController[T int64 | string] struct {
	Values  []T  `json:"values,optional"`
	Exclude bool `json:"exclude,optional"`
}

func (c *SelectController[T]) ClickHouseWhere(column string) (string, []T) {
	if len(c.Values) == 0 {
		return "", nil
	}

	var whereString string
	if c.Exclude {
		whereString = " not in ? "
	} else {
		whereString = " in ? "
	}

	return " " + column + " " + whereString, c.Values
}
