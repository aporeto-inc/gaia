init: codegen
default: codegen

codegen: codegen_v1

codegen_v1:
	@echo '* Generating Midgard models'
	rm -rf codegen
	monogen -f v1/specs -L elemental
	monogen -f v1/specs -L monodoc
	rm -f ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	rm -rf ./v1/apidoc/* && cp -a codegen/monodoc/* ./v1/apidoc
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
