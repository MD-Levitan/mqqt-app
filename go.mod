module github.com/MD-Levitan/mqqt-app

go 1.13

//replace github.com/MD-Levitan/bboltstore => ../bboltstore

require (
	github.com/MD-Levitan/bboltstore v0.0.0-20200422144831-0823678cda2f
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/sessions v1.2.0
	github.com/sirupsen/logrus v1.5.0
	go.etcd.io/bbolt v1.3.4
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	gopkg.in/yaml.v2 v2.2.8
)
