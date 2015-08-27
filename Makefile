PID = /tmp/goPusher.pid

vendor:
	@gb vendor update -all

build:
	@echo "Building..."
	@gb build

serve:
	@make restart
	@fswatch -o -r ./src/ | xargs -n1 -I{}  make restart || make kill

kill:
	@echo ""
	@echo ""
	@echo ""
	@echo "Trying to kill old instance..."
	@kill `cat $(PID)` || true

restart:
	@make kill
	@make build
	@bin/goPusher --env dev & echo $$! > $(PID)

.PHONY: vendor build serve kill restart
