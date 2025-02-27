# 📌 API de Gestion des Tâches en Go (Gin)

👥 **Groupe Go** : Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX

🔗 **Dépôt GitHub** : [go-todo-api](https://github.com/Lucasmes93/go-todo-api)

---

## 📝 Description du Projet
Cette API permet d'ajouter, récupérer, modifier et supprimer des tâches en utilisant **Go** et le framework **Gin**.  
Elle repose sur une architecture **RESTful** et offre une gestion simple des tâches via les requêtes `GET`, `POST`, `PUT` et `DELETE`.

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
- **Exemple de réponse JSON** :
  ```json
  [
    {
      "id": 1,
      "title": "Faire les courses"
    },
    {
      "id": 2,
      "title": "Apprendre Go"
    }
  ]
  ```

📌 **Ajouter une nouvelle tâche**
- **Méthode** : `POST`
- **URL** : `/tasks`
- **Exemple de requête JSON** :
  ```json
  {
    "title": "Acheter du pain"
  }
  ```
- **Exemple de réponse JSON** :
  ```json
  {
    "id": 3,
    "title": "Acheter du pain"
  }
  ```

📌 **Modifier une tâche existante**
- **Méthode** : `PUT`
- **URL** : `/tasks/:id`
- **Exemple de requête JSON** :
  ```json
  {
    "title": "Faire du sport"
  }
  ```
- **Exemple de réponse JSON** :
  ```json
  {
    "id": 1,
    "title": "Faire du sport"
  }
  ```

📌 **Supprimer une tâche**
- **Méthode** : `DELETE`
- **URL** : `/tasks/:id`
- **Exemple de réponse JSON en cas de succès** :
  ```json
  {
    "message": "Tâche supprimée"
  }
  ```
- **Exemple de réponse JSON si la tâche n'existe pas** :
  ```json
  {
    "error": "Tâche non trouvée"
  }
  ```

---

## 🐳 **Exécuter avec Docker**

Si vous souhaitez exécuter l'API dans un **contenant Docker**, utilisez la commande suivante pour **construire l'image Docker** :

```sh
docker build --pull --rm -f "Dockerfile" -t "gotodoapi:latest" "."
```

Puis, exécutez le conteneur :
```sh
docker run -p 8080:8080 gotodoapi:latest
```

L'API sera accessible sur **[http://localhost:8080](http://localhost:8080)**.

---

## ✍️ **Auteurs**
Ce projet a été réalisé dans le cadre du Groupe Go par :  
**Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX.**

