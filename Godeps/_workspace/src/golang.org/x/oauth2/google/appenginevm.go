// Copyright 2014 The oauth2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appenginevm

package google

import (
	"time"

	"github.com/goffee/goffee/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/goffee/goffee/Godeps/_workspace/src/golang.org/x/oauth2"
	"google.golang.org/appengine"
)

// AppEngineTokenSource returns a token source that fetches tokens
// issued to the current App Engine application's service account.
// If you are implementing a 3-legged OAuth 2.0 flow on App Engine
// that involves user accounts, see oauth2.Config instead.
//
// The provided context must have come from appengine.NewContext.
func AppEngineTokenSource(ctx oauth2.Context, scope ...string) oauth2.TokenSource {
	return &appEngineTokenSource{
		ctx:         ctx,
		scopes:      scope,
		fetcherFunc: aeVMFetcherFunc,
	}
}

var aeVMFetcherFunc = func(ctx oauth2.Context, scope ...string) (string, time.Time, error) {
	c, ok := ctx.(context.Context)
	if !ok {
		return "", time.Time{}, errInvalidContext
	}
	return appengine.AccessToken(c, scope...)
}
