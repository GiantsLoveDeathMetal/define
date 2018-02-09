from setuptools import setup
from setuptools import find_packages


setup(
    name='cli-dictionary',
    version='0.0.0',
    author='GiantsLoveDeathMetal',
    author_email='s.williamswynn.mail@gmail.com',
    packages=find_packages(
        include='define',
        exclude='tests',
    ),
    entry_points={
        'console_scripts': ['define=define.cli:main'],
    },
    install_requires=[
        'requests',
    ],
    classifiers=[
        "Environment :: Console",
        "Operation System :: OS Independent",
        "Programming Language :: Python",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.6",
    ],
)
