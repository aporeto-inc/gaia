init: codegen
default: codegen

codegen: codegen_v1

codegen_v1:
	@echo '* Generating servicelauncher models v1'
	rm -rf codegen
	monogen -f v1/specs -L elemental
	# monogen -f v1/specs -L pyelemental
	# rm -rf ./v1/python/* && cp -a codegen/pyelemental/* ./v1/python
	rm -rf ./v1/golang/*.go && cp codegen/elemental/*.go ./v1/golang
	rm -rf codegen

publish:
	git pull
	make codegen
	git commit -am "codegen"
	git push
