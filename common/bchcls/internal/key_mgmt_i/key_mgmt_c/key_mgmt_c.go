//go:binary-only-package

package key_mgmt_c

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/custom_errors"
	"github.com/smannathan/common-sdk/common/bchcls/data_model"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/graph"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i/key_mgmt_c/key_mgmt_g"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)