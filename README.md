
---

# GoBackendOfPcsSystem

## 项目简介
**GoBackendOfPcsSystem** 是一个学生学习状态检测系统的后端项目。该项目提供了一个后端服务，用于分析和处理关于学生在课堂上的表现和参与度的数据。

### 关于项目
- **开发目的**：为学生学习状态提供实时监控和分析。
- **项目特色**：采用 Go 语言开发，以高效性和稳定性为核心。

## 技术栈
- **Go**：后端主要开发语言，占比 96%。
- **Shell**：辅助脚本编写，占比 3%。
- **Dockerfile**：容器化部署，占比 1%。


# 克隆 Git 项目
git clone https://github.com/Pokezoom/GoBackendOfPcsSystem.git

# 判断处理器架构并下载安装 Golang
```shell
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
```
# 设置环境变量
export PATH=$PATH:/usr/local/go/bin


## 快速开始
1. **安装依赖**：
   ```bash
   cd GoBackendOfPcsSystem
   go mod tidy
   ```
2. **配置环境**：
   - 配置相关环境变量和参数。
3. **启动服务**：
   ```bash
   go run main.go
   ```

## 后端的算法参考
[点击此处访问Perception-System-of-Students-Classroom-Performance仓库](https://github.com/Pokezoom/Perception-System-of-Students-Classroom-Performance)


## 贡献指南
我们欢迎并感谢社区成员的任何贡献。请阅读我们的贡献指南来了解如何参与项目开发。

## 许可证
本项目采用 MIT。有关详细信息，请参阅 `LICENSE` 文件。

---
