package netsample

import (
	"context"

	"github.com/yandex/pandora/core"
)

type Aggregator interface {
	Start(context.Context) error
	Release(*Sample)
}

func WrapAggregator(a Aggregator) core.Aggregator { return &aggregatorWrapper{a} }

func UnwrapAggregator(a core.Aggregator) Aggregator {
	switch a := a.(type) {
	case *aggregatorWrapper:
		return a.Aggregator
	}
	return &aggregatorUnwrapper{a}
}

type aggregatorWrapper struct{ Aggregator }

func (a *aggregatorWrapper) Release(s core.Sample) { a.Aggregator.Release(s.(*Sample)) }

type aggregatorUnwrapper struct{ core.Aggregator }

func (a *aggregatorUnwrapper) Release(s *Sample) { a.Aggregator.Release(s) }
