package main

import (
	"encoding/xml"
	"time"
)

var fileXML = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>https://www.goinggo.net/</link>
	<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
    <item>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an object oriented language.</description>
        <link>https://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>
`

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"puDate"`
	Item        []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

func Read(doc string) ([]Item, error) {
	time.Sleep(time.Millisecond) // 模拟读文件的阻塞事件
	var d Document
	if err := xml.Unmarshal([]byte(fileXML), &d); err != nil {
		return nil, err
	}
	return d.Channel.Item, nil
}
