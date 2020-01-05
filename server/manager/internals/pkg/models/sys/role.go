package sys

type Role struct {
	BaseModel
	Name    string `json:"name" bson:"name,omitempty"`
	Explain string `json:"explain" bson:"explain,omitempty"`
}
