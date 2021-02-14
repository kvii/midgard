// Copyright 2020 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPL-3.0
// license that can be found in the LICENSE file.

package platform_test

import (
	"strings"
	"sync"
	"testing"

	"changkun.de/x/midgard/internal/clipboard/platform"
	"changkun.de/x/midgard/internal/types"
)

func TestLocalClipboardConcurrentRead(t *testing.T) {
	// This test check that concurrent read/write to the clipboard does
	// not cause crashes on some specific platform, such as macOS.
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		platform.Read(types.MIMEPlainText)
	}()
	go func() {
		defer wg.Done()
		platform.Read(types.MIMEImagePNG)
	}()
	wg.Wait()
}

func TestLocalClipboardWrite(t *testing.T) {
	s := "hi"
	platform.Write([]byte(s), types.MIMEPlainText)
	buf := platform.Read(types.MIMEImagePNG)
	if buf != nil {
		t.Errorf("write as text but can be captured as image: %s", string(buf))
	}
	if strings.Compare("", string(buf)) != 0 {
		t.Errorf("expect: %s, go: %v", "", string(buf))
	}

	s = "there"
	buf = nil
	platform.Write([]byte(s), types.MIMEImagePNG)
	buf = platform.Read(types.MIMEPlainText)
	if buf != nil {
		t.Errorf("write as image but can be captured as text: %s", string(buf))
	}
	if strings.Compare("", string(buf)) != 0 {
		t.Errorf("expect: %s, go: %v", "", string(buf))
	}
}
