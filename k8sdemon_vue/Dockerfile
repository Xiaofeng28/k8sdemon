FROM nginx:1.28
# 复制 Nginx 配置文件
COPY server.conf /etc/nginx/conf.d/
# 复制前端文件
COPY dist/ /usr/share/nginx/html/
# 移除默认配置
RUN rm /etc/nginx/conf.d/default.conf
# 暴露 80 端口
EXPOSE 80
# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --retries=3 --start-period=0s \
    CMD curl -f http://localhost/ || exit 1
# 运行命令
CMD ["nginx", "-g", "daemon off;"]