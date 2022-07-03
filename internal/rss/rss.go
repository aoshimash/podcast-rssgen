package rss

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
	//"github.com/spf13/cobra"
)

var (
	ErrParsePubDateTimeHour = errors.New("failed to parse PubDateTimeHour")
	ErrParsePubDateTimeMin  = errors.New("failed to parse PubDateTimeMin")
	ErrParseFileName        = errors.New("failed to parse FileName")
)

func GenRSSString(dir string, baseURLStr string, channelTitle string, pubDateTimeStr string, thumbnailURLStr string) (string, error) {
	_, err := url.Parse(thumbnailURLStr)
	if err != nil {
		return "", err
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}

	podcastItems := []PodcastRSSItem{}
	for _, file := range files {
		podcastItem, err := newPodcastRSSItem(file.Name(), baseURLStr, pubDateTimeStr)
		if err != nil {
			return "", err
		}
		podcastItems = append(podcastItems, *podcastItem)
	}

	podcast := PodcastRSS{
		ChannelTitle:    channelTitle,
		ThumbnailURLStr: thumbnailURLStr,
		PodcastRSSItems: podcastItems,
	}

	// create RSS data
	tmpl := template.Must(template.New("call").Parse(tmplStr))
	buff := new(bytes.Buffer)
	fw := io.Writer(buff)
	err = tmpl.ExecuteTemplate(fw, "podcastrss", podcast)
	if err != nil {
		return "", err
	}

	return string(buff.Bytes()), nil
}

const tmplStr = `
{{- define "podcastrss" -}}
<?xml version="1.0" encoding="utf-8"?>
<rss xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" version="2.0">
  <channel>
    <title>{{.ChannelTitle}}</title>
    <itunes:image href="{{.ThumbnailURLStr }}"/>
{{- range $item := .PodcastRSSItems }}
    <item>
      <title>{{$item.Title}}</title>
      <enclosure url="{{$item.URL}}" type="{{$item.AudioType}}" />
      <pubDate>{{$item.FormatedPubDate}}</pubDate>
    </item>
{{- end }}
  </channel>
</rss>
{{ end -}}
`

type PodcastRSS struct {
	ChannelTitle    string
	ThumbnailURLStr string
	PodcastRSSItems []PodcastRSSItem
}

type PodcastRSSItem struct {
	Title           string
	URL             url.URL
	AudioType       string
	FormatedPubDate string
}

func newPodcastRSSItem(filename string, baseURLStr string, pubDateTimeStr string) (*PodcastRSSItem, error) {
	re := regexp.MustCompile(`.*_(\d{8}).+$`)
	rawDateStr := re.ReplaceAllString(filename, "$1")
	t, err := time.Parse("20060102", rawDateStr)
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	url.Path = path.Join(url.Path, filename)

	pubDateTimeHour, err := strconv.Atoi(pubDateTimeStr[:2])
	if err != nil {
		return nil, err
	}
	if pubDateTimeHour < 0 && pubDateTimeHour > 30 {
		return nil, ErrParsePubDateTimeHour
	}

	pubDateTimeMin, err := strconv.Atoi(pubDateTimeStr[2:])
	if err != nil {
		return nil, err
	}
	if pubDateTimeMin < 0 && pubDateTimeMin > 60 {
		return nil, ErrParsePubDateTimeMin
	}

	pubDate := time.Date(t.Year(), t.Month(), t.Day(), pubDateTimeHour, pubDateTimeMin, 0, 0, time.Local)
	formatedPubDate := pubDate.Format("Mon, 02 Jan 2006 15:04:04 -0700")

	pos := strings.LastIndex(filename, ".")
	fileExt := filename[pos+1:]
	audioType := "audio/" + fileExt

	ret := &PodcastRSSItem{
		Title:           t.Format("2006年01月02日"),
		URL:             *url,
		AudioType:       audioType,
		FormatedPubDate: formatedPubDate,
	}

	return ret, nil
}
