FROM 192.168.10.14:5000/alpine:v1

MAINTAINER WangXue <649229856@qq.com>

RUN mkdir -p /var/deploy/ykt/supermarket/

# 文件移动
ADD supermarket /var/deploy/ykt/supermarket
ADD supermarket.conf /var/deploy/ykt/supermarket

COPY static/ /var/deploy/ykt/supermarket/static/

COPY views/ /var/deploy/ykt/supermarket/views/


# 切换目录
WORKDIR /var/deploy/ykt/supermarket/

EXPOSE 8080

RUN chmod +x supermarket

ENTRYPOINT ./supermarket -r 192.168.10.127 -s supermarket -m dev



