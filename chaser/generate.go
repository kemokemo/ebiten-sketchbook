package main

/* package images */
//go:generate file2byteslice -package=images -input=./_assets/images/enemy.png -output=./assets/images/enemy.go -var=EnemyImage
//go:generate file2byteslice -package=images -input=./_assets/images/player.png -output=./assets/images/player.go -var=PlayerImage
