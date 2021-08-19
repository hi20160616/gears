package gears

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// HttpGetBody request url to get Html Body, if get error occur it'll try n times.
func HttpGetBody(url string, n int) (string, error) {
	raw, err := http.Get(url)
	for err != nil && n > 0 {
		raw, err = http.Get(url)
		time.Sleep(time.Minute * 1)
		n--
	}
	if err != nil {
		return "", err
	}
	rawBody, err := ioutil.ReadAll(raw.Body)
	defer raw.Body.Close()
	if err != nil {
		return "", err
	}
	if raw.StatusCode != 200 {
		return "", err
	}
	return string(rawBody), nil
}

// HttpGetTitleViaTwitterJS get post title via twitter share javascripts' json data
func HttpGetTitleViaTwitterJS(rawBody string) string {
	var a = regexp.MustCompile(`(?m)<meta name="twitter:title" content="(?P<title>.*?)"`)
	rt := a.FindStringSubmatch(rawBody)
	if rt != nil {
		return rt[1]
	} else {
		return ""
	}
}

// HttpGetSiteViaTwitterJS get post site via twitter share javascripts' json data
func HttpGetSiteViaTwitterJS(rawBody string) string {
	var a = regexp.MustCompile(`(?m)<meta name="twitter:site" content="(?P<site>.*?)"`)
	rt := a.FindStringSubmatch(rawBody)
	if rt != nil {
		return rt[1]
	} else {
		return ""
	}
}

func HttpGetDateViaMeta(rawBody string) string {
	var a = regexp.MustCompile(`(?m)<meta name="parsely-pub-date" content="(?P<date>.*?)".*?/>`)
	rt := a.FindStringSubmatch(rawBody)
	if rt != nil {
		return rt[1]
	} else {
		return ""
	}
}

func HttpGetDateByHeader(rawBody string) string {
	var a = regexp.MustCompile(`(?m)"dateModified":\s*?"(?P<Datetime>.*?)",`)
	rt := a.FindStringSubmatch(rawBody)
	if rt != nil {
		return rt[1]
	} else {
		return ""
	}
}

func HttpGetDomain(url string) string {
	var a = regexp.MustCompile(`(?m)https?://(\w+.\w+.\w+)/`)
	rt := a.FindStringSubmatch(url)
	if rt != nil {
		return rt[1]
	} else {
		return ""
	}
}
