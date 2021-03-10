//go:binary-only-package

package key_mgmt_c

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/data_model"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/common/graph"
	"common/bchcls/internal/key_mgmt_i/key_mgmt_c/key_mgmt_g"
	"common/bchcls/utils"

	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)