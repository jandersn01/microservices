module github.com/jandersn01/microservices/order

go 1.26.1

require(
	github.com/jandersn01/microservices-proto/golang/order v0.0.0-00010101000000-000000000000;
	github.com/jandersn01/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
)

replace github.com/jandersn01/microservices-proto/golang/order => ../../microservices-proto/golang/

replace github.com/jandersn01/microservices-proto/golang/payment => ../../microservices-proto/golang/payment
