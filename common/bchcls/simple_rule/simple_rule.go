//go:binary-only-package

package simple_rule

import (
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"

	"encoding/json"
	"math"
	"strconv"
	"strings"
)