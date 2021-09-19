package change

import (
	"fmt"
	"sort"
)

func Change(coins []int, sum int) ([]int, error) {
	if sum == 0 {
		return []int{}, nil
	}
	if len(coins) == 0 {
		return nil, fmt.Errorf("no coins")
	}
	sort.Slice(coins, func(i, j int) bool { return coins[j] < coins[i] })
	if sum < coins[len(coins)-1] {
		return nil, fmt.Errorf("can't change for a value smaller than target")
	}

	cached := map[int][]int{}
	for i := 1; i < sum+1; i++ {
		minCoins(cached, i, coins)
	}
	result, ok := cached[sum]
	if ok {
		return result, nil
	}
	return nil, fmt.Errorf("no combination found")
}

func minCoins(cachedAmounts map[int][]int, amount int, coins []int) {
	result := []int{}
	for _, coin := range coins {
		if coin <= amount {
			result = append(result, coin)
		}
	}
	changes := make([][]int, len(result))
	for i, coin := range result {
		if cc, ok := cachedAmounts[amount-coin]; ok {
			cc = append(cc, coin)
			changes[i] = cc
		} else {
			changes[i] = []int{coin}
		}
	}
	var minChanges []int
	for _, ch := range changes {
		if minChanges == nil && sum(ch, amount) {
			minChanges = ch
		} else if len(ch) < len(minChanges) && sum(ch, amount) {
			minChanges = ch
		}
	}
	if minChanges != nil {
		sort.Ints(minChanges)
		cachedAmounts[amount] = minChanges
	}
}

func sum(coins []int, amount int) bool {
	total := 0
	for _, coin := range coins {
		total += coin
	}
	return total == amount
}
