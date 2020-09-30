# live-reload

### How to use

1. Copy binary from release page
2. Paste to /usr/local/bin/ directory
3. Set endpoint in Dokerfile

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
