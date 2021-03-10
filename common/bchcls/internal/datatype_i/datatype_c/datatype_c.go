//go:binary-only-package

package datatype_c

import (
	"common/bchcls/cached_stub"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/datatype/datatype_interface"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/common/graph"
	"common/bchcls/utils"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)