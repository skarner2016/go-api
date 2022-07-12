package area

import (
	"fmt"
	"strings"
)

type AreaName string
type AreaCode int64

type Area struct {
	Name      AreaName
	TransName string
	AreaCode  AreaCode
	Icon      string
}

const (
	China    AreaName = "Chinese"
	HongKong AreaName = "HongKong"
	USA      AreaName = "USA"
)

func NewArea() *Area {
	return &Area{}
}

func (a *Area) GetAreas() map[AreaName]Area {
	fmt.Println("icon:", a.getIconByName(China))

	return map[AreaName]Area{
		China: {
			Name:      China,
			TransName: a.getTranslationNameByName(China),
			AreaCode:  86,
			Icon:      a.getIconByName(China),
		},
		HongKong: {
			Name:      HongKong,
			TransName: a.getTranslationNameByName(HongKong),
			AreaCode:  852,
			Icon:      a.getIconByName(HongKong),
		},
		USA: {
			Name:      USA,
			TransName: a.getTranslationNameByName(USA),
			AreaCode:  1,
			Icon:      a.getIconByName(USA),
		},
	}
}

func (a *Area) GetAreaList() []Area {
	areaList := make([]Area, 0)
	areas := a.GetAreas()
	for _, area := range areas {
		areaList = append(areaList, area)
	}

	return areaList
}

// 获取icon （地区/国旗）
func (a *Area) getIconByName(name AreaName) string {
	// TODO 配置oss
	nameString := string(name)
	ossDomain := "https://www.go-api.com/area/%s.png"

	return fmt.Sprintf(ossDomain, strings.ToLower(nameString))
}

// 翻译后的名称
func (a *Area) getTranslationNameByName(name AreaName) string {
	// TODO 多语言
	return string(name)
}
