// Code generated by "callbackgen -type Stream -interface"; DO NOT EDIT.

package kucoin

import (
	"github.com/c9s/bbgo/pkg/exchange/kucoin/kucoinapi"
)

func (s *Stream) OnCandleEvent(cb func(c *kucoinapi.WebSocketCandle)) {
	s.candleEventCallbacks = append(s.candleEventCallbacks, cb)
}

func (s *Stream) EmitCandleEvent(c *kucoinapi.WebSocketCandle) {
	for _, cb := range s.candleEventCallbacks {
		cb(c)
	}
}

func (s *Stream) OnOrderBookL2Event(cb func(c *kucoinapi.WebSocketOrderBookL2)) {
	s.orderBookL2EventCallbacks = append(s.orderBookL2EventCallbacks, cb)
}

func (s *Stream) EmitOrderBookL2Event(c *kucoinapi.WebSocketOrderBookL2) {
	for _, cb := range s.orderBookL2EventCallbacks {
		cb(c)
	}
}

func (s *Stream) OnTickerEvent(cb func(c *kucoinapi.WebSocketTicker)) {
	s.tickerEventCallbacks = append(s.tickerEventCallbacks, cb)
}

func (s *Stream) EmitTickerEvent(c *kucoinapi.WebSocketTicker) {
	for _, cb := range s.tickerEventCallbacks {
		cb(c)
	}
}

func (s *Stream) OnAccountBalanceEvent(cb func(c *kucoinapi.WebSocketAccountBalance)) {
	s.accountBalanceEventCallbacks = append(s.accountBalanceEventCallbacks, cb)
}

func (s *Stream) EmitAccountBalanceEvent(c *kucoinapi.WebSocketAccountBalance) {
	for _, cb := range s.accountBalanceEventCallbacks {
		cb(c)
	}
}

func (s *Stream) OnPrivateOrderEvent(cb func(c *kucoinapi.WebSocketPrivateOrder)) {
	s.privateOrderEventCallbacks = append(s.privateOrderEventCallbacks, cb)
}

func (s *Stream) EmitPrivateOrderEvent(c *kucoinapi.WebSocketPrivateOrder) {
	for _, cb := range s.privateOrderEventCallbacks {
		cb(c)
	}
}

type StreamEventHub interface {
	OnCandleEvent(cb func(c *kucoinapi.WebSocketCandle))

	OnOrderBookL2Event(cb func(c *kucoinapi.WebSocketOrderBookL2))

	OnTickerEvent(cb func(c *kucoinapi.WebSocketTicker))

	OnAccountBalanceEvent(cb func(c *kucoinapi.WebSocketAccountBalance))

	OnPrivateOrderEvent(cb func(c *kucoinapi.WebSocketPrivateOrder))
}