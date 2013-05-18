SHELL := /bin/bash

doc: clean
	phpdoc run --visibility=public,protected,private  --title pe2013 -f ./src/pe.php -t doc

clean:
	rm -rf doc/*

test_all: test1 test2 test3 test4 test5 test6 test7 

test1:
	diff -s <(php src/izbori.php ../../Tests/1/) ../../Tests/1/Result.txt

test2:
	diff -s <(php src/izbori.php ../../Tests/2/) ../../Tests/2/Results.txt

test3:
	diff -s <(php src/izbori.php ../../Tests/3/) ../../Tests/3/Result.txt

test4:
	diff -s <(php src/izbori.php ../../Tests/4/) ../../Tests/4/Result.txt

test5:
	diff -s <(php src/izbori.php ../../Tests/5/) ../../Tests/5/Result.txt

test6:
	diff -s <(php src/izbori.php ../../Tests/6/) ../../Tests/6/Result.txt

test7:
	diff -s <(php src/izbori.php ../../Tests/7/) ../../Tests/7/Result.txt
