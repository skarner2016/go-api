package area

import (
	"fmt"
	"testing"
)

func TestArea_GetArea(t *testing.T) {
	areas := NewArea().GetAreas()

	for areaName, area := range areas {
		fmt.Println(areaName, area)
	}
}

func TestArea_GetAreaList(t *testing.T) {
	areas := NewArea().GetAreaList()

	for areaName, area := range areas {
		fmt.Println(areaName, area)
	}
}