module github.com/oceanweave/testhello

go 1.21.4

// 通过 commit-id 进行拉取
// go get -u github.com/oceanweave/testgomod@54453547c566697b28954d2290d5859ba846536a
// v0.0.0-20231118074428-54453547c566 三个部分分别表示，版本号，commit 时间，commit-id
require github.com/oceanweave/testgomod v0.0.0-20231118074428-54453547c566 // indirect
