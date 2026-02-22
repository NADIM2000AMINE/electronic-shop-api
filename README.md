# Electronic Shop Management API

## 📝 Description du projet
Ce projet est une API REST complète développée en Go pour gérer plusieurs boutiques d'électronique. Il intègre une architecture avec une isolation stricte multi-tenant, une gestion des rôles internes (SuperAdmin, Admin), une page publique pour les clients, et une génération dynamique de liens WhatsApp pour contacter les boutiques.

## 🛠 Stack Technique
- **Langage** : Go (1.25)
- **Framework Web** : Fiber v2
- **Base de données** : PostgreSQL 14
- **ORM** : GORM v2
- **Authentification** : JWT (golang-jwt) & bcrypt
- **Infrastructure** : Docker & Docker Compose

## ⚙️ Prérequis
Pour exécuter ce projet du premier coup, vous devez avoir installé sur votre machine :
- **Docker** et **Docker Compose** (fortement recommandé)
- **Go** (version 1.21 ou supérieure, pour exécuter le script de Seed)
- **Git**

## 🚀 Installation et Exécution étape par étape

**1. Cloner le repository**
\`\`\`bash
git clone https://github.com/NADIM2000AMINE/electronic-shop-api.git
cd electronic-shop-api
\`\`\`

**2. Configurer les variables d'environnement**
Créez un fichier `.env` à la racine du projet et copiez-y cette configuration :
\`\`\`env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=electronic_shop
JWT_SECRET=super-secret-jwt-key-2026
SERVER_PORT=3000
\`\`\`

**3. Lancer l'infrastructure avec Docker**
Exécutez cette commande pour construire l'API et lancer la base de données PostgreSQL. Le port `5433` de la machine hôte est mappé sur le `5432` du conteneur pour éviter les conflits locaux.
\`\`\`bash
docker compose up --build -d
\`\`\`
*Patientez quelques secondes le temps que la base de données s'initialise.*

**4. Injecter les données de test (Seed)**
Pour peupler la base de données avec 2 boutiques, 2 utilisateurs et 3 produits, exécutez le script depuis votre machine locale (qui ciblera le conteneur Docker via le port 5433) :
\`\`\`bash
DB_HOST=localhost DB_PORT=5433 DB_USER=postgres DB_PASSWORD=postgres123 DB_NAME=electronic_shop go run scripts/seed.go
\`\`\`

**5. Tester l'API**
Le serveur tourne désormais sur `http://localhost:3000`. 
Vous pouvez importer la collection Postman incluse (`electronic-shop.postman_collection.json`) pour tester les routes.

### 🔐 Identifiants de test générés
- **SuperAdmin** : `super@techstore.com` / `password123` (Peut gérer le numéro WhatsApp et voir les profits)
- **Admin** : `admin@techstore.com` / `password123` (Gère le stock et les transactions, ne voit pas le prix d'achat)