#

# Makefile for the go-haproxy client.
#
# SYNOPSIS:
#
#   make tests 		 - runs the module's test suite.
#   make clean-tests - removes all files generated while running the test suite.
#   make clean 		 - removes all files generated by make.

# NOTE:
#
#	The test suite of the go-haproxy client is located under the $(TEST_DIR)
#	directory on the contrary to the official Golang way of defining tests.
#	The reason behind this approach is that I simply don't like maintaining
#	my test files along with the main source code.
#
# 	However, this approach causes an obvious issue, i.e., you can not import
# 	and subsequently test the client's non-exported functions. A temporary
# 	workaround in order to be able testing all client's functions and also
# 	continue maintaing the test suite in a separate folder, is to symlink the
# 	complete test suite to its proper location prior to running it, and
# 	unlink it as soon as the test run completes, either successfully or not.

#

# Source code directory location.
CODE_DIR = haproxy

# Tests directory location.
TEST_DIR = tests

# Filename format of testing files.
TEST_FILE_PATTERN = "*_test.go"


.PHONY: tests clean clean-tests

# We want to run the `clean-tests` target even in case of failure, so we've
# created a wrapper target named `_tests`. We could run the current command
# from our shells, however this is more handy since you can never forget it!
tests:
	make _tests || make clean-tests

clean-tests:
	find $(CODE_DIR) -name $(TEST_FILE_PATTERN) -type l | xargs -I{} unlink {}

clean: clean-tests

_tests:
	find $(TEST_DIR) -name $(TEST_FILE_PATTERN) | xargs -I{} ln -r -s {} $(CODE_DIR)
	go test -v $(CODE_DIR)/*.go
	make clean-tests
