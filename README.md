# Chronicles of Aethelgard

![Chronicles of Aethelgard](chronicles.jpeg)

## Workshop IA au Snowcamp 2025

L'objectif de ce workshop est de voir comment utiliser un modèle de langage (LLM) pour nous aider à créer un jeu de rôle textuel (RPG).

Nous verrons comment utiliser un modèle de langage pour générer des textes pour diverses situations (description d'une pièce, création d'un personnage, Bot PNJ ...). Ce sera l'occasion de voir comment utiliser l'API d'Ollama pour interagir avec un LLM.

✋ **VOUS N'AVEZ PAS BESOIN DE GPU** 🎉 *(tous mes exemples fonctionnent aussi sur un Raspberry Pi5 8GB RAM)*. 

🐣 Nous verrons donc comment tirer parti d'un **Small Language Model** (ou Tiny Language Model).

🍒 Et si il nous reste du temps, nous tenterons d'intégrer nos outils dans un jeu de rôle textuel (RPG). 

### Prérequis :

- Docker & Docker Compose
- VSCode
- Git

Et ensuite : **À FAIRE AVANT DE VENIR** 🙏

```bash
docker pull k33g/chronicles-of-aethelgard:0.0.1
docker pull FROM golang:1.23.4-alpine 
```

> Le Dockerfile de `k33g/chronicles-of-aethelgard` est disponible dans le répertoire `docker`.

Pour vérifier que tout fonctionne :

```bash
docker compose up
```
> Si vous n'avez pas fait les `docker pull` précédents, Docker Compose les fera pour vous.

> 🖐️ Même si vous ne codez pas en Go, le code sera très simple.