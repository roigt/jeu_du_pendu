# JEU DU PENDU (HANGMAN) en mode console

## Pré-requis
- Go 1.25
---
### Fonctionnement
Le pendu est un jeu consistant à trouver un mot en devinant quelles sont les lettres qui le composent.
---
### Organisation du projet
- main.go (point d'entrée)
- word.txt (fichier contenant les mots)
- **jeu_pendu ( /dictionary, /hangman, /test)**:
  - dictionary (repertoire): contient les fonctions de chargement du dictionnaire de mots et de choix aléatoire de mot.
  - hangman(repertoire): contient la logique de fonctionnement du jeu et la sortie console associée.
  - test (repertoire): contient les tests des différentes fonctionnalités du jeu.
---
### Lancement du jeu
`
go run .
`
---
### Lancement des tests
```
go test -v .\jeu_pendu\test\
```

