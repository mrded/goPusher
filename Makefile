PID = /tmp/goPusher.pid

vendor:
	@gb vendor update -all

build:
	@echo "Building..."
	@gb build all

serve:
	@make restart

kill:
	@echo ""
	@echo ""
	@echo ""
	@echo "Trying to kill old instance..."
	@kill `cat $(PID)` || true

restart:
	@make kill
	@make build
	@bin/goPusher & echo $$! > $(PID)

.PHONY: vendor build serve kill restart
