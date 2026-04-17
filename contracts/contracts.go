package contracts

type Producer string

const (
	UserService    Producer = "user.service"
	ProductService Producer = "product.service"
	OrderService   Producer = "order.service"
)

type EventType string

const (
	UserCreated   EventType = "USER_CREATED"
	OrderCreated  EventType = "ORDER_CREATED"
	OrderPayedFor EventType = "ORDER_PAYED_FOR"
	FundsAdded    EventType = "FUNDS_ADDED"
	FundsDebitted EventType = "FUNDS_DEBITTED"
)

type Topic string

const (
	WalletEvents Topic = "wallet.events"
	OrderEvents  Topic = "order.events"
	UserEvents   Topic = "user.events"
)
