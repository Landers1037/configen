# conf for {{.app}}
# description: {{.des}}
server {
    listen {{.listen -}};
    server_name {{.domain -}};
    location / {
        proxy_set_header REMOTE-HOST $remote_addr;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://localhost:{{.port}};
    }
}