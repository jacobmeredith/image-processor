package job

import (
	"image"
)

type Options struct {
	InputDirectory  string
	OutputDirectory string
	Verbose         bool
	Scale           float64
	Quality         int
	Crop            CropOptions
}

type CropOptions struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Error struct {
	Step    string
	Message string
}

type Job struct {
	Input  string
	File   image.Image
	Output string
	Cancel bool
	Errors []Error
}
