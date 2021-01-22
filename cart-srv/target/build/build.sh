export GO111MODULE=auto
export GOPROXY=https://goproxy.cn,direct
export GOPRIVATE="code.aliyun.com"

# 打包
copy() {
    # 生成项目文件
    rm -rf ./target
    mkdir ./target
    cp -R ./build ./target/
    CGO_ENABLED=0 GOOS=linux go build -a -o ./target/cart-srv
    mkdir ./target/logs
    cp -R ./config ./target
}
copy

