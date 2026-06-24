package server

import (
	"testing"

	minio "github.com/minio/minio-go/v7"
)

func TestParseBucketLookup(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    minio.BucketLookupType
		wantErr bool
	}{
		{"empty defaults to auto", "", minio.BucketLookupAuto, false},
		{"auto", "auto", minio.BucketLookupAuto, false},
		{"dns", "dns", minio.BucketLookupDNS, false},
		{"path", "path", minio.BucketLookupPath, false},
		{"case insensitive", "DNS", minio.BucketLookupDNS, false},
		{"invalid", "vhost", minio.BucketLookupAuto, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseBucketLookup(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseBucketLookup(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("parseBucketLookup(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
