# Project Overview
This project is a Go language-based concurrent image processing pipeline that supports image loading, resizing, grayscale conversion, and saving. The project provides both synchronous and asynchronous processing modes, and includes comprehensive unit tests and benchmark tests.

---

# Program Features
The images used in this project are from Google Images, specifically campus photos of Northwestern University. The images are in JPEG format with a 1:1 aspect ratio, meaning they are square images.
•	Image Loading: Converts images to Go's image.Image objects.
•	Image Resizing: Uses the [github.com/nfnt/resize](github.com/nfnt/resize) library to resize images to 500x500 pixels.
•	Grayscale Conversion: Converts loaded color images to grayscale.
•	Image Saving: Saves the processed images to specified output files.
•	Concurrent and Non-Concurrent Modes: Supports concurrent processing as well as non-concurrent processing.
•	Benchmark Testing: Includes benchmark tests to evaluate the performance of both concurrent and non-concurrent processing modes.


---

# How to Run Applications
To start the application, use the following commands:
•	Concurrent Mode: go run main.go -goroutines=true
•	Non-Concurrent Mode: go run main.go -goroutines=false


---

# How to Run Unit Tests
To run the unit tests, use the following command:
•	go test ./image_processing

---

# How to Run Benchmark Tests

To run the performance benchmark tests, use the following command：
•	go test ./image_processing -bench=.

---

# Compare processing times with and without goroutines
•	With goroutines: 226.2446ms
•	Without goroutines: 374.9332ms

This demonstrates that goroutines fully utilize the concurrency capabilities of multi-core CPUs, reducing the time spent on blocking and idle waiting. As a result, using goroutines can significantly increase the throughput of data pipelines, making them an effective way to optimize the performance of Go applications.


---

# Use of AI Assistants

•	Asked about the purpose of adding unit tests and how to add them
•	Asked how to create benchmark test files
•	Asked what a deadlock is and how to resolve it
•	Asked which commands to use for more detailed performance comparisons
