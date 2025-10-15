# 使用一个预装了 Go 和 Node.js 的基础镜像来简化步骤
# (你也可以选择一个纯净系统然后自己安装)
# 这里我们选择一个包含 Go 的，然后手动安装 Node.js
FROM golang:1.25

# 安装 Node.js 20.x, pnpm 和 supervisor
RUN apt-get update && \
    apt-get install -y curl supervisor && \
    curl -fsSL https://deb.nodesource.com/setup_24.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g pnpm

# 设置工作目录
WORKDIR /app

# 复制所有项目文件
COPY . .

# 设置 Go 代理以加速依赖下载
ENV GOPROXY=https://goproxy.cn,direct
# 下载 Go 依赖
RUN go mod download

# 进入前端目录安装依赖并构建
WORKDIR /app/frontend
RUN pnpm install --force
RUN pnpm build

# 回到根目录
WORKDIR /app

# 复制 supervisor 配置文件
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# 暴露前端预览服务器的端口 (默认 4173) 和后端 Go 服务的端口
EXPOSE 4173
EXPOSE 8123

# 启动 supervisor
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]