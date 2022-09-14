package model

import "sort"

type TireAnalyticsResponse struct {
	Count int     `json:"count" example:"35" validate:"required"`
	Sizes []*Size `json:"list" validate:"required"`
}

type Size struct {
	Rank    int     `json:"rank" example:"1"`
	Size    string  `json:"size" example:"175/70 R14"`
	Index   float32 `json:"index" example:"24.344"`
	Percent float32 `json:"percent" example:"7.89"`
	Plates  struct {
		Count int      `json:"count" example:"3"`
		List  []string `json:"list" example:"A412AY142,C109HA142,K093YE70"`
	} `json:"plates"`
	Cars struct {
		Count int      `json:"count" example:"2"`
		List  []string `json:"list" example:"Nissan Almera 2011,ВАЗ 2101-2107 2005"`
	} `json:"cars"`
}

// FirstOrCreateSize - получить или создать размер size в списке t.sizes
func (t *TireAnalyticsResponse) FirstOrCreateSize(findSize string) *Size {
	for _, size := range t.Sizes {
		if size.Size == findSize {
			return size
		}
	}
	newSize := &Size{
		Size: findSize,
	}
	t.Sizes = append(t.Sizes, newSize)
	t.Count++
	return newSize
}

// CaclSizesParams - Отсортироваить Sizes по index и подсчитать Rank, Percent
func (t *TireAnalyticsResponse) CaclSizesParams() {
	sort.Slice(t.Sizes, func(i, j int) bool {
		return t.Sizes[i].Index > t.Sizes[j].Index
	})
	var indexSum float32
	for _, size := range t.Sizes {
		indexSum += size.Index
	}
	if indexSum == 0 {
		return
	}
	for i, size := range t.Sizes {
		size.Rank = i+1
		size.Percent = (100 * size.Index) / indexSum
	}
}
