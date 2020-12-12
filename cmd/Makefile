SERVICE=is-it-free
RUNAS=$(USER)

build:
	go build -o $(SERVICE) main.go record.go pool.go gym.go
	cp $(SERVICE).sh $(SERVICE)-final.sh
	echo IIF_TOKEN="$(token)" >> $(SERVICE)-final.sh
	echo IIF_URL=$(influxdb) >> $(SERVICE)-final.sh
	echo IIF_GYM_SOURCE_URL=$(gym) >> $(SERVICE)-final.sh
	echo IIF_POOL_SOURCE_URL=$(pool) >> $(SERVICE)-final.sh

install:
	sudo install -m755 $(SERVICE) /usr/local/bin/$(SERVICE)
	sudo cp $(SERVICE).service /etc/systemd/system/$(SERVICE)@.service
	sudo cp influxdb.service /etc/systemd/system/influxdb.service
	sudo cp $(SERVICE)-final.sh /etc/sysconfig/$(SERVICE)
	sudo systemctl daemon-reload
	sudo systemctl enable --now $(SERVICE)@$(RUNAS).service
	sudo systemctl enable --now influxdb

clean:
	rm -f $(SERVICE) || true
	rm -f $(SERVICE)-final.sh || true

uninstall:
	sudo rm /etc/systemd/system/$(SERVICE)@.service
	sudo rm /usr/local/bin/$(SERVICE)
