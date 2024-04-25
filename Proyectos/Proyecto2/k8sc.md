

## Comandos para crear cluster de kubernetes

### Iniciar config
```
gcloud init
```

### Setear el proyecto
```
gcloud config set project so1-1s-202111478
```

### Configurar la zona
```
gcloud config set compute/zone us-central1-a
```

### Creacion del cluster
```
gcloud container clusters create cluster-so1-p2 --num-nodes=1 --tags=allin,allout --enable-legacy-authorization --issue-client-certificate --preemptible --machine-type=n1-standard-2
```

### Credenciales

Se establece conexion en GCP y se copia el comando, en este caso es este

```
gcloud container clusters get-credentials cluster-so1p2 --zone us-central1-a --project so1-1s-202111478
```

### Comandos para kafka

```
kubectl create ns kafka

kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka

kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n kafka

```

### Direcciones de CLOUD RUN

```
INGRESS
http://34.29.26.106

FRONTEND
https://so1p2-webfront-7fugfcr5eq-uc.a.run.app

BACKEND
https://so1p2-webapi-7fugfcr5eq-uc.a.run.app

```