module p1

go 1.19


// 不同的项目下使用 require 和 replace
require (
	p2 v0.0.0	// v是小写，大写会出现报错
)
replace (
	p2 => ../p2
)
