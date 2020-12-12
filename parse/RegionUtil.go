package parse

import (
	"crawler/csvutil"
	http2 "crawler/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	thy "github.com/wozaifeiyang0/thylog"
	"net/http"
	"strconv"
	"time"
)

func region2(url string, csvName string, selector string, start int, end int) [][]string {

	// 防止网址认为是恶意进攻加延时 2秒钟
	time.Sleep(time.Duration(100) * time.Millisecond)

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)
	// 设置请求投
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	resp, err := (&http.Client{Timeout: time.Second * 3}).Do(req)
	if err != nil {
		thy.Error.Println(url)
		thy.Info.Println(err)
		// 防止网址认为是恶意进攻加延时 2秒钟
		time.Sleep(time.Duration(2000) * time.Millisecond)
		resp, err = (&http.Client{}).Do(req)
	}

	if resp.StatusCode != 200 {
		thy.Error.Printf("请求    %s    异常\n", url)
		// 防止网址认为是恶意进攻加延时 2秒钟
		time.Sleep(time.Duration(2000) * time.Millisecond)
		resp, err = (&http.Client{}).Do(req)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		thy.Error.Printf("错误信息：%s", err.Error())
	}

	var rowsArray [][]string
	index := 1
	doc.Find(selector).Each(func(i int, region *goquery.Selection) {

		code := region.Text()[:start]
		name := enc.ConvertString(region.Text()[end:])

		if code != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), code), name)
			rowsArray = append(rowsArray, row)
			index++
			thy.Info.Printf("行政区划信息，编码：%s，名称：%s \n", code, name)
		}

	})
	csvutil.Write(csvName, rowsArray)

	return rowsArray
}

// 行政区划通用请求方式
func Region(url string, csvName string, selector string, start int, end int) [][]string {

	// 防止网址认为是恶意进攻加延时 2秒钟
	time.Sleep(time.Duration(50) * time.Millisecond)

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")

	resp, _ := http2.Get(url, 3, 5, 5)

	var rowsArray [][]string
	if resp != nil {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			thy.Error.Printf("错误信息：%s", err.Error())
			return rowsArray
		}
		index := 1
		doc.Find(selector).Each(func(i int, region *goquery.Selection) {

			code := region.Text()[:start]
			name := enc.ConvertString(region.Text()[end:])

			if code != "" {
				var row []string
				row = append(append(append(row, strconv.Itoa(index)), code), name)
				rowsArray = append(rowsArray, row)
				index++
				thy.Info.Printf("行政区划信息，编码：%s，名称：%s \n", code, name)
			}

		})

		csvutil.Write(csvName, rowsArray)
	}
	return rowsArray
}

// 请求省数据接口
func Province(url string, csvName string) [][]string {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)
	// 设置请求投
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		thy.Error.Println(url)
		thy.Info.Println(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		thy.Info.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var rowsArray [][]string
	index := 1
	doc.Find(".provincetable .provincetr a").Each(func(i int, s *goquery.Selection) {
		// name
		provinceHtml, _ := s.Attr("href")
		provinceCode := string([]rune(provinceHtml)[:2])
		provinceName := enc.ConvertString(s.Text())

		if provinceCode != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), provinceCode), provinceName)
			rowsArray = append(rowsArray, row)
			index++
			thy.Info.Printf("行政区划信息，编码：%s，名称：%s \n", provinceCode, provinceName)
		}

	})

	csvutil.Write(csvName, rowsArray)

	return rowsArray
}

func getProvince(url string, csvName string) [][]string {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(url, csvName)
		fmt.Println(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var rowsArray [][]string
	index := 1
	doc.Find(".provincetable .provincetr a").Each(func(i int, s *goquery.Selection) {
		// name
		provinceHtml, _ := s.Attr("href")
		provinceCode := string([]rune(provinceHtml)[:2])
		provinceName := enc.ConvertString(s.Text())

		if provinceCode != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), provinceCode), provinceName)
			rowsArray = append(rowsArray, row)
			index++
		}

	})

	csvutil.Write(csvName, rowsArray)

	return rowsArray
}

func getCity(url string, csvName string) [][]string {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(url, csvName)
		fmt.Println(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var rowsArray [][]string
	index := 1
	doc.Find(".citytable .citytr").Each(func(i int, city *goquery.Selection) {

		cityCode := city.Text()[:4]
		cityName := enc.ConvertString(city.Text()[12:])

		if cityCode != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), cityCode), cityName)
			rowsArray = append(rowsArray, row)
			index++
		}

	})

	csvutil.Write(csvName, rowsArray)

	return rowsArray
}

func getCounty(url string, csvName string) [][]string {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(url, csvName)
		fmt.Println(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var rowsArray [][]string
	index := 1
	doc.Find(".countytable .countytr").Each(func(i int, county *goquery.Selection) {

		countyCode := county.Text()[:6]
		countyName := enc.ConvertString(county.Text()[12:])

		if countyCode != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), countyCode), countyName)
			rowsArray = append(rowsArray, row)
			index++
		}

	})

	csvutil.Write(csvName, rowsArray)

	return rowsArray
}

func getTown(url string, csvName string) [][]string {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(url, csvName)
		fmt.Println(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var rowsArray [][]string
	index := 1
	doc.Find(".towntable .towntr").Each(func(i int, town *goquery.Selection) {

		townCode := town.Text()[:9]
		townName := enc.ConvertString(town.Text()[12:])

		if townCode != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), townCode), townName)
			rowsArray = append(rowsArray, row)
			index++
		}

	})

	csvutil.Write(csvName, rowsArray)

	return rowsArray
}

func getVillage(url string, csvName string) [][]string {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println(url, csvName)
		fmt.Println(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	var rowsArray [][]string
	index := 1
	doc.Find(".villagetable .villagetr").Each(func(i int, village *goquery.Selection) {

		villageCode := village.Text()[:12]
		villageName := enc.ConvertString(village.Text()[15:])

		if villageCode != "" {
			var row []string
			row = append(append(append(row, strconv.Itoa(index)), villageCode), villageName)
			rowsArray = append(rowsArray, row)
			index++
		}

	})

	csvutil.Write(csvName, rowsArray)

	return rowsArray
}
