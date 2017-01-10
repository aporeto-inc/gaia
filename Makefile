init: install_monolithe install_monolithe_plugins
default: codegen

install_monolithe:
	pip install -U git+https://github.com/aporeto-inc/monolithe.git

install_monolithe_plugins:
	pip install 'git+https://github.com/aporeto-inc/elemental.git#subdirectory=monolithe'
	pip install 'git+https://github.com/aporeto-inc/pyelemental.git#subdirectory=monolithe'

codegen:
	monogen -f specs -L elemental
	# monogen -f specs -L pyelemental
	monogen -f specs -L html
	rm -f golang/*.go && cp codegen/elemental/1.0/*.go golang
	# rm -rf python/*.py python/requirements.txt MANIFEST.in && cp codegen/pyelemental/gaia/*.py python/gaia && cp codegen/pyelemental/requirements.txt python && cp codegen/pyelemental/MANIFEST.in python && cp codegen/pyelemental/setup.py python
	rm -rf apidoc/* && cp -a codegen/html/* apidoc
	rm -rf codegen
