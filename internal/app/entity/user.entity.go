package entity

type User struct {
	FirstName string `firestore:"firstname"`
	LastName  string `firestore:"lastname"`
	Username  string `firestore:"username"`
	Email     string `firestore:"email"`
	CCNumber  string `firestore:"cc_num"`
	CCType    string `firestore:"cc_type"`
	Country   string `firestore:"country"`
	City      string `firestore:"city"`
	Currency  string `firestore:"currency"`
}
