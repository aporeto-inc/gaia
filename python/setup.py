# -*- coding: utf-8 -*-

from setuptools import setup
import os

packages = ['gaia']
resources = []
api_version_path = "./gaia"

# Needed because git+https://github.com/aporeto-inc/gaia.git#subdirectory=python
# Subdirectory is not handled by dependency_links
# Removed when we will have a public pypi
[os.system("pip install %s" % line) for line in open('requirements.txt')],

setup(
    name='gaia',
    version="1.0",
    url='',
    author='',
    author_email='',
    packages=packages,
    description='SDK for Aporeto REST API',
    include_package_data=True,
    data_files=resources
)