#反向代理
  反向代理指代理后端服务器响应客户端请求的一个中介服务器，代理的对象是服务端。
实战一
实现效果：
在浏览器输入 http://127.0.0.1:83 , 从 nginx 服务器跳转到 linux 系统 tomcat 主页面。
具体配置：
    server {
        listen       83;
        server_name  localhost;   #监听地址
        location  / {
           root html;  #/html目录
           proxy_pass http://172.52.1.2:8080;  #请求转向
           index  index.html index.htm;      #设置默认页
        }
    }
实战二
实现效果：
根据在浏览器输入的路径不同，跳转到不同端口的服务中。
具体配置：
    server {
        listen       83;
        server_name  localhost;   #监听地址
        location  ~ /admin/ {  
           proxy_pass http://172.52.1.2:8080;
        }

        location  ~ /server/ {  
           proxy_pass http://172.52.1.2:8081;         
        } 
    }