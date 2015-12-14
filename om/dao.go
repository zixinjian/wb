package om
import (
	"fmt"
)
type sTable struct {
	Name string
}
func Table(name string) *sTable{
	return &sTable{Name:name}
}
func (table *sTable)AddMap(valueMap map[string]interface{})(string, string){
	columns, marks, values := GetValuesSql(valueMap)
	query := fmt.Sprintf("INSERT INTO %s%s%s (%s%s%s) VALUES (%s)", Q, table.Name, Q, Q, columns, Q, marks)
	return DbAdd(query, values...)
}

//func (table *sTable)UpdateMap(valueMap map[string]interface{})(string, string){
//	sn, ok := valueMap[cc.Sn]
//	if !ok {
//		return st.SnNotFound, ""
//	}
//	var names []string
//	var values []interface{}
//	for k, v := range valueMap {
//		values = append(values, v)
//		names = append(names, k)
//	}
//	values = append(values, sn)
//	sep := fmt.Sprintf("%s = ?, %s", Q, Q)
//	setColumns := strings.Join(names, sep)
//	query := fmt.Sprintf("UPDATE %s%s%s SET %s%s%s = ? WHERE %s = ?", Q, table.Name, Q, Q, setColumns, Q, cc.Sn)
//	return DbUpdate(query, values...)
//}

//func (table *sTable)UpdateOrAddMap (valueMap map[string]interface{})(string, string){
//	return "", ""
//}
//func (table *sTable)IsExist(string)bool{
//	query := fmt.Sprintf("SELECT %s%s%s FROM %s%s%s WHERE %s%s%s = ?", Q, cc.Sn, Q, Q, table.Name, Q, Q, cc.Sn, Q)
//	_, list := DbQueryValues(query)
//	return len(list) > 0
//}