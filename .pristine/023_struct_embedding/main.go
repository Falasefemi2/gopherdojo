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
//
// TODO: implement newAccount, accountLabel, accountID and transfer.
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

// newAccount builds an Account with the embedded Entity filled in.
func newAccount(id int, kind, owner string) Account {
	return Account{}
}

// accountLabel returns the promoted entity label.
func accountLabel(a Account) string {
	return ""
}

// accountID returns the promoted entity ID.
func accountID(a Account) int {
	return 0
}

// transfer sets the owner through a pointer and returns the updated value.
func transfer(a *Account, owner string) Account {
	return Account{}
}

func main() {
	a := newAccount(7, "user", "ada")
	fmt.Println(accountLabel(a), accountID(a))
	a = transfer(&a, "grace")
	fmt.Println("new owner:", a.Owner)

	m := Manager{Account: a, Level: 2}
	fmt.Println("manager:", m.Label(), "level", m.Level)
}
