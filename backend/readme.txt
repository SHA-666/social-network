I. Gestion des utilisateurs

1-Inscription et Authentification :

Règle métier : Un utilisateur doit fournir une adresse email valide et un mot de passe pour s'inscrire.
Règle métier : Un utilisateur ne peut pas avoir plusieurs comptes avec la même adresse email.
Règle métier : Les mots de passe doivent être stockés de manière sécurisée (par exemple, hachés).

2-Règles d'authentification et sécurité :

Règle métier : Après plusieurs tentatives de connexion échouées, l’utilisateur est bloqué temporairement pour éviter les attaques par force brute.
Règle métier : Les sessions d'utilisateur expirent après un certain temps d'inactivité pour renforcer la sécurité.
Règle métier : Lors de l'inscription, les utilisateurs doivent choisir un nom d'utilisateur unique, et il ne peut y avoir qu’un seul nom d’utilisateur par compte.

II. Gestion des profils utilisateurs

1-Création et modification du profil :

Règle métier : Un utilisateur peut ajouter une photo de profil et une description personnelle.
Règle métier : Un utilisateur peut modifier ses informations personnelles, mais certaines données (comme l'email) sont protégées.
Règle métier : Un utilisateur peut choisir de rendre son profil public ou privé. Si public, certaines informations seront visibles par tous les utilisateurs, sinon elles seront visibles uniquement par ses amis.

III. Interaction avec les autres utilisateurs
Ajouter des amis / suivre des utilisateurs :

Règle métier : Un utilisateur peut envoyer une demande d’amitié à un autre utilisateur, mais cette demande peut être acceptée ou rejetée.
Règle métier : Un utilisateur peut suivre un autre utilisateur sans être ami, ce qui permet de voir ses publications sans lui envoyer une demande d’amitié.
Règle métier : Lorsqu’un utilisateur accepte une demande d’amitié, il devient un ami, et les deux utilisateurs peuvent partager du contenu de manière plus privilégiée (par exemple, en ajustant les paramètres de visibilité des publications).

Règle métier : Si un utilisateur bloque un autre utilisateur, celui-ci ne pourra plus voir son profil ou ses publications. De plus, les utilisateurs bloqués ne peuvent pas envoyer de messages.

IV. Publication de contenu

1-Création de posts (publications) :

Règle métier : Un utilisateur peut publier du texte, des photos, des vidéos, etc., sur son mur ou dans un groupe.
Règle métier : Un utilisateur peut taguer d'autres utilisateurs dans ses publications. Cela affecte qui voit la publication selon les paramètres de confidentialité.
Règle métier : Les publications peuvent être privées, visibles par les amis, ou publiques. L’utilisateur choisit cette visibilité au moment de la publication.

2-Gestion des commentaires et réactions :

Règle métier : Un utilisateur peut commenter une publication et réagir à celle-ci (par exemple, "j’aime", "j’adore", etc.).
Règle métier : Un utilisateur peut supprimer ou modifier ses propres commentaires, mais ne peut pas supprimer les commentaires des autres.

5. Fil d’actualités (Newsfeed)
Tri des publications :
Règle métier : Le fil d’actualités d’un utilisateur affiche les publications de ses amis et des pages qu'il suit. Cependant, un algorithme trie et priorise ces publications en fonction de plusieurs critères comme la pertinence, l’engagement ou les préférences de l’utilisateur.
Règle métier : Un utilisateur peut choisir de masquer ou de signaler une publication qu'il trouve inappropriée ou de désabonner un ami ou une page pour ne plus voir ses publications.
6. Gestion des groupes et événements
Création de groupes :
Règle métier : Un utilisateur peut créer un groupe avec des paramètres de confidentialité (public, privé, fermé).
Règle métier : Les groupes peuvent avoir des règles spécifiques (par exemple, qui peut publier, qui peut ajouter des membres).
Gestion des événements :
Règle métier : Un utilisateur peut créer un événement (par exemple, une fête, une réunion) et inviter des amis.
Règle métier : Un utilisateur peut accepter ou refuser une invitation à un événement, ou marquer qu’il participe.
7. Notifications
Règle métier : Un utilisateur recevra des notifications pour des événements importants comme une nouvelle demande d'amitié, un message privé, une mention dans un commentaire, ou un événement à venir.

Règle métier : Les notifications peuvent être ajustées dans les paramètres pour être activées/désactivées pour différentes interactions.

8. Modération du contenu
Gestion des contenus inappropriés :
Règle métier : Si un utilisateur signale une publication ou un commentaire comme inapproprié, un système de modération vérifie si le contenu viole les règles du site.
Règle métier : Si le contenu est jugé inapproprié, il peut être supprimé ou un avertissement peut être envoyé à l’utilisateur.
Règle métier : En cas de comportement répété ou grave, un utilisateur peut être banni temporairement ou permanent.
9. Gestion de la confidentialité et des paramètres
Paramètres de confidentialité :
Règle métier : L’utilisateur peut définir qui peut voir ses informations (profil, publications, amis).
Règle métier : Les utilisateurs peuvent choisir de rendre certaines informations visibles à tous (par exemple, la photo de profil), mais d’autres (comme la liste d’amis) uniquement visibles par les amis.
Règle métier : Un utilisateur peut supprimer son compte ou désactiver son compte, et ses informations seront effacées ou rendues inaccessibles selon la politique de confidentialité du site.


social-network-backend/
├── cmd/
│   └── server/
│       └── main.go                 # Point d'entrée
├── internal/
│   ├── domain/                     # 🔵 DOMAIN LAYER
│   │   ├── entities/
│   │   ├── repositories/
│   │   ├── services/
│   │   └── usecases/
│   ├── application/                # 🟡 APPLICATION LAYER
│   │   ├── dto/
│   │   └── validators/
│   ├── infrastructure/             # 🔴 INFRASTRUCTURE LAYER
│   │   ├── database/
│   │   ├── config/
│   │   └── external/
│   └── presentation/               # 🟢 PRESENTATION LAYER
│       ├── controllers/
│       ├── routes/
│       ├── middleware/
│       └── server/
├── pkg/                           # Code réutilisable
│   ├── logger/
│   ├── response/
│   └── errors/
├── configs/
├── docs/
├── go.mod
├── go.sum
├── Makefile
└── README.md






tree -d -I "*.log|*.tmp|node_modules"
.
├── application
│   ├── dto  
│   └── usecases
├── cmd       ----> main.go file
│   └── tls 
├── core      -----> domain layer (principale layer)
│   ├── entities -> types
│   └── repositories -> declaration only of interfaces
├── delivery
│   ├── http
│   │   ├── handlers
│   │   └── middlewares 
│   └── ws 
├── infrastructure -> this layer handle all dependecie of sql and 
│   ├── db 
│   └── repositories -> implemente the interfaces defined in core layer
└── internal -> utilitaires
    ├── app
    └── assets
        └── uploads
            ├── cover_pic
            ├── posts
            └── profile_pic

            



├── application                    -----> layer application(depend on core layer)
│   ├── dto                        
│   └── usecases                   -> Implementation of app service (ex:creatpost..) (NO Interaction WHIT DB OR .. (only define how))
├── cmd                            -----> main.go comande
│   └── tls                        
├── core                           -----> layer core -> principale (no dependecie)
│   ├── entities                   -> (types/structs)
│   └── repositories               -> Interfaces des repositories (only declaration)
├── delivery                       -----> layer delivery
│   ├── http                       -> API REST
│   │   ├── handlers               -> handler HTTP
│   │   └── middlewares            -> Middlewares (auth, logging, CORS)
│   └── ws                         -> WebSocket handlers - hub setup
├── infrastructure                 -----> layer Infrastructure 
│   ├── db                         -> Config db
│   └── repositories               -> (implemente teh interfaces declared in the core) 
└── internal                       -----> Utilitaires internes
    ├── app                        
    └── assets                     
        └── uploads                -> Fichiers uploadés cote client
            ├── cover_pic          
            ├── posts              
            └── profile_pic        
                                        //celan archi is not aboute folder structure is 
                                        //aboute layer dependecie
                                        //uncle bob golden rule :
                                            Infrastructure ──→ Core ←── Application
                                                                ↑
                                                             Delivery

                                        blog post original "The Clean Architecture"  2012 sur blog Stack OverflowE4developer
                                        livre "Clean Architecture" 2014 ed                   
