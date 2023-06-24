"""
A setuptools based setup module.

See:
https://packaging.python.org/en/latest/distributing.html
https://github.com/pypa/sampleproject
"""
import os

from setuptools import setup, find_packages

# pylint: disable=redefined-builtin

setup(
    name="aas-core3.0-golang-dev-scripts",
    version="0.0.1",
    description="Provide development scripts for aas-core3.0-golang.",
    url="https://github.com/aas-core-works/aas-core3.0-golang",
    author="Marko Ristin",
    author_email="marko@ristin.ch",
    classifiers=[
        "Development Status :: 5 - Production/Stable",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3.8",
        "Programming Language :: Python :: 3.9",
        "Programming Language :: Python :: 3.10",
    ],
    license="License :: OSI Approved :: MIT License",
    keywords="asset administration shell code generation industry 4.0 industrie i4.0",
    install_requires=[
        "icontract>=2.6.1,<3",
        "aas-core-meta@git+https://github.com/aas-core-works/aas-core-meta@44756fb#egg=aas-core-meta",
        "aas-core-codegen@git+https://github.com/aas-core-works/aas-core-codegen@607f65c#egg=aas-core-codegen",
    ],
    # fmt: off
    extras_require={
        "dev": [
            "black==23.1.0",
            "mypy==1.3.0",
            "pylint==2.17.1",
        ]
    },
    # fmt: on
    packages=find_packages(exclude=["continuous_integration_of_dev_scripts"]),
)
