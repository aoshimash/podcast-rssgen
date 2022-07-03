# podcast-rssgen

podcast-rssgen is a CLI tool to create RSS Feed for podcast.

## Usage

### Golang

```
$ go run ./main.go [DIR] [BASE_URL] [CHANNEL_TITLE] [PUB_DATE_TIME] [THUMBNAIL_URL] [flags]
```

sample

```
$ go run ./main.go data/test "http://192.168.1.1:8888/audrey" "オードリーのオールナイトニッポン" 2700 "https://192.168.1.1:8888/audrey/thumbnail.png"
```

### Docker

```
$ docker run -it -v "$(pwd)/data:/data aoshimash/podcast-rssgen:latest [DIR] [BASE_URL] [CHANNEL_TITLE] [PUB_DATE_TIME] [THUMBNAIL_URL] [flags]
```

sample

```
$ docker run -it -v "$(pwd)/data:/data" aoshimash/podcast-rssgen:latest data/test "http://192.168.1.1:8001/audreyeeee" "オードリーのオールナイトニッポン" 2700 "https://192.168.1.1:8888/thumbnail.png"
```

## Sample Output

``` rss
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">
  <channel>
    <title>オードリーのオールナイトニッポン</title>
    <itunes:image href="https://192.168.1.1:8888/thumbnail.png"/>
    <item>
      <title>2021年01月01日</title>
      <enclosure url="http://192.168.1.1:8001/audreyeeee/audrey_20210101.aac" type="audio/aac" />
      <pubDate>Sat, 02 Jan 2021 03:00:00 +0000</pubDate>
    </item>
    <item>
      <title>2021年01月02日</title>
      <enclosure url="http://192.168.1.1:8001/audreyeeee/audrey_20210102.aac" type="audio/aac" />
      <pubDate>Sun, 03 Jan 2021 03:00:00 +0000</pubDate>
    </item>
    <item>
      <title>2021年01月03日</title>
      <enclosure url="http://192.168.1.1:8001/audreyeeee/audrey_20210103.aac" type="audio/aac" />
      <pubDate>Mon, 04 Jan 2021 03:00:00 +0000</pubDate>
    </item>
  </channel>
</rss>
```
