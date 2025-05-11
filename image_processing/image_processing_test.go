package imageprocessing

import (
	"os"
	"testing"
)

func TestReadImage(t *testing.T) {
	img, err := ReadImage("../images/image1.jpeg")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if img == nil {
		t.Fatal("Expected a valid image, got nil")
	}
}

func TestWriteImage(t *testing.T) {
	img, err := ReadImage("../images/image1.jpeg")
	if err != nil {
		t.Fatalf("Can't read image: %v", err)
	}

	outputPath := "../images/output/test_output.jpeg"
	err = WriteImage(outputPath, img)
	if err != nil {
		t.Fatalf("Can't write image: %v", err)
	}

	// Check if file exists
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Fatalf("Expected output file %s to exist, but it does not", outputPath)
	}

	// Cleanup
	os.Remove(outputPath)
}

func BenchmarkPipeline(b *testing.B) {
	img, err := ReadImage("../images/image1.jpeg")
	if err != nil {
		b.Fatalf("Can't read image: %v", err)
	}

	for i := 0; i < b.N; i++ {
		resized := Resize(img)
		gray := Grayscale(resized)
		err := WriteImage("../images/output/benchmark_output.jpeg", gray)
		if err != nil {
			b.Fatalf("Can't write image: %v", err)
		}
	}
}
