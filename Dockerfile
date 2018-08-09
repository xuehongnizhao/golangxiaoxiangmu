FROM 192.168.10.14:5000/alpine:v1

MAINTAINER ZhangQiang <494000616@qq.com>

RUN mkdir -p /var/deploy/msgnotification/

# 文件移动
ADD msgnotification /var/deploy/msgnotification
ADD msgnotification.conf /var/deploy/msgnotification

COPY static/ /var/deploy/msgnotification/static/

COPY views/ /var/deploy/msgnotification/views/


# 切换目录
WORKDIR /var/deploy/msgnotification/

EXPOSE 8080

RUN chmod +x msgnotification

ENTRYPOINT ./msgnotification -r 192.168.10.127 -s msgnotification -m dev



