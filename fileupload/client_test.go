package fileupload

import (
	"os"
	"testing"

	stripe "github.com/siriele/stripe-go"
	. "github.com/siriele/stripe-go/utils"
)

const (
	expectedSize = 734
	expectedMime = "application/pdf"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestFileUploadNewThenGet(t *testing.T) {
	f, err := os.Open("test_data.pdf")
	if err != nil {
		t.Errorf("Unable to open test file upload file %v\n", err)
	}

	uploadParams := &stripe.FileUploadParams{
		Purpose: "dispute_evidence",
		File:    f,
	}

	target, err := New(uploadParams)
	if err != nil {
		t.Error(err)
	}

	if target.Size != expectedSize {
		t.Errorf("(POST) Size %v does not match expected size %v\n", target.Size, expectedSize)
	}

	if target.Purpose != uploadParams.Purpose {
		t.Errorf("(POST) Purpose %v does not match expected purpose %v\n", target.Purpose, uploadParams.Purpose)
	}

	if target.Mime != expectedMime {
		t.Errorf("(POST) Mime %v does not match expected mime %v\n", target.Mime, expectedMime)
	}

	res, err := Get(target.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if target.ID != res.ID {
		t.Errorf("(GET) File upload id %q does not match expected id %q\n", target.ID, res.ID)
	}
}
