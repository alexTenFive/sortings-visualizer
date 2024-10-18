# Sorting Visualization Program

This program visualizes various sorting algorithms using the [Pixel](https://github.com/faiface/pixel) graphics library in Go. It allows users to see the step-by-step execution of different sorting algorithms on randomly generated lists of integers.

## Features

- Implements the following sorting algorithms:
  - Bubble Sort
  - Selection Sort
  - Insertion Sort
  - Shell Sort
  - Quick Sort
  - Heap Sort
- Visual representation of sorting steps.
- Adjustable number of items to sort.
- Command-line options to select the sorting algorithm.

## Requirements

- Go 1.15 or later
- Pixel library for graphics
- Go modules

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/sorting-visualization.git
   cd sorting-visualization
   ```

2. Install the required dependencies:
   ```bash
   go mod tidy
   ```

## Usage

Run the program with the following command:

```bash
go run main.go -sort=<algorithm> -items=<count>
```

### Command-Line Flags

- `-sort`: Choose the sorting algorithm to use. Available options:
  - `quick` (default)
  - `bubble`
  - `selection`
  - `insertion`
  - `heap`
  - `shell`
  - `default` (Go's built-in sort)
  
- `-items`: Set the number of items to sort (default is 100).

### Example

To run the program using Bubble Sort with 50 items:

```bash
go run main.go -sort=bubble -items=50
```

## Visuals

The program creates a window that displays the sorting process. Each integer is represented as a bar, and the color changes to indicate which elements are currently being compared and swapped. The elapsed time for sorting is displayed at the bottom of the window.

## Code Structure

- `IntSlice`: A type that implements the `sort.Interface` for sorting integers.
- Sorting algorithms:
  - `BubbleSort`: Implements bubble sort.
  - `SelectionSort`: Implements selection sort.
  - `InsertionSort`: Implements insertion sort.
  - `ShellSort`: Implements shell sort.
  - `QuickSort`: Implements quick sort.
  - `HeapSort`: Implements heap sort.
- `shaper`: A utility for drawing the visual representation of the sorting process.

## Contributing

Feel free to fork the repository and submit pull requests for improvements or additional features!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
