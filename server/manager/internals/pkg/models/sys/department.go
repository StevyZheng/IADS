package sys

type Department struct {
	BaseModel
	Name string `json:"name" bson:"name,omitempty"`
}
