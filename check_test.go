package main

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		cfg     string
		wantErr bool
	}{
		{
			name:    "With all allowed",
			data:    data,
			cfg:     configAllAllowed,
			wantErr: false,
		},
		{
			name:    "With RND not allowed, but package RND allowed",
			data:    data,
			cfg:     configPackageRNDAllowed,
			wantErr: false,
		},
		{
			name:    "With RND not allowed",
			data:    data,
			cfg:     configRNDNotAllowed,
			wantErr: true,
		},
		{
			name:    "With RND blocked",
			data:    data,
			cfg:     configRNDBlocked,
			wantErr: true,
		},
		{
			name:    "With RND allowed, package RND blocked",
			data:    data,
			cfg:     configRNDIsOKPackageRNDBlocked,
			wantErr: true,
		},
		// Add more test cases as necessary
	}

	for _, tt := range tests { //nolint:varnamelen
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := parseConfig([]byte(tt.cfg))
			if err != nil {
				t.Errorf("parseConfig() error = %v", err)
			}
			_, err = check([]byte(tt.data), cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
