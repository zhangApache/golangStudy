package main

import "log"

type Human struct {
	name string
	age int
	wegit int
}

type Skills []string
type Student struct {
	Human      //匿名字段，默认Student包含Human中所有字段
	Skills
	int
	speciality string
}

func change(stu *Student) *Student {
	stu.speciality = "AI"
	return stu
}
func main() {
	mark := Student{Human:Human{"Mark", 24, 60}, speciality:"Computer Science"}
	log.Println("His name is ", mark.name)
	log.Println("His age is ", mark.age)
	log.Println("His wegit is ", mark.wegit)
	log.Println("His speciality is ", mark.speciality)

	mark.Skills = []string{"anatomy"}
	stu := change(&mark)
	log.Println("Mark changed his speciality")
	log.Println("His speciality is ", stu.speciality)
}
