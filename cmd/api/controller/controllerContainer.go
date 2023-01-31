package controller

import (
	"go.uber.org/zap"
)

type Container struct {
}

func NewControllerContainer(
	logger *zap.Logger,
) *Container {
	return &Container{}
}
