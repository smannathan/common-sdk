//go:binary-only-package

package graph

import (
	"common/bchcls/cached_stub"
	"common/bchcls/custom_errors"
	"common/bchcls/internal/common/global"
	"common/bchcls/simple_rule"

	"encoding/json"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)