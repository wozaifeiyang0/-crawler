package main

import (
	"crawler/fileutil"
	"crawler/parse"
	thy "github.com/wozaifeiyang0/thylog"
)

func main() {

	httpUrl := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2019/"
	path := "./data"

	//创建文件夹目录
	path = fileutil.CreateDateDir(path)

	thy.SetLogFile(path, "error.log")
	provinceName := path + "/province.csv"
	provinces := parse.Province(httpUrl+"index.html", provinceName)

	// 查询省数据
	for _, province := range provinces {

		// 查询市级数据
		cityName := path + "/city.csv"
		citys := parse.Region(httpUrl+province[1]+".html", cityName, ".citytable .citytr", 4, 12)

		// 查询县级数据
		countyCsvName := path + "/county.csv"
		for _, city := range citys {

			countys := parse.Region(httpUrl+province[1]+"/"+city[1]+".html", countyCsvName, ".countytable .countytr", 6, 12)

			// 查询镇级数据
			townCsvName := path + "/town.csv"
			for _, county := range countys {

				towns := parse.Region(httpUrl+province[1]+"/"+city[1][2:]+"/"+county[1]+".html", townCsvName, ".towntable .towntr", 9, 12)

				// 查询村级数据
				villageCsvName := path + "/village.csv"
				for _, town := range towns {
					parse.Region(httpUrl+province[1]+"/"+city[1][2:]+"/"+county[1][4:]+"/"+town[1]+".html", villageCsvName, ".villagetable .villagetr", 12, 15)
				}
			}
		}

	}

	thy.Info.Println("下载数据完成")

}
