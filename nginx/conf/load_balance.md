负载均衡
实战一
实现效果：
在浏览器地址栏输入 *http://192.168.4.32/example/a.html* ，平均到 8080 和 8081 端口中，实现负载均衡效果。
具体配置：
    upstream myserver {
      server 172.52.1.2:8080;
      server 172.52.1.2:8081;
    }
    server {
        listen       80;   #监听端口
        server_name  172.52.1.2;   #监听地址

        location  / {       
           root html;  #html目录
           index index.html index.htm;  #设置默认页
           proxy_pass  http://myserver;  #请求转向 myserver 定义的服务器列表      
        } 
    }
nginx 分配服务器策略
轮询（默认）
按请求的时间顺序依次逐一分配，如果服务器down掉，能自动剔除。

权重
weight 越高，被分配的客户端越多，默认为 1。比如：
      upstream myserver {
        server 172.52.1.2:8080 weight=10;
        server 172.52.1.2:8081 weight=5;
      }

ip
按请求 ip 的 hash 值分配，每个访客固定访问一个后端服务器。比如：
      upstream myserver {
        ip_hash;  
        server 172.52.1.2:8080;
        server 172.52.1.2:8081;
      }

fair
按后端服务器的响应时间来分配，响应时间短的优先分配到请求。比如：
      upstream myserver {
        fair;  
        server 172.52.1.2:8080;
        server 172.52.1.2:8081;
      }