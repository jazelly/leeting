package distanceK

import (
	"math"
)

// Beats 7.5% time and 10% space

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	if k == 0 {
		return []int{target.Val}
	}
	nvsPtr, targetIndex, targetD := transformToArray(root, target, k)
	return findWithArray(*nvsPtr, k, targetIndex, targetD)
}

func transformToArray(root *TreeNode, target *TreeNode, k int) (*[]int, int, int) {
	layer := []*TreeNode{root}
	d := 1
	nextLine := []int{-1, 1}
	targetIndex := 0
	targetD := 0
	nvs := []int{-1}
	for len(layer) > 0 {
		thisLine := nextLine
		nextLine = make([]int, int(math.Pow(2, float64(d)))+1)
		nextLayer := []*TreeNode{}

		// prepare nextLine and update thisLine
		li := 0
		ni := 1
		for i := 1; i < len(thisLine); i++ {
			if thisLine[i] < 0 {
				nextLine[ni] = -1
				ni++
				nextLine[ni] = -1
				ni++
			} else {
				thisLine[i] = layer[li].Val
				if layer[li] == target {
					targetIndex = int(math.Pow(2, float64(d-1))) - 1 + i
					targetD = d
				}
				if layer[li].Left != nil {
					nextLayer = append(nextLayer, layer[li].Left)
					nextLine[ni] = 1
				} else {
					nextLine[ni] = -1
				}
				ni++
				if layer[li].Right != nil {
					nextLayer = append(nextLayer, layer[li].Right)
					nextLine[ni] = 1
				} else {
					nextLine[ni] = -1
				}
				ni++
				li++
			}
		}

		layer = nextLayer
		nvs = append(nvs, thisLine[1:]...)
		d++
	}
	return &nvs, targetIndex, targetD
}

func findWithArray(nvs []int, k int, targetIndex int, targetD int) []int {
	res := []int{}
	resPtr := &res
	// down
	resPtr = findDown(nvs, k, targetIndex, targetD, -1, -1, resPtr)

	// up
	cd := targetD
	ci := targetIndex

	for cd > 1 && k > 0 {
		from := 0
		if ci/2*2 == ci {
			from = 0
		} else {
			from = 1
		}
		ci /= 2
		k--
		cd--

		if k == 0 {
			resPtr = findDown(nvs, k, ci, cd, -1, -1, resPtr)
		} else if k >= 1 {
			tis := ci
			tie := ci
			if from == 0 {
				tis *= 2
			} else {
				tis = tis*2 + 1
			}
			tie = tis
			for i := cd + 1; i < cd+k; i++ {
				tis *= 2
				tie = tie*2 + 1
			}

			if nvs[ci*2] > 0 && nvs[ci*2+1] > 0 {
				resPtr = findDown(nvs, k, ci, cd, tis, tie, resPtr)
			}
		}
	}
	return *resPtr
}

func findDown(nvs []int, k int, ci int, cd int, avoids int, avoide int, res *[]int) *[]int {
	// at cd + k depth, the left bound and right bound
	downLeft := ci
	downRight := ci
	for i := cd; i < cd+k; i++ {
		downRight = downRight*2 + 1
		downLeft *= 2
	}

	for j := downLeft; j <= downRight && j < len(nvs) && j >= 1; j++ {
		if nvs[j] >= 0 {
			if j >= avoids && j <= avoide {
				continue
			}
			*res = append((*res), nvs[j])
		}
	}
	return res
}
