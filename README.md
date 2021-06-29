
# Projet-Forum

### :one: L'ambition du projet : 

Le but de ce projet est de créer un forum dans le sens général, c’est-à-dire un site qui permet de créer des posts qui soient du texte et/ou image. Chaque poste une ou plusieurs catégories de plus les utilisateurs peuvent y réagir en postant des commentaires et via des votes (likes, dislikes). En outre les posts et les commentaires seront accessibles par tout le monde, mais participer nécessitera de créer un compte utilisateur. Enfin on pourra filtrer les posts par catégorie, par posts qu'on a soi-même liké ou disliké.


### :two: Language utiliser :

<img alt="HTML5" src="https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white%22/%3E"/> 

``Hypertext Markup Language``

<img alt="js" src="https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=F7DF1E"/>

``JavaScript``

<img src="https://img.shields.io/badge/CSS-239120?&style=for-the-badge&logo=css3&logoColor=white"/>

``Cascading Style Sheets``

<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
"/>
``golang``

<img src="https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white"/> 

``SQLite``

### :three: Pré-requis :
Pour acquérire le projet en .zip [cliquez ici !](https://github.com/Luke859/Projet-Forum/archive/refs/heads/main.zip)

Tout d'abord installer les librairies ``sqlite3, bcrypt, uuid`` pour cela direction donc  dans le ***terminal*** de ***Visual Studio Code*** et utiliser la commande suivante ``go get`` :
* ``github.com/mattn/go-sqlite3``
* ``golang.org/x/crypto/bcrypt``
* ``github.com/google/uuid``


Le serveur est fin prêt à être lancé pour cela utiliser la commande `` go run serveur`` dans le ***terminal*** ensuite aller dans votre ***navigateur web*** et inscriver ``local host :8080`` dans la barre de navigation.
### :four: Environnement GIT:
Tout d'abord notre répertoire ``Projet-Forum`` contient 7 branches. Chaque membre du projet possède une branche individuelle, cela permet d'éviter des conflits entre les parties de codes de chacun. Elle permet aussi aux membres de travailler et tester son code.
 
<img width="300" alt="GITPERSO" src="https://user-images.githubusercontent.com/72868466/123279556-3f192800-d508-11eb-98cc-4b93b07ff61f.png">
<br>

Ensuite notre répertoire contient une branche ``dev`` et une branche ``main`` et ``reviews``. Dans notre cas, la ``dev`` permet de regrouper tous les bouts de code de chaque membre et de régler les conflits uniquement sur cette branche. Contrairement à la ``dev`` la branche ``main`` permet de récupérer tous les codes dont la version finale finale du projet. Enfin la branche ``reviews`` avait pour but de faire un point en milieu de projet.

 <img width="300" alt="GITDEVREV" src="https://user-images.githubusercontent.com/72868466/123381824-84ce0300-d591-11eb-957f-06b66605fa4c.png">
<br>

### :five: Architecture du projet :

Pour l'architecture du projet, nous avons créer 3 grands dossiers pour la répartition des fichiers :    
* BDD → contient la Base De Donnée (les fichiers ".db") et son schéma.     
* src → contient les fichiers en golang répartis en 2 sous-dossiers qui sont Base De Donnée pour tous les fichiers lier à la récupération de données de la Base De Donnée et, le sous-dossier Accueil qui contient tous les scripts nécessaire au bon fonctionnement de forum.     
* static → contient le CSS, les fichiers HTML et les fichiers Java Scripts.     

De plus, nous avons le fichier "server.go" à la racine du projet.   
Pour mieux comprendre, voici une capture d'écran de la disposition des fichiers ci-dessous:      

<img src="https://cdn.discordapp.com/attachments/754974960845324309/859418457091211274/Screenshot_from_2021-06-29_14-40-12.png">

### :six: Membres du groupes :

* **Pierric Come** _alias_ [@lancelot260](https://github.com/lancelot260)

* **Luke Jones** _alias_ [@Luke859](https://github.com/Luke859)

* **Florian Dagnas** _alias_ [@Flodagnas](https://github.com/Flodagnas)

* **Nathy Mellal** _alias_ [@nathymellal](https://github.com/nathymellal)



