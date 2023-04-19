Nginx 缓存
实战一
实现效果：
在3天内，通过浏览器地址栏访问 *http://127.0.0.1:85/image/1.png* ，不会从服务器抓取资源，3天后（过期）则从服务器重新下载。
具体配置：
# http 区域下添加缓存区配置
proxy_cache_path /tmp/nginx_proxy_cache levels=1 keys_zone=cache_one:512m inactive=60s max_size=1000m;

# server 区域下添加缓存配置
location ~ \.(gif|jpg|png|htm|html|css|js)(.*) {
     proxy_pass http://172.52.1.2:8080; #如果没有缓存则转向请求
     proxy_redirect off;
     proxy_cache cache_one;
     proxy_cache_valid 200 1h;            #对不同的 HTTP 状态码设置不同的缓存时间
     proxy_cache_valid 500 1d;
     proxy_cache_valid any 1m;
     expires 3d;
}
expires 是给一个资源设定一个过期时间，通过 expires 参数设置，可以使浏览器缓存过期时间之前的内容，减少与服务器之间的请求和流量。也就是说无需去服务端验证，直接通过浏览器自身确认是否过期即可，所以不会产生额外的流量。此种方法非常适合不经常变动的资源。