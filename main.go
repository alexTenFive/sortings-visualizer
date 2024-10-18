package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand"
	"sort"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const winHeight = 1080
const winWidth = 1920

const shapeWidth = 10
const shapeHeightRatio = 5
const shapePadding = 1

// IntSlice is a type that implements the sort.Interface for a slice of integers.
type IntSlice struct {
	values        []int
	informChannel chan [2]int
}

// Len returns the number of elements in the collection.
func (s IntSlice) Len() int {
	return len(s.values)
}

// Less reports whether the element at index i should sort before the element at index j.
func (s IntSlice) Less(i, j int) bool {
	return s.values[i] < s.values[j]
}

// Swap swaps the elements at indexes i and j.
func (s IntSlice) Swap(i, j int) {
	//fmt.Println("comp function called", randomList)
	s.values[i], s.values[j] = s.values[j], s.values[i]
	s.informChannel <- [2]int{s.values[i], s.values[j]}
}

// BubbleSort implements the bubble sort algorithm.
func (s *IntSlice) BubbleSort() {
	n := len(s.values)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s.Less(j+1, j) { // Using Less to compare values
				s.Swap(j, j+1) // Call Swap to swap elements
			}
		}
	}
}

func (s *IntSlice) SelectionSort() {
	n := len(s.values)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if s.values[j] < s.values[minIdx] {
				minIdx = j
			}
		}
		s.Swap(i, minIdx)
	}
}

func (s *IntSlice) InsertionSort() {
	n := len(s.values)
	for i := 1; i < n; i++ {
		key := s.values[i]
		j := i - 1
		for j >= 0 && s.values[j] > key {
			s.Swap(j, j+1)
			j--
		}
	}
}

func (s *IntSlice) ShellSort() {
	n := len(s.values)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := s.values[i]
			j := i
			for j >= gap && s.values[j-gap] > temp {
				s.Swap(j, j-gap)
				j -= gap
			}
		}
	}
}

func (s *IntSlice) QuickSort() {
	s.quickSort(0, len(s.values)-1)
}
func (s *IntSlice) quickSort(low, high int) {
	if low < high {
		pi := s.partition(low, high)
		s.quickSort(low, pi-1)
		s.quickSort(pi+1, high)
	}
}
func (s *IntSlice) partition(low, high int) int {
	pivot := s.values[high]
	i := low - 1
	for j := low; j < high; j++ {
		if s.values[j] < pivot {
			i++
			s.Swap(i, j)
		}
	}
	s.Swap(i+1, high)
	return i + 1
}

func (s *IntSlice) HeapSort() {
	n := len(s.values)
	for i := n/2 - 1; i >= 0; i-- {
		s.heapify(n, i)
	}
	for i := n - 1; i > 0; i-- {
		s.Swap(0, i)
		s.heapify(i, 0)
	}
}
func (s *IntSlice) heapify(n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && s.values[left] > s.values[largest] {
		largest = left
	}
	if right < n && s.values[right] > s.values[largest] {
		largest = right
	}
	if largest != i {
		s.Swap(i, largest)
		s.heapify(n, largest)
	}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sorting Visualization Program",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	randomList := IntSlice{
		values:        make([]int, *itemsCount),
		informChannel: make(chan [2]int),
	}

	for i := range randomList.values {
		randomList.values[i] = i + 1 // Generates a random integer between 0 and 100
	}
	rand.Shuffle(len(randomList.values), func(i, j int) {
		randomList.values[i], randomList.values[j] = randomList.values[j], randomList.values[i]
	})

	var (
		tn       time.Time
		timeStop = make(chan struct{})
	)
	go func() {
		tn = time.Now()
		switch *sortType {
		case "quick":
			randomList.QuickSort() // default sort
		case "bubble":
			randomList.BubbleSort()
		case "selection":
			randomList.SelectionSort()
		case "insertion":
			randomList.InsertionSort()
		case "heap":
			randomList.HeapSort()
		case "shell":
			randomList.ShellSort()
		case "default":
			sort.Sort(randomList)
		}
		close(randomList.informChannel)
		close(timeStop)
	}()

	shaper := newShaper(imdraw.New(nil), 5, 5, colornames.White, colornames.Darkcyan)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(5, winHeight-20), basicAtlas)

	var timeSince time.Duration

	for !win.Closed() {
		win.Clear(colornames.Black)

		targets := <-randomList.informChannel
		for _, v := range randomList.values {
			if targets[0] == v || targets[1] == v {
				shaper.drawShape(v, true)
				continue
			}
			shaper.drawShape(v, false)
		}

		shaper.draw(win)

		switch *sortType {
		case "quick":
			fmt.Fprintln(basicTxt, "Quick sort")
		case "bubble":
			fmt.Fprintln(basicTxt, "Bubble sort")
		case "heap":
			fmt.Fprintln(basicTxt, "Heap sort")
		case "selection":
			fmt.Fprintln(basicTxt, "Selection sort")
		case "insertion":
			fmt.Fprintln(basicTxt, "Insertion sort")
		case "shell":
			fmt.Fprintln(basicTxt, "Shell sort")
		case "default":
			fmt.Fprintln(basicTxt, "Default sort")
		}

		select {
		case <-timeStop:
		default:
			timeSince = time.Since(tn)
		}
		fmt.Fprintf(basicTxt, "Time: %.2fs\n", timeSince.Seconds())

		basicTxt.Draw(win, pixel.IM)
		basicTxt.Clear()

		win.Update()
		time.Sleep(time.Second / 60)

	}
}

type shaper struct {
	imd                          *imdraw.IMDraw
	shapeColor, shapeBorderColor color.Color
	startPositionX               float64
	startPositionY               float64
}

func newShaper(imd *imdraw.IMDraw, sPosX, sPosY float64, shapeColor, shapeBorderColor color.Color) *shaper {
	return &shaper{
		imd:              imd,
		shapeColor:       shapeColor,
		shapeBorderColor: shapeBorderColor,
		startPositionX:   sPosX,
		startPositionY:   sPosY,
	}
}

func (s *shaper) draw(win *pixelgl.Window) {
	s.imd.Draw(win)
	s.startPositionX = 5
	s.imd.Clear()
}

func (s *shaper) drawShape(value int, targeted bool) {
	s.imd.Color = s.shapeColor
	if targeted {
		s.imd.Color = s.shapeBorderColor
	}
	s.imd.Push(pixel.V(s.startPositionX, s.startPositionY), pixel.V(s.startPositionX+float64(shapeWidth), float64(value)*shapeHeightRatio))
	s.imd.Rectangle(0)

	s.imd.Color = s.shapeBorderColor
	s.imd.EndShape = imdraw.NoEndShape
	s.imd.Push(pixel.V(s.startPositionX, s.startPositionY), pixel.V(s.startPositionX, float64(value)*shapeHeightRatio))
	s.imd.Push(pixel.V(s.startPositionX, float64(value)*shapeHeightRatio), pixel.V(s.startPositionX+float64(shapeWidth), float64(value)*shapeHeightRatio))
	s.imd.Push(pixel.V(s.startPositionX+float64(shapeWidth), float64(value)*shapeHeightRatio), pixel.V(s.startPositionX+float64(shapeWidth), s.startPositionY))
	s.imd.Push(pixel.V(s.startPositionX, s.startPositionY), pixel.V(s.startPositionX+float64(shapeWidth), s.startPositionY))
	s.imd.Line(1)

	if targeted {
		s.imd.Color = colornames.Red
		s.imd.Push(pixel.V(s.startPositionX+shapeWidth/2, float64(value)*shapeHeightRatio+5), pixel.V(s.startPositionX+shapeWidth/2, float64(value)*shapeHeightRatio+15))
		s.imd.Line(2)
		s.imd.Push(pixel.V(s.startPositionX, float64(value)*shapeHeightRatio+8), pixel.V(s.startPositionX+shapeWidth/2, float64(value)*shapeHeightRatio+5))
		s.imd.Push(pixel.V(s.startPositionX+shapeWidth, float64(value)*shapeHeightRatio+8), pixel.V(s.startPositionX+shapeWidth/2, float64(value)*shapeHeightRatio+5))
		s.imd.Line(2)
	}

	s.startPositionX = s.startPositionX + shapePadding + shapeWidth
}

var sortType = flag.String("sort", "default", "[quick, bubble, selection, insertion, heap, shell, default]")
var itemsCount = flag.Int("items", 100, "Items count")

func main() {
	flag.Parse()
	pixelgl.Run(run)
}
