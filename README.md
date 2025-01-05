# Chronicles of Aethelgard

![Chronicles of Aethelgard](chronicles.jpeg)

## Snowcamp 2025

### Prérequis :

- Docker & Docker Compose
- VSCode
- Git

Et ensuite :

```bash
docker pull k33g/chronicles-of-aethelgard:0.0.0
docker pull FROM golang:1.23.4-alpine 
```

> Le Dockerfile de `k33g/chronicles-of-aethelgard` est disponible dans le répertoire `docker`.

Pour vérifier que tout fonctionne :

```bash
docker compose up
```
> Si vous n'avez pas fait les `docker pull` précédents, Docker Compose les fera pour vous.

> 🖐️ Même si vous ne codez pas en Go, le code sera très simple.