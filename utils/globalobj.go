package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

/*
存储一切有关Zinx框架的全局参数，供其他模块使用
一些参数也可以通过 用户根据 zinx.json来配置
*/
type GlobalObj struct {
	TcpServer ziface.IServer //当前Zinx的全局Server对象
	Host string //当前服务器主机IP
	TcpPort int //当前服务器主机监听端口号
	Name string //当前服务器名称
	Version string //当前Zinx版本号
	MaxPacketSize uint32 //都需数据包的最大值
	MaxConn int //当前服务器主机允许的最大链接个数
}

/*
提供init方法，默认加载
*/
func init(){
	//初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Host:          "0.0.0.0",
		TcpPort:       7777,
		Name:          "weiweiServerApp",
		Version:       "v0.4",
		MaxPacketSize: 4096,
		MaxConn:       12000,
	}

	//从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}

/*
定义一个全局的对象
*/
var GlobalObject  *GlobalObj

//读取用户的配置文件
func (g *GlobalObj) Reload(){
	data,err := ioutil.ReadFile("../conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将json数据解析到struct中
	err = json.Unmarshal(data,&GlobalObject)
	if err != nil {
		panic(err)
	}
}