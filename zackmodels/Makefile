init: codegen
default: codegen

codegen: codegen_v1 codegen_v2

codegen_v1:
	@echo '* Generating Zack models v1'
	rm -rf codegen
	monogen -f v1/specs -L elemental
	rm -f ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	rm -rf codegen

codegen_v2:
	@echo '* Generating Zack models v2'
	rm -rf codegen
	monogen -f v2/specs -L elemental
	rm -f ./v2/golang/*.go && cp codegen/elemental/*.go ./v2/golang
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
