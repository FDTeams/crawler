package engine

// 每个请求的url和解析函数
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

// 存放每个解析后存放的请求和数据
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}
