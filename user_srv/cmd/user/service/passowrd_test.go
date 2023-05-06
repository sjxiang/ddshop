package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPassowrd(t *testing.T) {
	assert.True(t, verifyPassowrd("$2a$10$u0ZZLugm8gxlY8GLZodzqeXn51IW8sRkHh2Zcxk.252St9FZMx8VC", "123456789"))
}

func TestHashPassowrd(t *testing.T) {
	res := hashPassowrd("123456789")
	t.Log(res)
}