package models
type  	Order struct {

id 					string
order_id  			string
customer			string 			
cutomer_phone		string
date_deliv			string
recipeint			string
product				string
notes				string
card_text			string
deliv_price			string
zero_price			string
customer_price		string
paid				string
cash  - по факту	string
waiting 			string
payment				string


}

func NewOrder(id, order_id, customer, cutomer_phone, date_deliv, recipeint, product, notes, card_text, 
	deliv_price, zero_price, customer_price, paid, cash, waiting, payment string) *Order {
	return &Order{id, order_id, customer, cutomer_phone, date_deliv, recipeint, product, notes, card_text, 
	deliv_price, zero_price, customer_price, paid, cash, waiting, payment}
}

