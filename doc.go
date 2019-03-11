package resume

import (
	"io/ioutil"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"github.com/axgle/mahonia"
	"unicode/utf8"
)

func ReadDoc(path string) (string, error){
	s, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	reader := strings.NewReader(string(s))
	if (!utf8.Valid(s)) {
		// gb2312 转成 utf-8
		reader = strings.NewReader(mahonia.NewDecoder("gbk").ConvertString(string(s)))
	}

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	// 获取文档内容
	text := doc.Find("body").Text()

	// 替换 &nbsp;
	text = strings.Replace(text, "\u00A0", " ", -1)
	// 将所有 \s 换成一个空格
	r, err := regexp.Compile("\\s+")
	if err != nil {
		return "", err
	}
	text = r.ReplaceAllString(text, " ")
	text = strings.Trim(text, " ")

	return text, nil
}