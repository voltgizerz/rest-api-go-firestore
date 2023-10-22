package entity

type User struct {
	DocRefID  string `json:"doc_ref_id"`
	FirstName string `json:"first_name" firestore:"firstname"`
	LastName  string `json:"last_name" firestore:"lastname"`
	Username  string `json:"username" firestore:"username"`
	Email     string `json:"email" firestore:"email"`
	CCNumber  string `json:"cc_num" firestore:"cc_num"`
	CCType    string `json:"cc_type" firestore:"cc_type"`
	Country   string `json:"country" firestore:"country"`
	City      string `json:"city" firestore:"city"`
	Currency  string `json:"currency" firestore:"currency"`
}
