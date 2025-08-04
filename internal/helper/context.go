package helper

import (
	"errors"
	"github.com/cynx-io/janus-gateway/internal/constant"
	"net/http"
)

func GetSiteKey(r *http.Request) (constant.SiteKey, error) {
	siteKey := r.Context().Value(constant.ContextKeySiteKey)
	if siteKey == nil {
		return "", errors.New("site key not found")
	}

	if sk, ok := siteKey.(constant.SiteKey); ok {
		return sk, nil
	} else {
		return "", errors.New("invalid site key type")
	}
}
