package om
import (
	"fmt"
	"strings"
)


func GetValuesSql(vMap map[string]interface{})(columns string, marks string, values []interface{}) {
	columnList := make([]string, 0)
	markList := make([]string, 0)
	values = make([]interface{}, 0)

	for k, v := range vMap {
		columnList = append(columnList, k)
		markList = append(markList, "?")
		values = append(values, v)
	}

	sep := fmt.Sprintf("%s, %s", Q, Q)
	marks = strings.Join(markList, Mark)
	columns = strings.Join(columnList, sep)
	return
}