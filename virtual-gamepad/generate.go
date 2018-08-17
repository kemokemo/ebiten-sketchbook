package main

/* package images */
//go:generate file2byteslice -package=images -input=./_images/directional_button.png -output=./internal/images/directional_button.go -var=Directional_button_png
//go:generate file2byteslice -package=images -input=./_images/directional_pad.png -output=./internal/images/directional_pad.go -var=Directional_pad_png
