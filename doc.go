package resume

import (
	"io/ioutil"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ReadDoc(path string) (string, error){
	s, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(s)))
	if err != nil {
		return "", err
	}

	// 获取文档内容
	text := doc.Find("body").Text()

	// 替换 &nbsp;
	text = strings.Replace(text, "\u00A0", "", -1)
	// 将所有 \s 换成一个空格
	r, err := regexp.Compile("\\s+")
	if err != nil {
		return "", err
	}
	text = r.ReplaceAllString(text, " ")
	text = strings.Trim(text, " ")

	return text, nil
}