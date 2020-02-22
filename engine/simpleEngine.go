package engine

import (
	"filmspider/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct {

}

func (s *SimpleEngine) Run(seed ...Request)  {
	var requests []Request

	for _, v := range seed  {
		requests = append(requests, v)
	}


	for len(requests) > 0  {
		request := requests[0]
		requests = requests[1:]

		parserResult, err := Worker(request)
		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Requests...)
	}


}

func Worker(request Request) (ParseResult, error)  {
	fmt.Print(request.Url + "\n")
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("获取失败：%s 错误：%v", request.Url, err)
		return ParseResult{}, err
	}
	parserResult := request.ParserFunc(body)
	return parserResult, nil
}
