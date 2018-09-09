package main

/* package images */
//go:generate file2byteslice -package=images -input=./_images/directional_button.png -output=./internal/images/directional_button.go -var=Directional_button_png
//go:generate file2byteslice -package=images -input=./_images/directional_pad.png -output=./internal/images/directional_pad.go -var=Directional_pad_png
//go:generate file2byteslice -package=images -input=./_images/a_button.png -output=./internal/images/a_button.go -var=A_button_png
//go:generate file2byteslice -package=images -input=./_images/b_button.png -output=./internal/images/b_button.go -var=B_button_png
//go:generate file2byteslice -package=images -input=./_images/fighter.png -output=./internal/images/fighter.go -var=Fighter_png
//go:generate file2byteslice -package=images -input=./_images/another_fighter.png -output=./internal/images/another_fighter.go -var=Another_fighter_png
//go:generate file2byteslice -package=images -input=./_images/bullet.png -output=./internal/images/bullet.go -var=Bullet_png
//go:generate file2byteslice -package=images -input=./_images/laser.png -output=./internal/images/laser.go -var=Laser_png
