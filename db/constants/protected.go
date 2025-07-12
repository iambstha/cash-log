package constants

var ProtectedTransactionTypes = []string{"income", "expense"}

var ProtectedCategories = map[string][]string{
	"income":  {"Salary", "Freelance"},
	"expense": {"Grocery", "Rent"},
}
