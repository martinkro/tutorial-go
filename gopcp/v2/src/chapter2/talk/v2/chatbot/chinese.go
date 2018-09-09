package chatbot

import(
	"fmt"
	"strings"
)

type simpleCN struct{
	name string
	talk Talk
}

func NewSimpleCN(name string,talk Talk)Chatbot{
	return &simpleCN{
		name:name,
		talk:talk
	}
}

func (robot *simpleCN)Name()string{
	return robot.name
}

func (robot* simpleCN)Begin()(string,error){
	return "请输入你的名字：",nil
}

func (robot* simpleCN)Hello(userName string)string{
	userName = string.TrimSpace(userName)
	if robot.talk != nil{
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("你好，%s！我可以为你做些什么？", userName)
}

func (robot* simpleCN)Talk(heard string)(saying string,end bool,err error){
}

func (robot* simpleCN)ReportError(err error)string{
}

func (robot *simpleCN)End()error{
	return nil
}