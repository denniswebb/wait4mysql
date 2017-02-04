# wait4mysql
[![](https://images.microbadger.com/badges/version/dhwebb/wait4mysql:light.svg)](https://hub.docker.com/r/dhwebb/wait4mysql/)[![](https://images.microbadger.com/badges/image/dhwebb/wait4mysql:light.svg)](https://microbadger.com/images/dhwebb/wait4mysql:light)

Waits on a MySQL container to become available before exiting.

## Description

Run this container while linked to a MySQL container to cause automation scripts to pause until MySQL is ready.

## Example

```
docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=my-secret mysql
docker run --link mysql dhwebb/wait4mysql
```

## Environment Variables

You can pass in the following environment variables: `HOST`, `PORT`, `USERNAME`, `PASSWORD`, `TIMEOUT`.