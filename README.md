# Golang-GeojsonHelper
[中文文档](#是什么)
## What is this?

Helps you locate whether a point(latitude and longitude) is in the GeoJson map

Find location info by latitude and longitude 

Compare the query efficiency of different geojson files

## What you need
[GeoJson](https://geojson.org/) Map [(Wiki)](https://en.wikipedia.org/wiki/GeoJSON)


#### How to get GeoJson File?

You can find them for certain countries or regions on certain websites. For example, you can find geojson file of China on [Aliyun](http://datav.aliyun.com/tools/atlas/)

Another idea is try to find it on GitHub or Google

Otherwise, you need to do:

1. First download a [shape file](http://en.wikipedia.org/wiki/Shapefile), on websites like [gadm](http://www.gadm.org/country) or [naturalearth](http://www.naturalearthdata.com/downloads/)
2. Generate geojson file by shape file on website [mapshaper](http://www.mapshaper.org/), import your shape file and you can use `simplify` to control the size of your geojson file, finally export it.

## How to use

`go get github.com/linvon/golang-geohelper`

Prepare an geojson file, and get the `key` like "name"

``` json
{
  "type": "Feature",
  "geometry": {
    "type": "Point",
    "coordinates": [125.6, 10.1]
  },
  "properties": {
    "name": "Dki Jakarta"
  }
}
```

Code and Run

``` go
m, err := geohelper.NewGeoMap("highgeo.json", "name") // file, key
if err != nil {
  log.Fatalln(err)
}

provinceName := m.FindLoc(-6.196893, 106.830407)
fmt.Println(provinceName)
```





## 是什么?

帮助你判断一个用经纬度表示的点是否在 geojson 地图中

通过经纬度反查地理位置信息

比较不同的 geojson 文件查询效率

## 需要什么

[GeoJson](https://geojson.org/) Map  [(Wiki)](https://en.wikipedia.org/wiki/GeoJSON)


#### 如何获取 GeoJson 文件?

你可以在某些网站上找到某些国家或地区的 geojson 文件。比如，你可以在[阿里云](http://datav.aliyun.com/tools/atlas/)的工具站上找到中国的 geojson 文件

另一个方法是尝试在 GitHub 或者 Google 上面找到相关文件

其他情况下，你需要这么做:

1. 首先下载一个[形状文件](http://en.wikipedia.org/wiki/Shapefile), 可以从这些网站上下载： [gadm](http://www.gadm.org/country) 或 [naturalearth](http://www.naturalearthdata.com/downloads/)
2. 在 [mapshaper](http://www.mapshaper.org/) 上通过形状文件生成 geojson 文件, 导入形状文件后你可以使用简化操作控制你的 geojson 文件大小，最后导出文件

## 如何使用

`go get github.com/linvon/golang-geohelper`

准备好一个 geojson 文件，并找到其中表示名称的 `key`比如 "name"

``` json
{
  "type": "Feature",
  "geometry": {
    "type": "Point",
    "coordinates": [125.6, 10.1]
  },
  "properties": {
    "name": "Dki Jakarta"
  }
}
```

编码并运行

``` go
m, err := geohelper.NewGeoMap("highgeo.json", "name") // file, key
if err != nil {
  log.Fatalln(err)
}

provinceName := m.FindLoc(-6.196893, 106.830407)
fmt.Println(provinceName)
```






