SERVICE=is-it-free
RUNAS=$(USER)

build:
	go build -o dist/$(SERVICE) cmd/service/main.go
	cp scripts/$(SERVICE).sh dist/$(SERVICE)-final.sh
	echo IIF_TOKEN="$(token)" >> dist/$(SERVICE)-final.sh
	echo IIF_URL=$(influxdb) >> dist/$(SERVICE)-final.sh
	echo IIF_GYM_SOURCE_URL=$(gym) >> dist/$(SERVICE)-final.sh
	echo IIF_POOL_SOURCE_URL=$(pool) >> dist/$(SERVICE)-final.sh

install:
	sudo install -m755 dist/$(SERVICE) /usr/local/bin/$(SERVICE)
	sudo cp services/$(SERVICE).service /etc/systemd/system/$(SERVICE)@.service
	sudo cp services/influxdb.service /etc/systemd/system/influxdb.service
	sudo cp dist/$(SERVICE)-final.sh /etc/sysconfig/$(SERVICE)
	sudo systemctl daemon-reload
	sudo systemctl enable --now $(SERVICE)@$(RUNAS).service
	sudo systemctl enable --now influxdb

clean:
	rm -f dist/$(SERVICE) || true
	rm -f dist/$(SERVICE)-final.sh || true

uninstall:
	sudo rm /etc/systemd/system/$(SERVICE)@.service
	sudo rm /usr/local/bin/$(SERVICE)
