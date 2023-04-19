FROM centos:7
ENV PATH /usr/local/go/bin:$PATH
RUN yum install -y wget
RUN yum install -y git
RUN cd /usr/local \
	&&	wget https://dl.google.com/go/go1.20.3.linux-amd64.tar.gz \
	&& tar  zxvf go1.20.3.linux-amd64.tar.gz -C /usr/local \
	&& rm -rf go1.20.3.linux-amd64.tar.gz \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w GO111MODULE=on
WORKDIR /wwwroot