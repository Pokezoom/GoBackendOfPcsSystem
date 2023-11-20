#!/bin/bash

# 检查并安装 Homebrew
if ! command -v brew &> /dev/null
then
    echo "安装 Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# 创建 pcs_project 文件夹
mkdir -p ~/pcs_project
cd ~/pcs_project

# 克隆 Git 项目
git clone https://github.com/Pokezoom/GoBackendOfPcsSystem.git

# 判断处理器架构并下载安装 Golang
ARCH=$(uname -m)
if [ "$ARCH" == "x86_64" ]; then
  # 对于 x86 架构
  curl -OL https://go.dev/dl/go1.19.10.darwin-amd64.tar.gz
  tar -C /usr/local -xzf go1.19.10.darwin-amd64.tar.gz
elif [ "$ARCH" == "arm64" ]; then
  # 对于 ARM 架构
  curl -OL https://go.dev/dl/go1.19.10.darwin-arm64.tar.gz
  tar -C /usr/local -xzf go1.19.10.darwin-arm64.tar.gz
fi

# 设置环境变量
export PATH=$PATH:/usr/local/go/bin

# 检查是否安装了 MySQL
if ! mysql --version &> /dev/null
then
    echo "安装 MySQL..."
    sudo brew install mysql
    sudo brew services start mysql
fi

# 设置 MySQL 密码
echo "设置 MySQL 密码..."
sudo mysqladmin -u root password 'keke990813'

# 创建数据库和表
echo "创建数据库和数据表..."
sudo mysql -u root -pkeke990813 -e "CREATE DATABASE IF NOT EXISTS dev_pcs;"
sudo mysql -u root -pkeke990813 dev_pcs < create_tables.sql

# 进入项目目录并执行 go mod tidy
cd GoBackendOfPcsSystem
go mod tidy

# 启动项目
go run main.go &

echo "安装和部署完成！"
