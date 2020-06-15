# http-sneak

Http Sneak is a lightweight Go HTTP server that accepts any request and returns a 200 code with all the request path, parameters and headers pretty printed.

Useful for debugging Proxies, Api Gateways, etc.

### Run locally
Compile and run the server locally (it requires Go >= 1.14 installed):
1. Run: `make run`
1. Send a request
```bash
$  curl localhost:8080/test

  Path: /test
  Headers:
  	 User-Agent: [curl/7.70.0]
  	 Accept: [*/*]
```

### Run in Kubernetes
Manifests for Pod and Service are provided. The `http-sneak` image is hosted in [dockerhub](https://hub.docker.com/repository/docker/rulox/http-sneak) for simplicity.

1. Apply the `http-sneak.yaml` manifests: `kubectl apply -f http-sneak.yaml`

Additionally, if you are using an Ingress Controller, an ingress resource is provided to access the `http-sneak` pod. Apply it by running:

`kubectl apply -f http-sneak-ingress.yaml`

This will create an ingress rule for the host "example.com" and path `/http-sneak`. Simply use curl to access the pod (Note: $IP is your Ingress Controller pod)
```bash
$ curl -H "Host: example.com" http://<$IP>:80/http-sneak
```

### License

MIT