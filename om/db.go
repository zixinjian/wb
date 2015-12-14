package om
import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
	"strings"
	"wb/st"
)

const SqlErrUniqueConstraint = "UNIQUE constraint failed: "

func DbCount(sql string, values []interface{})int64{
	o := orm.NewOrm()
	var maps []orm.Params
	if _, err := o.Raw(sql, values...).Values(&maps); err == nil {
		if len(maps) <= 0 {
			beego.Error("GetCount error: len(maps) < 0")
			return -1
		}
		if total, ok := maps[0]["count"]; ok {
			if total64, err := strconv.ParseInt(total.(string), 10, 64); err != nil {
				beego.Error("GetCount Parseint error: ", err)
				return -1
			}else{
				return total64
			}
		}
	}else{
		beego.Error("GetCount error: ", err)
		return -1
	}

	beego.Error("GetCount unknown error: ", maps)
	return 0
}

func DbQueryValues(query string, values ...interface{})(string, []map[string]interface{}){
	beego.Debug(fmt.Sprintf("RawQueryMaps sql: %s, values: %s", query, values))
	var resultMaps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw(query, values).Values(&resultMaps)
	if err == nil {
		retList := make([]map[string]interface{}, len(resultMaps))
		//		fmt.Println("old", resultMaps)
		for idx, oldMap := range resultMaps {
			var retMap = make(map[string]interface{}, len(oldMap))
			for key, value := range oldMap {
				retMap[strings.ToLower(key)] = value
			}
			retList[idx] = retMap
		}
		return st.Success, retList
	}
	beego.Error("QueryItems Query error:", err.Error())
	return st.Failed, make([]map[string]interface{}, 0)
}

func DbDelete(query string, values ...interface{}) string{
	o := orm.NewOrm()
	if res, err := o.Raw(query, values...).Exec(); err == nil {
		if i, e := res.RowsAffected(); e == nil && i > 0 {
			return st.Success
		}else {
			beego.Error("RawDelete failed", err)
			return st.UnKnownFailed
		}
	} else {
		beego.Error("RawDelete error", err)
		return st.UnKnownFailed
	}
	return st.UnKnownFailed
}
func DbUpdate(sql string, values ...interface{})(string, string){
	o := orm.NewOrm()
	if res, err := o.Raw(sql, values...).Exec(); err == nil {
		if i, e := res.RowsAffected(); e == nil && i > 0 {
			return st.Success, ""
		}else{
			beego.Error(e, i)
			return parseSqlError(err)
		}
	} else {
		beego.Error("Update error", err.Error())
		return parseSqlError(err)
	}
}

func DbAdd(query string, values ...interface{})(string, string){
	beego.Debug("Add Item", query, values)
	o := orm.NewOrm()
	if res, err := o.Raw(query, values...).Exec(); err == nil {
		if i, e := res.LastInsertId(); e == nil && i > 0 {
			return st.Success, ""
		} else {
			beego.Error(e, i)
			return parseSqlError(err)
		}
	} else {
		beego.Error("Add error", err)
		return parseSqlError(err)
	}
}

func parseSqlError(err error)(string, string){
	errStr := err.Error()
	if strings.HasPrefix(errStr, SqlErrUniqueConstraint) {
		itemAndField := strings.TrimPrefix(errStr, SqlErrUniqueConstraint)
		lstStr := strings.Split(itemAndField, ".")
		if len(lstStr) < 2 {
			return st.DuplicatedValue, itemAndField
		}
		field := strings.TrimSpace(lstStr[1])
		return st.DuplicatedValue, field
	}
	beego.Error("ParseSqlError unknown error", errStr)
	return st.UnKnownFailed, errStr
}