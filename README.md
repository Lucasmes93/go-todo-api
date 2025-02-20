# 📝 API de Gestion des Tâches en Go (Gin)## 👥 Groupe Go- **Ianis CHENNAF**- **Philippe Ivan MBARGA**- **Mateo OUDART**- **Salman Ali MADEC**- **Lucas MESSIA DOLIVEUX**
🔗 **Lien du dépôt GitHub** : [go-todo-api](https://github.com/Lucasmes93/go-todo-api)
---
## 📌 Description du ProjetCette API permet d'ajouter, récupérer et gérer des tâches en utilisant **Go** et le framework **Gin**.  Elle offre une gestion simple des tâches via des requêtes **RESTful** (`GET`, `POST`).
---
## 🚀 Installation & Exécution### **1️⃣ Prérequis**- **Go** installé (version 1.24 ou ultérieure)- **Git** installé
### **2️⃣ Cloner le dépôt**```sh
git clone https://github.com/Lucasmes93/go-todo-api.git
cd go-todo-api
 


3️⃣ Installer les dépendances
shCopierModifiergo mod tidy
4️⃣ Lancer le serveur
shCopierModifiergo run main.go
Le serveur tourne maintenant sur http://localhost:8080 🚀.
📡 Endpoints de l'API
📍 Obtenir toutes les tâches
GET /tasks
📌 Exemple de réponse
jsonCopierModifier[{"id": 1,"title": "Faire les courses"}]
📝 Ajouter une nouvelle tâche
POST /tasks
📌 Exemple de requête
jsonCopierModifier{"title": "Apprendre Go"}
📌 Exemple de réponse
jsonCopierModifier{"id": 2,"title": "Apprendre Go"}
🎯 Auteurs
Ce projet a été réalisé dans le cadre du Groupe Go par :
Ianis CHENNAF
Philippe Ivan MBARGA
Mateo OUDART
Salman Ali MADEC
Lucas MESSIA DOLIVEUX