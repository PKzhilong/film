package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items	[]interface{}
}

type Item struct {
	Url string
	Name string
	Type int
}

func NilParaser([]byte) ParseResult {
	return ParseResult{}
}