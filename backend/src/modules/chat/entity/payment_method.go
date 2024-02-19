package entity

type PaymentMethod struct {
	Code         string
	Name         string
	Description  string
	IconLightUrl string
	IconDarkUrl  string
	Child        []*PaymentMethod
}
