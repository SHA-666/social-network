social-network/
├── backend
│   ├── controllers
│   │   └── auth.go
│   ├── middlewares
│   │   └── cors.go
│   ├── pkg
│   │   ├── db
│   │   │   └── migration
│   │   │       └── sql
│   │   └── models
│   ├── routes.go
│   ├── serveur.go
│   ├── services
│   ├── utils
│   └── websockets
│       └── ws.go
├── frontend
│   ├── public
│   └── src
│       ├── components
│       │   ├── auth
│       │   ├── layout
│       │   └── posts
│       ├── context
│       ├── hooks
│       ├── lib
│       ├── pages
│       │   ├── {api}
│       │   ├── _app.js
│       │   ├── _document.js
│       │   ├── index.js
│       │   ├── login.js
│       │   ├── profile.js
│       │   └── register.js
│       └── styles
│           └── globals.css
├── go.mod
├── main.go
└── readme.txt

## 1. Configuration initiale
- Installer Node.js et npm
- Créer un nouveau projet Next.js
- Configurer ESLint et Prettier

## 2. Mise en place de la structure de base
- Créer les dossiers principaux (components, styles, lib, etc.)
- Configurer le routage de base

## 3. Authentification
- Implémenter l'inscription et la connexion
- Configurer NextAuth.js pour la gestion des sessions

## 4. Profils utilisateurs
- Créer une page de profil
- Implémenter la modification du profil

## 5. Posts
- Créer, afficher et supprimer des posts
- Implémenter les commentaires

## 6. Followers
- Ajouter la fonctionnalité de suivre/ne plus suivre
- Afficher les followers et les following

## 7. Groupes
- Créer et gérer des groupes
- Implémenter les posts de groupe

## 8. Chat en temps réel
- Configurer WebSockets
- Implémenter le chat privé et de groupe

## 9. Notifications
- Mettre en place un système de notifications
- Implémenter les notifications en temps réel


