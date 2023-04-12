package broadcaster

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	zeroAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	zeroHash = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
)

// Returns true if there is at least one filter value (or no filters at all) that matches an actual received value for every index i, or false otherwise
func filtersContainValues(topicValues []common.Hash, filters [][]common.Hash) bool {
	for i := 0; i < len(topicValues) && i < len(filters); i++ {
		filterValues := filters[i]

		// Empty filter for given index means: all values allowed
		valueFound := len(filterValues) == 0

		for _, filterValue := range filterValues {
			if filterValue == topicValues[i] {
				valueFound = true
				break
			}
		}
		if !valueFound {
			return false
		}
	}
	return true
}
