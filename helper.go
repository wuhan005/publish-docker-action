package main

import (
	"regexp"
	"strings"
	"time"
)

//const (
//	RefTypeBranch = "branch"
//	RefTypeTag    = "tag"
//	RefTypePull   = "pull"
//)

//func resolveRef(github GitHub) (string, string) {
//	var typ, name string
//	refs := strings.SplitN(github.Ref, "/", 3)
//	if len(refs) == 3 {
//		switch refs[1] {
//		case "heads":
//			typ = RefTypeBranch
//		case "tags":
//			typ = RefTypeTag
//		case "pull":
//			typ = RefTypePull
//		}
//		name = refs[2]
//	}
//	return typ, name
//}

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

	t := strings.Replace(tagFormat, "%TIMESTAMP%", string(time.Now().Unix()), -1)

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
