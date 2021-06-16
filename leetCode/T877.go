package leetCode

/*
*  @author liqiqiorz
*  @data 2021/6/16 23:07
 */

func stoneGame(piles []int) bool {
	var alex, li, time int
	for len(piles) > 0 {
		var stone, key int
		if piles[0] > piles[len(piles)-1] {
			stone = piles[0]
			key = 0
		} else {
			stone = piles[len(piles)-1]
			key = len(piles) - 1
		}
		piles = append(piles[:key], piles[key+1:]...)
		if time%2 == 0 {
			alex += stone
		} else {
			li += stone
		}
	}
	return alex > li
}
