# portster
Inside-the-box Docker REST service to easily access externally-mapped ports.

GET /port/<num> - Returns a port # of externally-mapped port or 404 if not mapped.

On docker-machine hosts, TLS security is enabled by default.  You need to make
sure to run your Docker images something like this:

```bash
#!/bin/bash

ACTIVE=`docker-machine active 2>/dev/null`
HOST_IP=`docker-machine ip $ACTIVE`

docker run -it -P -v ~/.docker/machine/certs:/mnt/certs -e "DOCKER_TLS_VERIFY=true" -e HOST_IP=$HOST_IP myImage
```

This mounts the cert pem files to a known point inside the Docker, /mnt/certs, and sets a flag to support TLS.

If you are not running in a TLS environment, for example the default configuration of the AWS ECS-enabled instance, you would run your image with something like this:

```bash
#!/bin/bash

docker run -it -P  -v /var/run/docker.sock:/var/run/docker.sock myImage
```

This mounts the (unsecured) Docker UNIX socket to a known point inside the Docker so Porster can find it.