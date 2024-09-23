project-root/
│
├── backend/
│   ├── pkg/
│   │   ├── db/
│   │   │   ├── migrations/
│   │   │   │   └── sqlite/
│   │   │   │       ├── 000001_create_users_table.up.sql
│   │   │   │       ├── 000001_create_users_table.down.sql
│   │   │   │       ├── 000002_create_posts_table.up.sql
│   │   │   │       ├── 000002_create_posts_table.down.sql
│   │   │   └── sqlite.go  # Gère la connexion à la DB SQLite et les migrations
│   │   └── models/        # Fichiers pour les modèles de base de données (Users, Posts, Followers)
│   ├── controllers/       # Gestion des requêtes pour les routes (auth, posts, profiles)
│   ├── middlewares/       # Authentification, gestion des sessions, validation des données
│   ├── services/          # Logique métier, services pour Followers, Notifications, Chats
│   ├── websockets/        # Gestion des connexions WebSocket (chat en temps réel, notifications)
│   ├── utils/             # Fonctions utilitaires (gestion des fichiers, validation des entrées)
│   └── server.go          # Point d'entrée principal du serveur backend
│
├── frontend/
│   ├── public/            # Fichiers statiques, images
│   ├── src/
│   │   ├── components/     # Composants React/Vue/Svelte (ex. Profile, Post, Group)
│   │   ├── pages/          # Pages principales (Home, Profile, Post Details, Groups)
│   │   ├── services/       # Gestion des appels API vers le backend (auth, posts, followers)
│   │   └── App.js          # Fichier principal de l'application frontend
│   ├── package.json        # Dépendances et scripts de build frontend
│   └── Dockerfile          # Dockerfile pour construire le conteneur frontend
│
├── docker-compose.yml      # Configuration Docker pour orchestrer les conteneurs backend et frontend
├── .env                    # Fichier pour les variables d'environnement (DB path, secrets)
└── README.md               # Documentation du projet




git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/SHA-666/social-network.git
git push -u origin main