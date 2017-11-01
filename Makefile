include domingo.mk

init: domingo_init
test: lint domingo_unit_tests

# custom linting until we fix the generation.
lint:
	@echo "Running linters..."
	@gometalinter --vendor --disable-all \
		--enable=vet \
		--enable=vetshadow \
		--enable=ineffassign \
		--enable=goconst \
		--enable=errcheck \
		--enable=varcheck \
		--enable=structcheck \
		--enable=gosimple \
		--enable=misspell \
		--enable=deadcode \
		--enable=staticcheck \
		--deadline 5m \
		--tests ./...

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'
	pip install 'git+https://github.com/aporeto-inc/pyelemental.git#subdirectory=monolithe'

codegen: codegen_squall codegen_zack codegen_vince codegen_midgard codegen_squall codegen_rufus codegen_yuffie codegen_barret codegen_highwind

codegen_zack:
	cd zackmodels && make codegen

codegen_vince:
	cd vincemodels && make codegen

codegen_midgard:
	cd midgardmodels && make codegen

codegen_squall:
	cd squallmodels && make codegen

codegen_rufus:
	cd rufusmodels && make codegen

codegen_yuffie:
	cd yuffiemodels && make codegen

codegen_barret:
	cd barretmodels && make codegen

codegen_highwind:
	cd highwindmodels && make codegen

publish:
	git pull
	make codegen -j 8
	git commit -am "codegen"
	git push
