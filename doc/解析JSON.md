目录

[toc]

2022 年 4 月 7 号

### 一、reflect 模式解析 JSON

使用内置的 reflect 的方法可以用于解析配置文件，但是对于 QPS 较高的场景不适合

解析方式如下：

struct 类型

```go
package reflectjson

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

```

解析 JSON

```go
package reflectjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java", "Go", "c"]
	}
}`

// 适用于解析配置文件，不适合在QPS高的场景，使用的机制采用的是反射，性能低
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
```

### 二、使用 easyJSON 解析 JSON

首先下载 easyJSON，`go get -u github.com/mailru/easyjson/...`

然后使用 easuJSON 生成解析 JSON 的代码，包含了解析 JSON 的方法，生成代码命令如下，struct_json.go 就是包含了 json 的 go 文件。

```visual basic
~/go/bin/easyjson struct_json.go
```

执行该命令后会生成一个 struct_json_easyjson.go 的文件，里面包含了 MarshalJSON()函数，UnmarshalJSON()函数等。

使用生成的 easyjson 代码解析 JSON

```go
package easyjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java", "Go", "c"]
	}
}`

// 适用于解析配置文件，不适合在QPS高的场景，使用的机制采用的是反射，性能低
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

// 使用easyjson 自动生成解析的代码
func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(v))
	}

}

// reflect 方式实现
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}

// easyJSON 生成的代码方法实现
func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err = e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}

```

### 三、bench 测试 reflect 和 easyjson 解析 JSON 效率

使用`go test -bench=.`命令测试，表示测试当前目录下的所有 bench 测试用例。

如上面两个 bench 测试方法，测试结果如下所示：

```visual basic
csgo/json/easyJson on  main [!?] via 🐹 v1.17.12
❯ go test -bench=.
{{Mike 30} {[Java Go c]}}
{"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","c"]}}
{{Mike 30} {[Java Go c]}}
{"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","c"]}}
goos: darwin
goarch: arm64
pkg: csgo/json/easyJson
BenchmarkEmbeddedJson-8   	  593142	      2009 ns/op
BenchmarkEasyJson-8       	 2047489	       589.2 ns/op
PASS
ok  	csgo/json/easyJson	4.453s
```

可以看出 easyjson 解析的效率更好。