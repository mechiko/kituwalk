package znakdb

type Order struct {
	Id                  int64  `db:"id"`
	CreateDate          string `db:"create_date"`
	Gtin                string `db:"gtin"`
	Quantity            int    `db:"quantity"`
	SerialNumberType    string `db:"serial_number_type"`
	TemplateId          int    `db:"template_id"`
	CisType             string `db:"cis_type"`
	ContactPerson       string `db:"contact_person"`
	ReleaseMethodType   string `db:"release_method_type"`
	CreateMethodType    string `db:"create_method_type"`
	PaymentType         string `db:"payment_type"`
	ProductionOrderId   string `db:"production_order_id"`
	ProductName         string `db:"product_name"`
	ProductCapacity     string `db:"product_capacity"`
	ProductShelfLife    string `db:"product_shelf_life"`
	ProductTemplate     string `db:"product_template"`
	Comment             string `db:"comment"`
	Version             string `db:"version"`
	State               string `db:"state"`
	Status              string `db:"status"`
	OrderId             string `db:"order_id"`
	Archive             int    `db:"archive"`
	Json                string `db:"json"`
	ServiceProviderId   string `db:"service_provider_id"`
	ServiceProviderName string `db:"service_provider_name"`
	ServiceProviderRole string `db:"service_provider_role"`
}

func (z *DbZnak) FindOrder(number int64) (order *Order, err error) {
	sess := z.dbSession
	order = &Order{}
	res := sess.Collection("order_mark_codes").Find("id", number)
	if err := res.One(order); err != nil {
		return order, err
	}
	return order, nil
}
