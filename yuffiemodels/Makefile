init: codegen
default: codegen

codegen: codegen_v1

codegen_v1:
	@echo '* Generating Yuffie models v1'
	rm -rf codegen
	monogen -f v1/specs -L elemental
	rm -rf ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
