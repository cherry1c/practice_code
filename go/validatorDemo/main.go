package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Phone  string `form:"phone" json:"phone" validate:"required"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

/*
required ：必填
email：验证字符串是email格式；例：“email”
url：这将验证字符串值包含有效的网址;例：“url”
max：字符串最大长度；例：“max=20”
min:字符串最小长度；例：“min=6”
excludesall:不能包含特殊字符；例：“excludesall=0x2C”//注意这里用十六进制表示。
len：字符长度必须等于n，或者数组、切片、map的len值为n，即包含的项目数；例：“len=6”
eq：数字等于n，或者或者数组、切片、map的len值为n，即包含的项目数；例：“eq=6”
ne：数字不等于n，或者或者数组、切片、map的len值不等于为n，即包含的项目数不为n，其和eq相反；例：“ne=6”
gt：数字大于n，或者或者数组、切片、map的len值大于n，即包含的项目数大于n；例：“gt=6”
gte：数字大于或等于n，或者或者数组、切片、map的len值大于或等于n，即包含的项目数大于或等于n；例：“gte=6”
lt：数字小于n，或者或者数组、切片、map的len值小于n，即包含的项目数小于n；例：“lt=6”
lte：数字小于或等于n，或者或者数组、切片、map的len值小于或等于n，即包含的项目数小于或等于n；例：“lte=6”
*/
func validatorParam(param interface{}) error {
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err) //Key: 'Users.Passwd' Error:Field validation for 'Passwd' failed on the 'min' tag
		}
		fmt.Printf("%#v\n", param)
		return nil
	}
	return nil
}

func main() {
	users := &Users{
		Phone:  "1326654487",
		Passwd: "12345",
		Code:   "12345",
	}
	if err := validatorParam(users); err != nil {
		return
	}
	return
}
