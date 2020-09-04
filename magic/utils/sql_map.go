package utils

import "database/sql"

// SQLMap x
func SQLMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	rt := make([]map[string]interface{}, 0)
	columns, _ := rows.Columns()
	value := make([]interface{}, len(columns))
	valuePrt := make([]interface{}, len(columns))
	for rows.Next() {
		for i := range columns {
			valuePrt[i] = &value[i]
		}
		err := rows.Scan(valuePrt...)
		if err != nil {
			return nil, err
		}
		mp := map[string]interface{}{}
		for i, e := range columns {
			var v interface{}
			if co, ok := value[i].([]byte); ok {
				v = string(co)
			} else {
				v = value[i]
			}
			mp[e] = v
		}
		rt = append(rt, mp)
	}
	return rt, nil
}
