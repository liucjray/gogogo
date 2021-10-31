# 簡易說明
1. 安裝 GO
   - ``sudo add-apt-repository ppa:longsleep/golang-backports``
   - ``sudo apt-get update``
   - ``sudo apt-get install -y golang-go``

2. github 下載專案

3. PORT
    + 確認 PORT 是否開啟
        - `` sudo iptables -L -n ``
    + 若未開啟則開啟所需 PORT
        - `` sudo iptables -I INPUT -p tcp --dport 8080 -j ACCEPT ``

4. SUPERVISOR
    + 配置檔
        - /etc/supervisor/conf.d/go.conf
        - ``
[program:go-gogogo]
directory=/home/{$USER}/codes/go-gogogo/gogogo
command=/usr/bin/go run /home/{$USER}/codes/go-gogogo/gogogo/index.go
autostart=true
autorestart=true
environment=GOPATH="/home/{$USER}/codes",HOME="/home/{$USER}"
stderr_logfile=/var/log/supervisor/go.err.log
;stdout_logfile=/var/log/supervisor/go.out.log
``
    + 端口報錯檢查
        - `` netstat -tulpn | grep :8080 ``
        - 若是有端口占用問題，則使用 kill {$ID}

    + 更新並且重啟 Supervisor config
        - `` sudo supervisorctl reload && sudo supervisorctl restart all ``

    + 上述報錯開起 stderr_logfile 並確認錯誤原因

2. NGINX
    + 配置檔使用 proxy_pass 即可
        - /etc/nginx/sites-available/default
        - ``
server {
    listen 80;
    listen [::]:80;
    server_name  {$DOMAIN};
    location / {
        proxy_pass http://localhost:8080;
    }
}
``