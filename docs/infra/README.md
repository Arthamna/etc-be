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

Managed Identities 
Create first, then assign the user-assigned role for ACR pull


```
```

Setup Azure container App

Make new container env if not exist (same region/location)

in Container, use ACR as image sources, then select registry, image, and tag

authentication type is managed identity

Set env variable from secret then input on container

Target Port :
fe : 3000
be : 8080

```

```

Urlnya sudah ada (di Application URL), tapi cara agar custom ??

(beli domain dulu, but I don't have money for that)

Azure DNS Public Zone

Pertama, cari CNAME target, muncul di overview -> Application URL 

So, sekarang setup pakai yang free aja


setelah mendapatkan application url be, maka commandnya bisa disesuaikan :

```
docker build -t etc-fe:dev --build-arg NEXT_PUBLIC_API_URL={YOUR_APPLICATION_URL_BE} .
docker run --rm -p 3000:3000 etc-fe:dev
```

Anyway, pastikan mengganti tag setiap kali deployment agar sesuai 

buat be, jangan lupa nonaktifkan scan .env (di main) sebelum deploy
```

docker tag etc-be:dev etcimage.azurecr.io/etc-be:prod1

docker push etcimage.azurecr.io/etc-be:prod1


docker tag etc-fe:dev etcimage.azurecr.io/etc-fe:prod
docker push etcimage.azurecr.io/etc-fe:prod
```


```


```

But how about proxy request ?
