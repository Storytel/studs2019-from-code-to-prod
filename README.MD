# From Code To Production

Presented during the STUDS event @ Storytel, 2019 by Johan Lejdung

## Prerequisites

Before continuing you will need to install the programming language [Go](https://golang.org/), [Docker](https://www.docker.com/products/docker-desktop) and finally [Kubernetes](https://kubernetes.io/) (although it's a bit more complicated and is optional).

If you have a Mac you can easily install the above by running:

```
brew install go docker
```

You might also have to enable Go Modules with this command:

```
export GO111MODULE=on;
```

To install Kubernetes you'll either have to have it installed on a server or download [MiniKube](https://kubernetes.io/docs/setup/minikube/).

## Go Microservice

In this repository you'll find a file called `main.go` that is our 'microservice' for this example.

It contains ~40 lines of code, which essentially:

- Contains an endpoint
  - `http://<IP>/`
  - Returns `Status 200 OK` & Prints a message
- It has a Ratelimiter, capped at 4 users per second
  - eg. No more than 4 users can call this at a time

### Connecting and executing

Once you've got Go installed you can run the application locally with:

```
go run main.go
```

Now try to connect to it with CURL for example:

```
curl http://localhost:8080 -v
```

## Docker Images

In order to be able to share this application with others - people, kubernetes ect... - we want to package it into a [Docker Image](https://docs.docker.com/v17.09/engine/userguide/storagedriver/imagesandcontainers/).

A Docker Image will package our code and it's dependencies up into a single file, ready to be executed.
To do this it requires a `Dockerfile`, it contains 11 lines of instructions for the build process.

To build the image you can execute the following command:

```
docker build -t <NAME> ./
```

You can then run it with:

```
docker run -p 8080:8080 <NAME>
```

_Where `<NAME>` is the name of the image._

Once it's running you can connect to it once again to assure it's working as expected:

```
curl http://localhost:8080 -v
```

We've now got a Image ready to share!
As a final action, you could upload this image to hub.docker.com. After you've logged in with your Docker instance you just need to execute:

```
docker push <NAME>
```

## Kubernetes

Assuming that we've uploaded the image we can now use it in Kubernetes.

In order to use Kubernetes you either have to install something like Minikube (to run it locally), or actually create a cluster at a Cloud Provider. I'm going to assume you have a cluster at Google Cloud Platform, and have installed the `kubectl` command line tool.

We will use two type of Resources in Kubernetes;

Deployment:

- Contains & executes the application(s)

Service:

- Allocates an IP and exposes the application(s) to the public web

---

You will find the Deployment file in `k8s/deployment.yaml`, the lines of interest are marked with a comment.

You can find the Service file in `k8s/service.yaml`, the lines of interest are marked with a comment.

---

Assuming you have a cluster in GCP you can upload both of these with:

```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

To monitor the resources in kubernetes you can use:

```
kubectl get pod -w | grep studs
```

or

```
kubectl get service -w | grep studs
```

Where `studs` is part of the name of your deployment/service.

If you don't have a cluster available here is a really good interactive tutorial: https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-interactive/ where you can perform similar operations.
