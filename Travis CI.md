
 - 添加.travis.yml文件告诉Travis CI去做什么。
 - 添加.travis.yml到git
 
 Travis CI provides a default build environment and a default set of steps for each programming language。
 
 
## atlas项目

 - https://docs.travis-ci.com/user/languages/go
 让我们的项目里的各个branch都被自动CI系统Travis CI监控。
 它会build我们push的每一个branch。如果unit test不过，这个branch对应的PR就不能被merge。
 - .travis.yml

```
 language: go

go:
  - 1.6
```
