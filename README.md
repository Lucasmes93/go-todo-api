# 📌 API de Gestion des Tâches en Go (Gin)

👥 **Groupe Go** : Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX

🔗 **Dépôt GitHub** : [go-todo-api](https://github.com/Lucasmes93/go-todo-api)

---

## 📝 Description du Projet
Cette API permet d'ajouter, récupérer et gérer des tâches en utilisant **Go** et le framework **Gin**. Elle offre une gestion simple des tâches via des requêtes **RESTful** (`GET`, `POST`).

---

## 🚀 Installation & Exécution

### 🔹 **Prérequis**
- **Go** installé (version **1.24 ou ultérieure**)
- **Git** installé

### 🔹 **Installation**

1️⃣ **Cloner le dépôt**
```sh
git clone https://github.com/Lucasmes93/go-todo-api.git
cd go-todo-api
```

2️⃣ **Installer les dépendances**
```sh
go mod tidy
```

3️⃣ **Lancer le serveur**
```sh
go run main.go
```
Le serveur tourne maintenant sur : **[http://localhost:8080](http://localhost:8080)**

---

## 🔗 **Endpoints de l'API**

📌 **Obtenir toutes les tâches**
- **Méthode** : `GET`
- **URL** : `/tasks`

📌 **Ajouter une nouvelle tâche**
- **Méthode** : `POST`
- **URL** : `/tasks`
- **Exemple de requête JSON** :
  ```json
  {
    "id": 1,
    "title": "Faire les courses"
  }
  ```
- **Exemple de réponse JSON** :
  ```json
  {
    "id": 2,
    "title": "Apprendre Go"
  }
  ```

---

## ✍️ **Auteurs**
Ce projet a été réalisé dans le cadre du Groupe Go par :  
**Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX.**  
