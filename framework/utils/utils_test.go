package utils_test

import (
	"testing"

	"github.com/Andreis3/encoder-video-golang/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJon(t *testing.T) {
	s := `{"name":"test"}`
	err := utils.IsJson(s)
	require.Nil(t, err)

	s = `name`
	err = utils.IsJson(s)
	require.NotNil(t, err)
}
