# plz

> NOTE: This package is a sub-deprendency of thrift-iterator/go. This batchcorp fork adds gls/goid_arm64.s to support compiling plumber for M1 macs.


PLZ has three parts

* plz: observability, test helper, message format, utility
* plz.sql: use sql to query anything
* plz.service: junction for distributed computing

Observability is the primary goal:

* countlog: log event, record how state change over the time
* dump: take snapshot of data, record the moment
* witch: a WEB UI to make log and snapshot visible