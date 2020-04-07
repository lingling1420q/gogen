package funcs

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	"github.com/silenceper/gogen/util"
)

//FuncMap 模板中允许使用的方法
var FuncMap = map[string]interface{}{
	"capitalizeFirst":   CapitalizeFirst,
	"lowerFirst":        LowerFirst,
	"camel":             Camel,
	"upperCamel":        UpperCamel,
	"underline":         Underline,
	"contains":          contains,
	"isEmpty":           IsEmpty,
	"isNotEmpty":        IsNotEmpty,
	"isNot":             IsNot,
	"isBuildIn":         IsBuildIn,
	"match":             match,
	"incr":              incr,
	"upper":             strings.ToUpper,
	"lower":             strings.ToLower,
	"trimSpace":         strings.TrimSpace,
	"toSortedMap":       ToSortedMap,
	"parseAPIArgument":  ParseAPIArgument,
	"parseAPIArguments": ParseAPIArguments,
	"inc":               Inc,

	"isArray":        isArray,
	"isBuiltIn":      isBuiltIn,
	"isBuiltInArray": isBuiltInArray,
	"hasPrefix":      strings.HasPrefix,
	"nullableType":   nullableType,
}

//APIArgument Inc
type APIArgument struct {
	Name     string
	Type     string
	Describe string
	Optional bool
	Default  interface{}
}

//Inc Inc
func Inc(v int) int {
	return v + 1
}

//ToSortedMap map to sort array
func ToSortedMap(m interface{}) []MapItem {
	slice := []MapItem{}
	if m != nil {
		for k, v := range m.(map[interface{}]interface{}) {
			slice = append(slice, MapItem{
				Key:   k.(string),
				Value: v,
			})
			sort.Sort(byKey(slice))
		}
	}
	return slice
}

//MapItem map k-v struct
type MapItem struct {
	Key   string
	Value interface{}
}

type byKey []MapItem

func (a byKey) Len() int           { return len(a) }
func (a byKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func nullableType(s string) string {
	if strings.HasPrefix(s, "[]") {
		return s
	}
	if isBuiltIn(s) {
		return s
	}
	return "*" + s
}

func isBuiltIn(t string) bool {
	if t == "int" || t == "int32" || t == "int64" || t == "string" || t == "bool" || t == "float32" || t == "float64" || t == "uint" || t == "uint32" || t == "uint64" {
		return true
	}
	return false
}

func isBuiltInArray(t string) bool {
	if !strings.HasPrefix(t, "[]") {
		return false
	}
	return isBuiltIn(t[:2])
}

func isArray(t string) bool {
	return strings.HasPrefix(t, "[]")
}

func match(s1, s2 string) bool {
	if s1 == "*" {
		return true
	}
	if strings.Contains(s1, s2) {
		return true
	}
	return false
}

func incr(i int) int {
	return i + 1
}

func contains(o interface{}, s string) bool {
	s1 := fmt.Sprintf("%s", o)
	return strings.Contains(s1, s)
}

//CapitalizeFirst 首字母大写
func CapitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	first := s[:1]
	return strings.ToUpper(first) + s[1:]
}

//LowerFirst 首字母小写
func LowerFirst(s string) string {
	if s == "" {
		return s
	}
	first := s[:1]
	return strings.ToLower(first) + s[1:]
}

//Camel 驼峰风格
func Camel(s string) string {
	bs := []byte{}
	captalize := false
	for i, c := range s {
		if i == 0 {
			bs = append(bs, byte(c))
			continue
		}
		if c == '_' {
			captalize = true
		} else {
			if captalize {
				c = unicode.ToUpper(c)
			}
			bs = append(bs, byte(c))
			captalize = false
		}
	}
	return LowerFirst(string(bs))
}

//UpperCamel 首字母大写的驼峰风格
func UpperCamel(s string) string {
	name := CapitalizeFirst(Camel(s))
	//符合golint规范的命名
	return util.LintName(name)
}

//Underline 下划线
func Underline(s string) string {
	bs := []byte{}
	for i, c := range s {
		if i == 0 {
			bs = append(bs, byte(c))
			continue
		}
		if unicode.IsLower(c) || c == '_' {
			bs = append(bs, byte(c))
		} else {
			bs = append(bs, '_', byte(unicode.ToLower(c)))
		}
	}
	return LowerFirst(string(bs))
}

//ParseAPIArgument ParseAPIArgument
func ParseAPIArgument(s string) *APIArgument {
	ss := strings.Split(s, ",")
	r := &APIArgument{}
	// nameType
	nameType := strings.Trim(ss[0], " ")
	nameTypeArray := strings.Split(nameType, ":")
	r.Name = strings.Trim(nameTypeArray[0], " ")
	r.Type = strings.Trim(nameTypeArray[1], " ")
	// Optional
	if strings.Contains(strings.ToLower(s), "optional") {
		r.Optional = true
	}
	// defaultValue
	if strings.Contains(strings.ToLower(s), "describe:") {
		descIdx := int32(strings.Index(strings.ToLower(s), "describe:"))
		descVal := strings.Split(s[descIdx+int32(len("describe:")):], ",")
		r.Describe = descVal[0]
	}
	// defaultValue
	if strings.Contains(strings.ToLower(s), "default:") {
		defaultIdx := int32(strings.Index(strings.ToLower(s), "default:"))
		defaultVal := strings.Split(s[defaultIdx+int32(len("default:")):], ",")
		r.Default = defaultVal[0]
	}
	if strings.ToLower(r.Name) == "default" {
		r.Name += "_"
	}
	// log.Debugf("%v", r)
	return r
}

//ParseAPIArguments 解析api 参数
func ParseAPIArguments(args []interface{}) []*APIArgument {
	r := []*APIArgument{}
	for _, arg := range args {
		s := arg.(string)
		r = append(r, ParseAPIArgument(s))
	}
	return r
}

// IsEmpty 判断是否为空
func IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}
	switch val.(type) {
	case string:
		return strings.TrimSpace(val.(string)) == ""
	case byte:
		return val == byte(0)
	case int:
		return val == int(0)
	case int32:
		return val == int32(0)
	case int64:
		return val == int64(0)
	case uint:
		return val == uint(0)
	case uint32:
		return val == uint32(0)
	case uint64:
		return val == uint64(0)
	case float32:
		return val == float32(0)
	case float64:
		return val == float64(0)

	default:
		return false
	}
}

//IsNotEmpty 判断不为空
func IsNotEmpty(val interface{}) bool {
	return !IsEmpty(val)
}

//IsBuildIn 是否内建类型
func IsBuildIn(t string) bool {
	return t == "int" || t == "int32" || t == "int64" || t == "string" || t == "float32" || t == "float64" || t == "uint" || t == "uint32" || t == "uint64"
}

//IsNot 判断不等
func IsNot(one, second string) bool {
	return strings.ToLower(one) != second
}
