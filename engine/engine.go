package engine

import (
	"filmspider/fetcher"
	"fmt"
	"log"
)

func Run(seed ...Request)  {
	var requests []Request

	for _, v := range seed  {
		requests = append(requests, v)
	}


	for len(requests) > 0  {
		request := requests[0]
		requests = requests[1:]

		fmt.Print(request.Url + "\n")
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("获取失败：%s 错误：%v", request.Url, err)
			continue
		}
		parserResult := request.ParserFunc(body)

		requests = append(requests, parserResult.Requests...)
	}


}
