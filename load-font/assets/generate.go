package assets

// This file converts all assets to the golang file (*.go) with 'go generate' command.

//go:generate file2byteslice -package=fonts -input=./fonts/x8y12pxTheStrongGamer.ttf -output=./fonts/thestronggamer.go -var=TheStrongGamer_ttf
