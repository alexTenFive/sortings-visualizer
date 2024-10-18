# Sorting Visualization Program

This is a Go program that visualizes different sorting algorithms using the Pixel library. The visualization shows the process of sorting a list of integers using various sorting algorithms.

## Features

- Supports multiple sorting algorithms: Quick Sort, Bubble Sort, Selection Sort, Insertion Sort, Heap Sort, and Shell Sort.
- Visualizes the sorting process in real-time.
- Displays the elapsed time taken by the sorting algorithm.

## Prerequisites

- Go 1.16 or higher
- Pixel library
- x/image package

## Installation

1. **Install Go**: If you don't have Go installed, download and install it from [here](https://golang.org/dl/).

2. **Get the required packages**:
    ```sh
    go get github.com/faiface/pixel
    go get golang.org/x/image/colornames
    go get golang.org/x/image/font/basicfont
    ```

## Usage

1. **Clone the repository**:
    ```sh
    git clone <repository_url>
    cd <repository_directory>
    ```

2. **Build and run the program**:
    ```sh
    go run main.go --sort=<algorithm> --items=<number_of_items>
    ```

    Replace `<algorithm>` with one of the following options: `quick`, `bubble`, `selection`, `insertion`, `heap`, `shell`, or `default`.

    Replace `<number_of_items>` with the number of items to sort.

## Example

To run the program with Quick Sort and 150 items:
```sh
go run main.go --sort=quick --items=150
