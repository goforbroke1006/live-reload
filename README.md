# live-reload

### How to use

1. Copy binary from release page

```
RUN curl -LO https://github.com/goforbroke1006/live-reload/releases/download/0.1.0/live-reload_linux_amd64
```

2. Paste to /usr/local/bin/ directory

```
RUN mv live-reload_linux_amd64 /usr/local/bin/live-reload
```

3. Set endpoint in Dockerfile

```
ENDPOINT [ "live-reload" ]
```

4. Use in docker-compose.yaml like

```
services:
  ...
  
  my-service-WIP:
    build:
      dockerfile: Dockerfile.my-service-WIP
      context: ./
    command: python /code/app.py
    volume:
      ./:/code/

```
