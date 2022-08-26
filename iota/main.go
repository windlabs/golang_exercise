package main

type Plans uint8

const (
	PlansProfession Plans = iota +1
	PlansAdvanced
	PlansExalted
	PlansUltimate
)
var PlansNameMap = map[Plans]string{
	PlansProfession:   "专业版",
	PlansAdvanced:     "高级版",
	PlansExalted:      "尊享版",
	PlansUltimate:     "旗舰版",
}

func (r Plans) Name() (rv string) {
	if rv, ok := PlansNameMap[r]; ok {
		return rv
	}

	return 
}

