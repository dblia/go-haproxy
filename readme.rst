.. _readme:

===============================
 HAProxy library written in Go
===============================


  **NOTE:** You should read the Caveats_ section first, before proceeding with
  the rest document

.. contents:: **Contents Table**
    :depth: 3

This is a simple library, written in Go, for interacting with the HAProxy UNIX
Socket commands interface.

Overview
========

The common way of interacting with the `HAProxy stats socket interface` is via a
common UNIX socket. However, even thought it is never done by default, it is
also possible to define multiple stats socket instance in HAProxy configuration,
and make them listen to a TCP port instead of a UNIX socket. The *go-haproxy*
library allows using *all* the supported ways for connecting to HAProxy, i.e., a
TCP port (either via IPv4 or IPv6), or via a UNIX socket.

The current library uses the *non-interactive* mode to connect to HAProxy stats
socket interface. This means you only send a single line per time. HAProxy will
process this line, send its response back to the client, and finally it will
close the connection immediately after the end of the response. In case you want
to send multiple commands in a single line, you need to delimit them with a
semi-colon (*';'*).

.. _HAProxy stats socket interface:
    https://cbonte.github.io/haproxy-dconv/1.7/management.html#9.3

Getting Started
===============

To facilitate a user, we ship a custom *Makefile* for managing the *go-haproxy*
library, i.e., ``build`` the library, run the ``tests``, etc.

Requirements
------------

There are no special Requirements for the current library, just a working `Go
environment`. The *go-haproxy* library is currently developed with the following
Stable release:

::

  $ go version
  go version go1.10.3 linux/amd64

.. _Go environment:
    https://golang.org/doc/install

Usage
-----

- Initialize a HAProxy connection object using a common *UNIX* socket:

.. code:: go

	haproxy := haproxy.HAProxyClient{
		AfNet:   "unix",
		Address: "/run/haproxy.sock",
		Timeout: time.Second * 2,
	}

- Initialize a HAProxy connection object using a *TCP* port:

.. code:: go

	// via IPv4
	haproxy := haproxy.HAProxyClient{
		AfNet:   "tcp4",
		Address: "10.0.0.1:8000",
		Timeout: time.Second * 2,
	}

	// via IPv6
	haproxy := haproxy.HAProxyClient{
		AfNet:   "tcp6",
		Address: "[fc00::]:8000",
		Timeout: time.Second * 2,
	}

- Run a HAProxy command and retrieve the response:

.. code:: go

	response, err := haproxy.SendCommand("show acl")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

For more information on using the current library, you can address the example
files under the ``samples/`` directory.

Development
-----------

As already said, you simply need a working Go environment and you'll have a
fully functioning development environment for the *go-haproxy* library.

- To properly format the source code, build the project and its dependencies, or
  run the tests, you are advised to use the *Makefile*:

::

    $ make build
    $ make fmt
    $ make tests
    $ make          // run everything

---

  **NOTE:** The test suite of the go-haproxy client is located under the
  $(TEST_DIR) directory on the contrary to the official Go lang way of defining
  tests.  The reason behind this approach is that I simply don't like
  maintaining my test files along with the main source code.

  However, this approach causes an obvious issue, i.e., you can not import and
  subsequently test the client's non-exported functions. A temporary workaround
  in order to be able testing all client's functions and also continue maintaining
  the test suite in a separate folder, is to symlink the complete test suite to
  its proper location prior to running it, and unlink it as soon as the test run
  completes, either successfully or not.

Caveats
=======

- This is still a work-in-progress library. More functionality in terms of
  improved error handling, handy HAProxy command wrappers, better unit tests,
  docs additions, etc, will be added in the future.
- Note that this library is more a mean to experiment with the Go language, its
  packaging system, the testing framework, and so on, rather than a tool tested
  in production environments.

.. vim: set textwidth=79 :
.. Local Variables:
.. mode: rst
.. fill-column: 79
.. End:
