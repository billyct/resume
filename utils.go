package resume

import "strings"

// 删除空字符串 https://dabase.com/e/15006/
func FilterSlice(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func IsZhiLianResume(s string) bool {
	return strings.Contains(s, "智联招聘")
}

