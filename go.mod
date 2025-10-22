module github.com/jhw66/hello_module  //Go 要求模块名必须与实际仓库路径完全一致，否则依赖解析会失败。

go 1.25.1

require rsc.io/quote v1.5.2

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
