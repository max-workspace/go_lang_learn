package people

import (
	"fmt"
	"time"
)

// people 设置为私有，强制使用构造方法获取people
type people struct {
	name       string
	age        int
	gender     int
	createTime time.Time
}

// init people struct
func init() {
	fmt.Println("init github.com/max_workspace/golang_learn/struct/people")
}

func NewPeople(name string, age, gender int) *people {
	return &people{name, age, gender, time.Now()}
}

// get people name
func (this *people) Name() string {
	return this.name
}

// get people age
func (this *people) Age() int {
	return this.age
}

// get people gender
func (this *people) Gender() int {
	return this.gender
}

// get people createTime
func (this *people) CreateTime() time.Time {
	return this.createTime
}

// get people gender description
func (this *people) GenderString() (genderString string) {
	if this.gender == 1 {
		genderString = "male"
	} else if this.gender == 0 {
		genderString = "female"
	}
	return genderString
}
