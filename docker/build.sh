
cd docker-service/
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
go build docker-service.go service.go
sudo chmod 755 docker-service



cd ../notice
go mod tidy
go build -o ./dockerfileDir/main

