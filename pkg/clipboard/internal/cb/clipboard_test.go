// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed by
// a GNU GPL-3.0 license that can be found in the LICENSE file.

package cb_test

import (
	"sync"
	"testing"

	"golang.design/x/midgard/pkg/clipboard/internal/cb"
	"golang.design/x/midgard/pkg/types"
)

func TestLocalClipboardConcurrentRead(t *testing.T) {
	// This test check that concurrent read/write to the clipboard does
	// not cause crashes on some specific platform, such as macOS.
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		cb.Read(types.ClipboardDataTypePlainText)
	}()
	go func() {
		defer wg.Done()
		cb.Read(types.ClipboardDataTypeImagePNG)
	}()
	wg.Wait()
}
