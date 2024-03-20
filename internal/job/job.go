package job

import "image"

type Options struct {
	InputDirectory  string
	OutputDirectory string
	Verbose         bool
	Scale           float64
}

type Job struct {
	Input  string
	File   image.Image
	Output string
}
