package progressbar

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Progressbar struct{
	Filled string
	Total int
	Start int
	Spent int
	Size int
	PBar string
	Scale float64
	IsStop bool
}

func NewProgressBar() Progressbar{
	return Progressbar{
		Filled: "#",
		Total: 100,
		Size: 50,
		Start: 0,
		Spent: 0,
		PBar: "",
		Scale: 2.0,
		IsStop: false,
	}
}

func (pbar *Progressbar) Set_filled(filled string) error{
	if len(filled) > 1{
		return errors.New("filled size must be 1")
	}
	pbar.Filled = filled

	return nil
}

func (pbar *Progressbar) Set_total(total int){
	pbar.Total = total
}

func (pbar *Progressbar) Set_start(start int) error{
	if start > pbar.Total{
		return errors.New("start value must be less than total")
	}
	pbar.Start = start

	return nil
}

func (pbar *Progressbar) Set_size(size int){
	pbar.Size = size
	if size != 50{
		pbar.Scale = math.Round(100.0 / float64(pbar.Start) * 1000) / 1000
	}
}

func (pbar *Progressbar) Reset(){
	pbar.IsStop = false
	percent := math.Round(float64(pbar.Start)/ float64(pbar.Total) * 1000) / 10
	progress := int(math.Floor(percent/ pbar.Scale))
	pbar.PBar = fmt.Sprintf("\r[%v%v] %.1f%%", strings.Repeat(pbar.Filled, progress), strings.Repeat(" ", pbar.Size - progress), percent)
}

func (pbar *Progressbar) Update(value int){
	if !pbar.IsStop {
		pbar.Spent += value
		percent := math.Round(float64(pbar.Spent)/ float64(pbar.Total) * 1000) / 10
		progress := int(math.Floor(percent/ pbar.Scale))
		pbar.PBar = fmt.Sprintf("\r[%v%v] %.1f%%", strings.Repeat(pbar.Filled, progress), strings.Repeat(" ", pbar.Size - progress), percent)
	}else{
		fmt.Println("This progress bar has stopped")
	}
}

func (pbar Progressbar) Show(){
	fmt.Print(pbar.PBar)
}

func (pbar *Progressbar) Stop(){
	pbar.IsStop = true
	fmt.Print("\n")
}