package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func resolveAutoTag(inputs *Inputs) {
	if !inputs.AutoTag {
		return
	}

	tags := make([]string, 0)
	// Add `latest` version
	tags = append(tags, "latest")
	tags = append(tags, getFormatTag(inputs.TagFormat)...)

	inputs.Tags = tags
}

func getFormatTag(tagFormat string) []string{
	tags := make([]string, 0)

	t := tagFormat
	t = strings.Replace(t, "%TIMESTAMP%", strconv.Itoa(int(time.Now().Unix())), -1) 	// %TIMESTAMP% Timestamp
	t = strings.Replace(t, "%YYYY%", strconv.Itoa(time.Now().Year()), -1)           	// %YYYY% Year
	t = strings.Replace(t, "%MM%", strconv.Itoa(int(time.Now().Month())), -1)       	// %MM% Month
	t = strings.Replace(t, "%DD%", strconv.Itoa(time.Now().Day()), -1)              	// %DD% Day
	t = strings.Replace(t, "%H%", strconv.Itoa(time.Now().Hour()), -1)              	// %H% Hour
	t = strings.Replace(t, "%m%", strconv.Itoa(time.Now().Minute()), -1)              	// %m% Minute
	t = strings.Replace(t, "%s%", strconv.Itoa(time.Now().Second()), -1)              	// %s% Second
	
	tags = append(tags, t)
	return tags
}

func resolveSemanticVersionTag(name string) []string {
	re := regexp.MustCompile(`^v?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	matches := re.FindAllStringSubmatch(name, -1)
	if matches == nil {
		return []string{name}
	}
	subs := matches[0]
	suffixes := make([]string, 0)

	prerelease := subs[4]
	if prerelease != "" {
		suffixes = append(suffixes, "-", prerelease)
	}

	// metadata := subs[5]
	// if metadata != "" {
	// 	suffixes = append(suffixes, ".", metadata)
	// }

	tags := make([]string, 0)
	suffix := strings.Join(suffixes, "")
	for n := 2; n <= 4; n++ {
		vs := make([]string, n-1)
		copy(vs, subs[1:n])
		tags = append(tags, strings.Join([]string{strings.Join(vs, "."), suffix}, ""))
	}
	return tags
}
