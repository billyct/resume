package resume

import (
	"regexp"
	"strings"
)

// 解析智联招聘的文档
type ZhiLianResume struct {
	content string
}

func NewZhiLianResume(path string) *ZhiLianResume {
	content, _ := ReadDoc(path)
	return &ZhiLianResume{content: content}
}

// 职位
func (r ZhiLianResume) JobWant() string {
	reg, _ := regexp.Compile("期望从事职业： (.+?) ")
	return reg.FindStringSubmatch(r.content)[1]
}

// 姓名
func (r ZhiLianResume) Name() string {
	reg, _ := regexp.Compile("ID：\\)\\w+ (.+?) 手机")
	return reg.FindStringSubmatch(r.content)[1]
}

// 手机
func (r ZhiLianResume) Phone() string {
	reg, _ := regexp.Compile("\\d{11}")
	return reg.FindString(r.content)
}

// 学历
func (r ZhiLianResume) Degree() string {
	degrees := []string {
		"专科",
		"本科",
		"硕士",
		"博士",
		"研究生",
	}

	for _, v := range degrees {
		if strings.Contains(r.content, v) {
			return v
		}
	}

	return ""
}

// 教育经历
func (r ZhiLianResume) Education() string {
	reg, _ := regexp.Compile("教育经历 (.+) 培训机构")
	return reg.FindStringSubmatch(r.content)[1]
}


// 性别
func (r ZhiLianResume) Gender() string {
	if strings.Contains(r.content, "女") {
		return "女"
	}

	return "男"
}

// 年龄
func (r ZhiLianResume) Age() string {
	reg, _ := regexp.Compile("(\\d+)岁")
	return reg.FindString(r.content)
}

// 出生年月
func (r ZhiLianResume) Birth() string {
	reg, _ := regexp.Compile("\\((\\d{4}年\\d*月)\\)")
	return reg.FindStringSubmatch(r.content)[1]
}

func (r ZhiLianResume) JobCurrent() string {
	return ""
}

func (r ZhiLianResume) WorkTime() string {
	reg, _ := regexp.Compile("\\d*年工作经验")
	return reg.FindString(r.content)
}

func (r ZhiLianResume) Industry() string {
	reg, _ := regexp.Compile("期望从事行业： (.+?) ")
	return reg.FindStringSubmatch(r.content)[1]
}

func (r ZhiLianResume) CompanyCurrent() string {
	return ""
}

func (r ZhiLianResume) Residence() string {
	reg, _ := regexp.Compile("现居住地：(.+?) \\|")
	return reg.FindStringSubmatch(r.content)[1]
}

func (r ZhiLianResume) SalaryCurrent() string {
	reg, _ := regexp.Compile("\\| (.+?元/月) ")
	return reg.FindStringSubmatch(r.content)[1]
}

func (r ZhiLianResume) SalaryWant() string {
	reg, _ := regexp.Compile("期望月薪： (.+?) ")
	return reg.FindStringSubmatch(r.content)[1]
}

