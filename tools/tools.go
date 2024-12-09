package tools

import (
	// This package is explicitly included to support the mockgen utility,
	// which was introduced after rebasing to v2.8.2.
	// The downstream fork builds the image using the vendor directory,
	// where module lookup is disabled:
	// "cannot find module providing package github.com/golang/mock/mockgen/model:
	// import lookup disabled by -mod=vendor".
	// To ensure the package is included in the vendor directory,
	// it must be an explicit dependency.
	_ "github.com/golang/mock/mockgen/model"
)
