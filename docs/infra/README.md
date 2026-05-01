# Setup infra on Azure

## Requirements

### VS Code Package
- Azure Tools

### OS
- Azure CLI
- Docker

### Azure
- Container Registry

---

## Build image Docker local 

### Backend
Pastikan `main.go` berada di folder yang sama dengan `Dockerfile`.

Jangan pakai `loadenv` on local

Pastikan sudah setup database online.

```
docker build -t etc-be:dev .
docker run --rm -p 8080:8080 --env-file .env etc-be:dev
```

### FrontEnd
```
docker build -t etc-fe:dev --build-arg NEXT_PUBLIC_API_URL={YOUR_API} .
docker run --rm -p 3000:3000 etc-fe:dev
```

---

## Setup ACR
Follow tutorial ini: [here](https://www.youtube.com/watch?v=7FjPZia53BU&theme=dark)

### Push Docker Image into ACR
U can see it on Docker Desktop.
```
az login
az acr login --name {name_acr} -g {name_resource_group}

# tag each repo and push, example :
    docker tag etc-be:dev etcimage.azurecr.io/etc-be:prod
    docker push etcimage.azurecr.io/etc-be:prod

    docker tag etc-fe:dev etcimage.azurecr.io/etc-fe:prod
    docker push etcimage.azurecr.io/etc-fe:prod
```

---

## Managed Identities
Create first, then assign the user-assigned role for ACR pull.

---

## Setup Azure Container App
- Make new container env if not exist (same region/location)
- In Container, use ACR as image source, then select registry, image, and tag
- Authentication type is managed identity
- Set environment variable from secret, then input on container

### Target Port
- FE: `{YOUR_FE_PORT}`
- BE: `{YOUR_BE_PORT}`

---


Setelah mendapatkan application URL BE, maka command FE bisa disesuaikan :

```
docker build -t etc-fe:dev --build-arg NEXT_PUBLIC_API_URL={YOUR_APPLICATION_URL_BE} .
docker run --rm -p 3000:3000 etc-fe:dev
```

## Application URL

Akan muncul di overview tiap container, flow :
Overview => Application Url

---

## Notes
- Pastikan mengganti tag setiap kali deployment agar sesuai.
- Untuk BE, jangan lupa nonaktifkan scan `.env` di `main.go` sebelum deploy.
