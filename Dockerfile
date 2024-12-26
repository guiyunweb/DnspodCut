# 使用基于 debian 的镜像
FROM debian:latest

# 更新 apt 并安装必要的软件
RUN apt-get update && apt-get install -y iputils-ping  locales

# 设置中文环境
RUN sed -i -e 's/# zh_CN.UTF-8 UTF-8/zh_CN.UTF-8 UTF-8/' /etc/locale.gen && locale-gen

# 设置环境变量
ENV LANG zh_CN.UTF-8
ENV LANGUAGE zh_CN:zh
ENV LC_ALL zh_CN.UTF-8

# 复制 Go 项目文件到镜像中
COPY . /app

# 设置工作目录
WORKDIR /app

# 编译并运行 Go 项目
CMD ["go", "run", "main.go"]
