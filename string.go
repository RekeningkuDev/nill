package nill

type String Type[string]

func NewString(value string) String {
	return String{Valid: true, V: value}
}
