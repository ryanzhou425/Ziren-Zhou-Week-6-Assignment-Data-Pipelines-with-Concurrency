package main

import (
	"flag"
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"strings"
	"time"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				fmt.Printf("Can't load image %s: %v\n", p, err)
				continue
			}
			job := Job{
				InputPath: p,
				Image:     img,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			if err := imageprocessing.WriteImage(job.OutPath, job.Image); err != nil {
				fmt.Printf("Can't save image %s: %v\n", job.OutPath, err)
				out <- false
			} else {
				out <- true
			}
		}
		close(out)
	}()
	return out
}

func main() {
	useGoroutines := flag.Bool("goroutines", true, "Use goroutines for processing")
	flag.Parse()

	imagePaths := []string{
		"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	fmt.Printf("With goroutines: %v\n", *useGoroutines)
	startTime := time.Now()

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success")
		} else {
			fmt.Println("Failed")
		}
	}

	duration := time.Since(startTime)
	fmt.Printf("\nPipeline completed in %v\n", duration)
}
