package resume

import "testing"

var pathZhiLian = "./test_documents/test_zhilian.txt"
var resumeZhiLian = NewZhiLianResume(pathZhiLian)

func TestZhiLianResume_JobWant(t *testing.T) {
	expected := "Java开发工程师"
	if resumeZhiLian.JobWant() != expected {
		t.Errorf("TestZhiLianResume_JobWant should be %s", expected)
	}
}

func TestZhiLianResume_Phone(t *testing.T) {
	expected := "15088111111"
	if resumeZhiLian.Phone() != expected {
		t.Errorf("TestZhiLianResume_Phone should be %s", expected)
	}
}

func TestZhiLianResume_Degree(t *testing.T) {
	expected := "本科"
	if resumeZhiLian.Degree() != expected {
		t.Errorf("TestZhiLianResume_Degree should be %s", expected)
	}
}

func TestZhiLianResume_Education(t *testing.T) {
	expected := "2014.03 - 2016.06 浙江理工大学 工商管理 本科 2008.09 - 2011.06 浙江树人大学 应用化学 大专 培训经历 2011.06 - 2011.10 Java软件开发"
	if resumeZhiLian.Education() != expected {
		t.Errorf("TestZhiLianResume_Education should be %s", expected)
	}
}

func TestZhiLianResume_Gender(t *testing.T) {
	expected := "男"
	if resumeZhiLian.Gender() != expected {
		t.Errorf("TestZhiLianResume_Gender should be %s", expected)
	}
}

func TestZhiLianResume_Name(t *testing.T) {
	expected := "billyct"
	if resumeZhiLian.Name() != expected {
		t.Errorf("TestZhiLianResume_Name should be %s", expected)
	}
}

func TestZhiLianResume_Age(t *testing.T) {
	expected := "30岁"
	if resumeZhiLian.Age() != expected {
		t.Errorf("TestZhiLianResume_Age should be %s", expected)
	}
}

func TestZhiLianResume_Birth(t *testing.T) {
	expected := "1988年7月"

	if resumeZhiLian.Birth() != expected {
		t.Errorf("TestZhiLianResume_Birth should be %s", expected)
	}
}

func TestZhiLianResume_WorkTime(t *testing.T) {
	expected := "6年工作经验"
	if resumeZhiLian.WorkTime() != expected {
		t.Errorf("TestZhiLianResume_WorkTime should be %s", expected)
	}
}

func TestZhiLianResume_Industry(t *testing.T) {
	expected := "计算机软件、互联网/电子商务、IT服务(系统/数据/维护)"
	if resumeZhiLian.Industry() != expected {
		t.Errorf("TestZhiLianResume_Industry should be %s", expected)
	}
}

func TestZhiLianResume_Residence(t *testing.T) {
	expected := "杭州 滨江区"
	if resumeZhiLian.Residence() != expected {
		t.Errorf("TestZhiLianResume_Industry should be %s", expected)
	}
}

func TestZhiLianResume_SalaryCurrent(t *testing.T) {
	expected := "10001-15000元/月"
	if resumeZhiLian.SalaryWant() != expected {
		t.Errorf("TestZhiLianResume_SalaryCurrent should be %s", expected)
	}
}

func TestZhiLianResume_SalaryWant(t *testing.T) {
	expected := "10001-15000元/月"
	if resumeZhiLian.SalaryWant() != expected {
		t.Errorf("TestZhiLianResume_SalaryWant should be %s", expected)
	}
}
