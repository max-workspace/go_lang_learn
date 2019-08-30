package people

import (
	"fmt"
	"time"
)

// people 设置为私有，强制使用构造方法获取people
type people struct {
	Name       string
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

// get people age
func (self *people) Age() int {
	return self.age
}

// get people gender
func (self *people) Gender() int {
	return self.gender
}

// get people createTime
func (self *people) CreateTime() time.Time {
	return self.createTime
}

// get people gender description
func (self *people) GenderString() (genderString string) {
	if self.gender == 1 {
		genderString = "male"
	} else if self.gender == 0 {
		genderString = "female"
	}
	return genderString
}
