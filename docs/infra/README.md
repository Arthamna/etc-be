Setup infra on Azure :

# Requirements :

VS Code Package :

- Azure Tools


OS :
- Azure CLI
- Docker

Azure :
- Container Registry 



Coba build image docker di local dulu to make sure this works

---

Build local be :

pastikan main.go berada di folder sama dengan dockerfile


jangan pakai loadenv local (?)

setup supabase sudah, lewat transaction pooler

```
docker build -t etc-be:dev .
docker run --rm -p 8080:8080 --env-file .env etc-be:dev
```

docker_image ready lewat command

---


Build local fe :



```
docker build -t etc-fe:dev --build-arg NEXT_PUBLIC_API_URL=http://localhost:8080 .
docker run --rm -p 3000:3000 etc-fe:dev
```

Setup ACR

Follow this [here](https://www.youtube.com/watch?v=7FjPZia53BU&theme=dark)


Push Docker Image into ACR

U can see it on docker desktop

```
az login
az acr login --name etcimage -g DefaultResourceGroup-EA
```


```

docker tag etc-be:dev etcimage.azurecr.io/etc-be

docker push etcimage.azurecr.io/etc-be

docker tag etc-fe:dev etcimage.azurecr.io/etc-fe

docker push etcimage.azurecr.io/etc-fe
```

But how about proxy request ?
