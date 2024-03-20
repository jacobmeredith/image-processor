package job

import (
	"image"
)

type Options struct {
	InputDirectory  string
	OutputDirectory string
	Verbose         bool
	Scale           float64
}

type Error struct {
	Step    string
	Message string
}

type Job struct {
	Input  string
	File   image.Image
	Output string
	Errors []Error
}
