package models

type Product struct{
	ID    string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string   `json:"name" bson:"name"`
	Type  string   `json:"type" bson:"type"`
	Price string   `json:"price" bson:"price"`
}
