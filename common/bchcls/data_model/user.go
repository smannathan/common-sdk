//go:binary-only-package

package data_model

import (
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/key_mgmt_i/key_mgmt_c/key_mgmt_g"
	"github.com/smannathan/common-sdk/common/bchcls/internal/user_mgmt_i/user_mgmt_c/user_mgmt_g"

	"crypto/rsa"
	"encoding/json"
	"reflect"
)