package ut
import (
	"time"
	"fmt"
	"math"
	"strings"
	"text/template"
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

func Int16ToBytes(i int16)[]byte{
	buff :=  bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, i)
	return buff.Bytes()
}
func BytesToInt16(b []byte) int16 {
	buff := bytes.NewBuffer(b)
	var r int16
	binary.Read(buff, binary.BigEndian, &r)
	return r
}
func BytesToUint16(b []byte) uint16 {
	buff := bytes.NewBuffer(b)
	var r uint16
	binary.Read(buff, binary.BigEndian, &r)
	return r
}

func BytesToHex(bs []byte) string{
	dst := make([]byte, hex.EncodedLen(len(bs)))
	hex.Encode(dst, bs)
	return string(dst)
}

var gId uint32
var gOldTime time.Time

func TUId() string {
	now := time.Now()
	if gOldTime.After(now) {
		now = gOldTime
	}
	if gId > 99 {
		gId = 0
		now.Add(time.Second)
	}
	gOldTime = now
	ret := fmt.Sprintf("%s%02d", now.Format("20060102150405"), gId)
	gId = gId + 1
	return ret
}

func init() {
	gId = 0
	gOldTime = time.Now()
}

func GetStringValue(vMap map[string]interface{}, key string)string{
	if value, ok := vMap[key];ok {
		if v, vok := value.(string); vok {
			return v
		}
	}
	return ""
}
func GetCreateTime()string{
	return time.Now().Format("20060102150405")
}
func GetToday() (totay string){
	t := time.Now()
	return t.Format("2006.01.02")
}

func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func IsNullStr(str interface{}) bool {
	return strings.EqualFold(str.(string), "")
}


func TemplateFormat(temp string, data interface{}) string{
	t := template.New("temp")
	t, _ = t.Parse(temp)
	b := new(bytes.Buffer)
	t.Execute(b, data)
	return b.String()
}