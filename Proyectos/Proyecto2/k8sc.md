

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
gcloud container clusters get-credentials cluster-so1-p2 --zone us-central1-a --project so1-1s-202111478
```

