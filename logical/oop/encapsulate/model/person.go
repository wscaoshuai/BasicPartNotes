package model
import(
	"fmt"
)

type person struct{
	Name string 
	age int //其他包无法引用
	sal float64

}

func NewPerson(name string)*person  {
	return &person{
		Name : name ,
	}
}

func (p *person) SetAge(age int)  {
	if age > 0 && age < 150{
		p.age = age 
	}else{
		fmt.Println("年龄不对")
		//此处也可以给一个默认值
	}
}

func (p *person) GetAge()int  {
	return p.age

}


func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水不对")
		//此处也可以给一个默认值
	}
}


func (p *person) GetSal()float64  {
	return p.sal
}