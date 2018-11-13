# GoMybatis
# 文档网站站点
https://zhuxiujia.github.io/gomybatis.io/#/
# 优势
GoMybatis 是根据java版 Mybatis3 的实现,基于Xorm的Engine和govaluate表达式及反射实现。
GoMybatis 内部在初始化时反射分析mapper xml生成golang的func代码，默认支持绝大部分的Java版的mybatis标签和规范,
支持标签
`<select>,<update>,<insert>,<delete>,<trim>,<if>,<set>,<foreach>`
# 使用教程,代码文件请查看/example文件夹
```
go get github.com/zhuxiujia/GoMybatis
go get "database/sql"
go get github.com/go-sql-driver/mysql
```
mapper.go 文件案例
```
//定义mapper文件的接口和结构体
type ExampleActivityMapper interface {
	SelectAll(result *[]Activity) error
	SelectByCondition(Name string, StartTime time.Time, EndTime time.Time, Page int, Size int, result *[]Activity) error
	UpdateById(arg Activity, result *int64) error
	Insert(arg Activity, result *int64) error
	CountByCondition(name string, startTime time.Time, endTime time.Time, result *int) error
}
//定义mapper文件的接口和结构体，也可以只定义结构体就行
//mapper.go文件 函数必须为2个参数（前面的参数（可空），最后一个为指针返回数据(可空)） error 为返回错误(必须定义)
type ExampleActivityMapperImpl struct {
	ExampleActivityMapper
	SelectAll         func(result *[]Activity) error
	SelectByCondition func(name string, startTime time.Time, endTime time.Time, page int, size int, result *[]Activity) error `mapperParams:"name,startTime,endTime,page,size"`
	UpdateById        func(arg Activity, result *int64) error
	Insert            func(arg Activity, result *int64) error
	CountByCondition  func(name string, startTime time.Time, endTime time.Time, result *int) error                            `mapperParams:"name,startTime,endTime"`
}
```

xml文件案例:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">
<mapper namespace="ActivityMapperImpl">
    <resultMap id="BaseResultMap" type="model.Activity">
        <id column="id" property="id" jdbcType="VARCHAR"/>
        <result column="name" property="name" jdbcType="VARCHAR"/>
        <result column="pc_link" property="pcLink" jdbcType="VARCHAR"/>
        <result column="h5_link" property="h5Link" jdbcType="VARCHAR"/>
        <result column="remark" property="remark" jdbcType="VARCHAR"/>
        <result column="create_time" property="createTime" jdbcType="TIMESTAMP"/>
        <result column="delete_flag" property="deleteFlag" jdbcType="INTEGER"/>
    </resultMap>
    <!--SelectAll(result *[]Activity)error-->
    <select id="selectAll" resultMap="BaseResultMap">
        select * from biz_activity where delete_flag=1 order by create_time desc
    </select>
</mapper>
```
实际使用mapper
```
import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
	"os"
	"fmt"
	"io/ioutil"
	"github.com/zhuxiujia/GoMybatis"
)
func main() {
  var err error
  	//mysql链接格式为         用户名:密码@(数据库链接地址:端口)/数据库名称   例如root:123456@(***.mysql.rds.aliyuncs.com:3306)/test
  	engine, err := GoMybatis.Open("mysql", "*?charset=utf8&parseTime=True&loc=Local") //此处请按格式填写你的mysql链接，这里用*号代替
  	if err != nil {
  		panic(err.Error())
  	}
  
  	file, err := os.Open("Example_ActivityMapper.xml")
  	if err != nil {
  		panic(err)
  	}
  	defer file.Close()
  
  	bytes, _ := ioutil.ReadAll(file)
  	var exampleActivityMapperImpl ExampleActivityMapperImpl
  	//设置对应的mapper xml文件
  	GoMybatis.UseProxyMapper(&exampleActivityMapperImpl, bytes, engine)
  
  	//使用mapper
  	var result []Activity
  	exampleActivityMapperImpl.SelectAll(&result)
  
  	fmt.Println(result)
}
```
