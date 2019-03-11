package resume

type Resume interface {
	JobWant() string // 职位
	Name() string
	Gender() string
	Phone() string
	Birth() string
	Age() string
	Degree() string // 学历
	Education() string // 教育经历
	JobCurrent() string // 职务
	WorkTime() string
	Industry() string
	CompanyCurrent() string
	Residence() string // 居住地
	SalaryCurrent() string
	SalaryWant() string // 期望薪资
}
