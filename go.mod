module github.com/oceanweave/testhello

go 1.21.4

// 通过 commit-id 进行拉取  go get -u github.com/oceanweave/testgomod@{commit-id(前7位及以上都行）}
// go get -u github.com/oceanweave/testgomod@54453547c566697b28954d2290d5859ba846536a
// v0.0.0-20231118074428-54453547c566 三个部分分别表示，版本号，commit 时间，commit-id
//require github.com/oceanweave/testgomod v0.0.0-20231118074428-54453547c566 // indirect

// 利用 go get -u github.com/oceanweave/testgomod@v1.0.1 命令进行拉取
// 同时还支持分支拉取，此处不做展示，命令类似  go get -u github.com/oceanweave/testgomod@{分支名}
require github.com/oceanweave/testgomod v1.0.1 // indirect
