package apiware

import (
	"reflect"
	"sync"
)

type (
	// ParamsAPI 是 一个web api的参数模型
	ParamsAPI struct {
		name   							string
		params 							[]*Param
		structType 						reflect.Type        //创建一个新的非指针结构体
		rawStructPointer 				interface{}			//原始结构体的指针
		defaultValues					[]byte				//原始结构体指针的集合
		paramNameMapper 				ParamNameMapper		//根据结构体字段名创建一个参数
		bodyDecoder 					BodyDecoder			//从请求体中解码出参数
		maxMemory 						int64				//when request Content-Type is multipart/form-data, the max memory for body.
	}
	// Schema ParamsAPI的收集器
	Schema struct {
		lib map[string]*ParamsAPI
		sync.RWMutex
	}

	// ParamNameMapper 从结构体参数名映射参数名
	ParamNameMapper func(fieldName string) (paramName string)

	// BodyDecoder 从请求体中解码参数
	BodyDecoder func(dest reflect.Value, body []byte) error
)

// 默认ParamsAPI的收集器
var (
	defaultSchema = &Schema{
		lib: map[string]*ParamsAPI{},
	}
)

func NewParamsAPI(structPointer interface{},paramNameMapper ParamNameMapper,bodyDecoder BodyDecoder,useDefaultValues bool)(*ParamsAPI, error,) {
	name := reflect.TypeOf(structPointer).String() //获取结构体指针系统类型，而非具体值
	v := reflect.ValueOf(structPointer)			   //获取结构体指针的reflect.Value
	if v.Kind() != reflect.Ptr {
		return nil, NewError(name, "*", "绑定对象必须是结构体指针")
	}
	v = reflect.Indirect(v) //获取v所指向的值
	if v.Kind() != reflect.Struct {
		return nil, NewError(name, "*", "绑定对象必须是结构体指针")
	}

	var paramsAPI = &ParamsAPI{
		name:             name,
		params:           []*Param{},
		structType:       v.Type(),
		rawStructPointer: structPointer,
	}


}


