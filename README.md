# k8s-playground

## Cluster

### Vagrant + Ansible

#### Required

```
$ vagrant plugin install vagrant-vbguest
```

#### Create Cluster

```
$ vagrant up
```

### Docker Desktop

#### Create Cluster

Preferences -> Kubernetes -> Enable

## Deploy

```
$ kubectl apply -f manifests/api
$ kubectl apply -f manifests/processor
$ kubectl apply -f manifests/resource
```
