package main

func main() {

}

func maxProfit(prices []int) int {
	var maxP int
	for left, right := 0, 1; right < len(prices); right++ {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			if profit > maxP {
				maxP = profit
			}
		} else {
			left = right
		}
	}

	return maxP
}
