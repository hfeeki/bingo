bingo
=====

A web framework for golang. 

. mixed mango and go-json-rest
go-json-rest 提供了resource的管理、日志（处理前和处理后，错误后）、压缩、json缩进
mango的日志是通过middleware来实现的（默认是处理前）
mango.Routing返回一个middleware，其将url转化成对某个app的调用


Route Interface
=============
type Route struct {
    Method string
    PathExp string
    Action interface{}
}

type Router struct {
    Routes []Route  // routes storage 
}

func (self *Router) AddRoutes(routes ...Route) {
    
}



