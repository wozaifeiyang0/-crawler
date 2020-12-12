package main

import "crawler/parse"

func main() {

	parse.Region("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/13/01/130101.html  " , "./test/towns.csv", ".towntable .towntr", 9, 12)
}
