
package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email string
}
/*
Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象，这和Java或者C++中的this类似，
如果要访问当前对象的字段通过{{.FieldName}}，但是需要注意一点：这个字段必须是导出的(字段首字母必须是大写的)，
否则在渲染的时候就会报错
*/
func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}!{{.email}}")
	p := Person{UserName: "baijing.biz"}
	t.Execute(os.Stdout, p)
}