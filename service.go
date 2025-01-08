package main

import "context"

type Pricefetcher interface{
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	
}

func MockPriceFetcher(ctx )