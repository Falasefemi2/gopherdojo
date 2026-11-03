// 023: Struct Embedding. Embedding a type promotes its fields and methods
// to the outer struct, which reads like inheritance without being it.
//
// Scenario: an accounts service reuses one identity record across account
// and manager records, then labels them through the promoted method.
//
// Rules:
//  1. Account embeds Entity, so ID, Kind and Label are accessible on Account.
//  2. newAccount fills the embedded Entity and the Owner field.
//  3. accountLabel and accountID must use the promoted method and field.
//  4. transfer sets Owner through a pointer and returns the updated Account.
//  5. Manager embeds Account, keeping Entity fields and Label promoted two
//     levels deep.
package main

import "fmt"

// Entity is the common identity every record carries.
type Entity struct {
	ID   int
	Kind string
}

// Label formats the entity as kind#id.
func (e Entity) Label() string {
	return fmt.Sprintf("%s#%d", e.Kind, e.ID)
}

// Account embeds Entity, promoting ID, Kind and Label.
type Account struct {
	Entity
	Owner string
}

// Manager embeds Account, so Entity stays promoted through two levels.
type Manager struct {
	Account
	Level int
}

func newAccount(id int, kind, owner string) Account {
	return Account{Entity: Entity{ID: id, Kind: kind}, Owner: owner}
}

func accountLabel(a Account) string {
	return a.Label()
}

func accountID(a Account) int {
	return a.ID
}

func transfer(a *Account, owner string) Account {
	a.Owner = owner
	return *a
}

func main() {
	a := newAccount(7, "user", "ada")
	fmt.Println(accountLabel(a), accountID(a))
	a = transfer(&a, "grace")
	fmt.Println("new owner:", a.Owner)

	m := Manager{Account: a, Level: 2}
	fmt.Println("manager:", m.Label(), "level", m.Level)
}
