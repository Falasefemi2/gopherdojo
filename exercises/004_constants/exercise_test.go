package main

import "testing"

func TestPermissionFlags(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{"Read", Read, 1},
		{"Write", Write, 2},
		{"Execute", Execute, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}

	if all := Read | Write | Execute; all != 7 {
		t.Errorf("Read|Write|Execute = %d, want 7", all)
	}
}

func TestConstants(t *testing.T) {
	if KB != 1024 {
		t.Errorf("KB = %d, want 1024", KB)
	}
	if MaxUploadBytes != 5*1024*1024 {
		t.Errorf("MaxUploadBytes = %d, want %d", MaxUploadBytes, 5*1024*1024)
	}
}

func TestHasPermission(t *testing.T) {
	tests := []struct {
		name string
		mask int
		perm int
		want bool
	}{
		{"all bits", 7, 4, true},
		{"subset", Read | Execute, Execute, true},
		{"one bit missing", Read | Write, Execute, false},
		{"empty mask", 0, Read, false},
		{"empty permission", 7, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPermission(tt.mask, tt.perm); got != tt.want {
				t.Errorf("hasPermission(%d, %d) = %t, want %t", tt.mask, tt.perm, got, tt.want)
			}
		})
	}
}

func TestPermissionString(t *testing.T) {
	tests := []struct {
		name string
		mask int
		want string
	}{
		{"none", 0, "---"},
		{"read", Read, "r--"},
		{"write", Write, "-w-"},
		{"execute", Execute, "--x"},
		{"read and execute", Read | Execute, "r-x"},
		{"all", 7, "rwx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permissionString(tt.mask); got != tt.want {
				t.Errorf("permissionString(%d) = %q, want %q", tt.mask, got, tt.want)
			}
		})
	}
}

func TestFitsUpload(t *testing.T) {
	tests := []struct {
		name string
		size int64
		want bool
	}{
		{"empty file", 0, true},
		{"one byte", 1, true},
		{"exactly at limit", MaxUploadBytes, true},
		{"one byte over", MaxUploadBytes + 1, false},
		{"negative size", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fitsUpload(tt.size); got != tt.want {
				t.Errorf("fitsUpload(%d) = %t, want %t", tt.size, got, tt.want)
			}
		})
	}
}
