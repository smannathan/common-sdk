//go:binary-only-package

package simple_rule

import (
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"

	"encoding/json"
	"math"
	"strconv"
	"strings"
)