package main

import (
	"ds/linkedList"

	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func main() {
	println(Hello())
	test(linkedList.New())
}

func test(list linkedList.IList) {

}
