package main

import "fmt"

// Member 成员接口
type Member interface {
	Desc() string // 输出成员描述信息
}

// Teacher 老师
type Teacher struct {
	name    string // 名称
	subject string // 所教课程
}

// NewTeacher 根据姓名、课程创建老师
func NewTeacher(name string, subject string) *Teacher {
	return &Teacher{
		name:    name,
		subject: subject,
	}
}

func (t *Teacher) Desc() string {
	return fmt.Sprintf("%s班主任老师负责教%s", t.name, t.subject)
}

// Student 学生
type Student struct {
	name     string // 姓名
	sumScore int    // 考试总分数
}

// NewStudent 创建学生对象
func NewStudent(name string, sumScore int) *Student {
	return &Student{
		name:     name,
		sumScore: sumScore,
	}
}

func (t *Student) Desc() string {
	return fmt.Sprintf("%s同学考试总分为%d", t.name, t.sumScore)
}
