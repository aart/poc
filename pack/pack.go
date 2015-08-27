package pack

type Empty struct {
}

type Person struct {
	Name      string
	Firstname string
}

type Goods struct {
	Id             string
	Description    string
	Bulk           bool
	TotalLoading   int
	TotalNetWeight int
	TotalVolume    int
	TotalPackage   int
	TotalPallets   int
}

type Endpoint struct {
	Id     string
	Detail string
}

type TransportOrder struct {
	BusinessId  string
	Carrier     string
	Express     bool
	ContractRef string
	Goods       Goods
	Origin      Endpoint
	Destination Endpoint
}
