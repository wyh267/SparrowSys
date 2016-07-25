package SparrowCache

import (
	"testing"
	"utils"
)

func TestNetwork(t *testing.T) {

	logger, _ := utils.New("test_network")
	nn := NewNetwork(5, 7777, logger)

	nn.Start()

}
