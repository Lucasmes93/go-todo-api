# 📌 API de Gestion des Tâches en Go (Gin)

👥 **Groupe Go** : Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX

🔗 **Dépôt GitHub** : [go-todo-api](https://github.com/Lucasmes93/go-todo-api)

---

## 📝 **Description du Projet**
Cette API permet d'**ajouter, récupérer, modifier et supprimer** des tâches en utilisant **Go** et le framework **Gin**.  
Elle inclut une **persistance des tâches** via un fichier **JSON (`tasks.json`)**, permettant de conserver les données après un redémarrage.

---

## 🚀 **Installation & Exécution**

### **🔹 Prérequis**
- **Go** installé (version **1.16 ou ultérieure**)
- **Git** installé
- **Docker** (optionnel, si tu veux utiliser un conteneur)

---

### **🔹 Installation**

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
📌 **Le serveur tourne sur :** [http://localhost:8080](http://localhost:8080)

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

## 💾 **Persistance des Données**
L’API utilise un fichier **JSON (`tasks.json`)** pour stocker les tâches.  
✅ **Les tâches restent enregistrées après un redémarrage.**  
✅ **Chaque ajout, modification ou suppression est sauvegardé automatiquement.**

---

## 🐳 **Utilisation avec Docker**
Si tu veux exécuter l'API dans un **conteneur Docker**, suis ces étapes :

1️⃣ **Construire l’image Docker**
```sh
docker build --pull --rm -f "Dockerfile" -t "gotodoapi:latest" .
```

2️⃣ **Lancer le conteneur**
```sh
docker run -p 8080:8080 gotodoapi:latest
```

📌 **L’API sera accessible sur :** [http://localhost:8080](http://localhost:8080)

---

## 📂 **Gestion des fichiers ignorés (`.gitignore` & `.dockerignore`)**
📌 **Fichiers ignorés dans `.gitignore`** :
```sh
tasks.json
```
📌 **Fichiers ignorés dans `.dockerignore`** :
```sh
.idea
.github
README.md
tasks.json
```
✅ **Cela évite que `tasks.json` (les données locales) soit ajouté à Git ou copié dans l’image Docker.**

---

## ✍️ **Auteurs**
Ce projet a été réalisé dans le cadre du Groupe Go par :  
**Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX.**
