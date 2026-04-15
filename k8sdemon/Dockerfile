FROM alpine:3.22
# 设置工作目录
WORKDIR /opt/k8sdemon
# 复制必要的文件
COPY k8sdemon .
COPY config.yaml .
# 赋予可执行权限
RUN chmod +x k8sdemon
# 暴露端口
EXPOSE 8080
# 运行命令
CMD ["./k8sdemon"]