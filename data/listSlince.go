package data

import (
	"github.com/aws/aws-sdk-go/service/lightsail"
	"sort"
)

type lightWrapper struct {
	instance []*lightsail.Instance
	by       func(p, q *lightsail.Instance) bool
}
type SortBy func(p, q *lightsail.Instance) bool

func (lw lightWrapper) Len() int {
	return len(lw.instance)
}

func (lw lightWrapper) Swap(i, j int) {
	lw.instance[i], lw.instance[j] = lw.instance[j], lw.instance[i]
}

func (lw lightWrapper) Less(i, j int) bool {
	return lw.by(lw.instance[i], lw.instance[j])
}

func SortInstance(instance []*lightsail.Instance, by SortBy) {
	sort.Sort(lightWrapper{instance, by})

}
