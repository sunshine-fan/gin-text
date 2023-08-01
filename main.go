package main

import (
	"hello/routes"
)

//生成token
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InExbWkiLCJpc3MiOiJteS1wcm9qZWN0IiwiZXhwIjoxNjkwOTYwODM2fQ.d9zSG3jSkfWAIbnfZHmFs8wv_x1Ti3ePFKzLk2-WjVw

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

type UserInfo struct {
	// 可根据需要自行添加字段
	Username string
	Password string
}

func main() {
	router := routes.NewRouter()
	router.Run()
}

