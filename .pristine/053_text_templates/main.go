// 053: Text Templates. text/template fills placeholders from data and can
// branch with actions; Execute writes the rendered result to an io.Writer.
//
// Scenario: an access email greets the account holder and states the level of
// access granted by the account record.
//
// Rules:
//  1. Parse the template in accountTemplate, named "account".
//  2. Execute it into a bytes.Buffer and return the buffer's exact contents.
//  3. Admin accounts end with "Admin access enabled"; others end with
//     "Read-only access".
//  4. The greeting is "Hello <name>!" followed by a newline; there is no
//     trailing newline after the access line.
//
// TODO: implement renderAccount. Keep main() as the demo harness.
package main

import "fmt"

// Account is the data passed to the access email template.
type Account struct {
	Name  string
	Admin bool
}

const accountTemplate = `Hello {{.Name}}!
{{if .Admin}}Admin access enabled{{else}}Read-only access{{end}}`

// renderAccount renders the access email for one account.
func renderAccount(a Account) (string, error) {
	// TODO
	return "", nil
}

func main() {
	email, err := renderAccount(Account{Name: "Ada", Admin: true})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(email)
}
