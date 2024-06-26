package greedy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	inputsGasAvailableAtIndices        = [][]int{{1, 2}, {1, 2, 3, 4, 5}, {204, 918, 18, 17, 35, 739, 913, 14, 76, 555, 333, 535, 653, 667, 52, 987, 422, 553, 599, 765, 494, 298, 16, 285, 272, 485, 989, 627, 422, 399, 258, 959, 475, 983, 535, 699, 663, 152, 606, 406, 173, 671, 559, 594, 531, 824, 898, 884, 491, 193, 315, 652, 799, 979, 890, 916, 331, 77, 650, 996, 367, 86, 767, 542, 858, 796, 264, 64, 513, 955, 669, 694, 382, 711, 710, 962, 854, 784, 299, 606, 655, 517, 376, 764, 998, 244, 896, 725, 218, 663, 965, 660, 803, 881, 482, 505, 336, 279}}
	inputsGasNeededFromThisIndexToNext = [][]int{{2, 1}, {3, 4, 5, 1, 2}, {273, 790, 131, 367, 914, 140, 727, 41, 628, 594, 725, 289, 205, 496, 290, 743, 363, 412, 644, 232, 173, 8, 787, 673, 798, 938, 510, 832, 495, 866, 628, 184, 654, 296, 734, 587, 142, 350, 870, 583, 825, 511, 184, 770, 173, 486, 41, 681, 82, 532, 570, 71, 934, 56, 524, 432, 307, 796, 622, 640, 705, 498, 109, 519, 616, 875, 895, 244, 688, 283, 49, 946, 313, 717, 819, 427, 845, 514, 809, 422, 233, 753, 176, 35, 76, 968, 836, 876, 551, 398, 12, 151, 910, 606, 932, 580, 795, 187}}
	outputsMinIndexToCompleteCircuit   = []int{1, 3, 31}
)

func TestCanCompleteCircuit(t *testing.T) {

	assert := require.New(t)

	for i := 0; i < len(inputsGasAvailableAtIndices); i++ {

		assert.Equal(outputsMinIndexToCompleteCircuit[i], canCompleteCircuit(inputsGasAvailableAtIndices[i], inputsGasNeededFromThisIndexToNext[i]), fmt.Sprintf("unexpected output for inputs %v, %v", inputsGasAvailableAtIndices[i], inputsGasNeededFromThisIndexToNext[i]))
	}

}
