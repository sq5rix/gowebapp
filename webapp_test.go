package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {

	orig := "https://www.washingtonpost.com/national/former-jfk-supervisor-admits-to-taking-foreign-bribes/2019/06/21/260b930a-9448-11e9-956a-88c291ab5c38_story.html"
	re := regexp.MustCompile(`([a-z0-9]+-)+[a-z0-9]+`)
	s := re.FindString(orig)
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.Title(s)
	re = regexp.MustCompile(`\b\w{1,3}\b`)
	s = re.ReplaceAllStringFunc(s, strings.ToUpper)
	t.Error(orig)
	t.Error(s)
}
