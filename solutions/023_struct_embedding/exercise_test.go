package main

import "testing"

func TestNewAccount(t *testing.T) {
	a := newAccount(7, "user", "ada")

	if a.ID != 7 {
		t.Errorf("newAccount(7, ...).ID = %d, want 7", a.ID)
	}
	if a.Kind != "user" {
		t.Errorf("newAccount(7, %q, ...).Kind = %q, want %q", "user", a.Kind, "user")
	}
	if a.Owner != "ada" {
		t.Errorf("newAccount(..., %q).Owner = %q, want %q", "ada", a.Owner, "ada")
	}
	if got, want := a.Label(), "user#7"; got != want {
		t.Errorf("newAccount(7, \"user\", \"ada\").Label() = %q, want %q", got, want)
	}

	zero := newAccount(0, "", "")
	if zero.ID != 0 || zero.Kind != "" || zero.Owner != "" {
		t.Errorf("newAccount(0, \"\", \"\") = %+v, want all zero fields", zero)
	}
}

func TestAccountPromotion(t *testing.T) {
	tests := []struct {
		name      string
		account   Account
		wantLabel string
		wantID    int
	}{
		{
			name:      "service",
			account:   Account{Entity: Entity{ID: 3, Kind: "svc"}, Owner: "ops"},
			wantLabel: "svc#3",
			wantID:    3,
		},
		{
			name:      "user",
			account:   Account{Entity: Entity{ID: 42, Kind: "user"}, Owner: "ada"},
			wantLabel: "user#42",
			wantID:    42,
		},
		{
			name:      "zero entity",
			account:   Account{Owner: "nobody"},
			wantLabel: "#0",
			wantID:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountLabel(tt.account); got != tt.wantLabel {
				t.Errorf("accountLabel(%+v) = %q, want %q", tt.account, got, tt.wantLabel)
			}
			if got := accountID(tt.account); got != tt.wantID {
				t.Errorf("accountID(%+v) = %d, want %d", tt.account, got, tt.wantID)
			}
			if got := tt.account.Label(); got != tt.wantLabel {
				t.Errorf("promoted Label() = %q, want %q", got, tt.wantLabel)
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	a := newAccount(1, "user", "ada")
	got := transfer(&a, "bob")

	if a.Owner != "bob" {
		t.Errorf("after transfer(&a, \"bob\"), a.Owner = %q, want %q", a.Owner, "bob")
	}
	if got.Owner != "bob" {
		t.Errorf("transfer returned Owner %q, want %q", got.Owner, "bob")
	}
	if got.ID != 1 || got.Kind != "user" {
		t.Errorf("transfer changed the embedded entity: %+v, want ID 1 and Kind \"user\"", got)
	}
}

func TestManagerPromotion(t *testing.T) {
	m := Manager{
		Account: Account{Entity: Entity{ID: 9, Kind: "admin"}, Owner: "grace"},
		Level:   3,
	}

	if m.ID != 9 {
		t.Errorf("promoted ID = %d, want 9", m.ID)
	}
	if m.Kind != "admin" {
		t.Errorf("promoted Kind = %q, want %q", m.Kind, "admin")
	}
	if m.Owner != "grace" {
		t.Errorf("promoted Owner = %q, want %q", m.Owner, "grace")
	}
	if got, want := m.Label(), "admin#9"; got != want {
		t.Errorf("promoted Label() = %q, want %q", got, want)
	}
	if got, want := m.Account.Label(), "admin#9"; got != want {
		t.Errorf("m.Account.Label() = %q, want %q", got, want)
	}
}
