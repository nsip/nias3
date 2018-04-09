# nias3
NIAS as peer-to-peer triple store

See https://github.com/nsip/nias3/wiki/Design-specification

## NSIP Infrastructure As A Service

NIAS (NSIP Infrastructure As A Service) is a suite of tools intended to make the interoperability infrastructure promoted by NSIP more readily available as software components, targeted in particular to desktop deployments which do not have ready access to enterprise-scale infrastructure (data hubs).

NIAS uses the following core concepts:

* Data streaming (as a local bus for data), feeding Microservices
* Local validation and conversion of data
* Graph database representation of data
* Accommodation of multiple data standards
* [CQRS](https://martinfowler.com/bliki/CQRS.html), to enable Eventual Concurrency
* Peer-to-peer synchronisation of data between nodes, to obviate the need for server/client configuration

NIAS has gone through the following implementation stages:
* [NIAS](https://github.com/nsip/nias), implemented in Ruby with Apache Kafka data streaming, Redis as graph database, LMDS as data store. Initial proof of concept; stream chaining, filtering, format conversion, multiple standards in graph. Simple Analytics front-end.
* [NIAS2](https://github.com/nsip/nias2), implemented in Go with NATS Streaming, Ledis as graph database and data store. Optimised for performance; intended for desktop deployment. Specific to NAPLAN use cases.
* [NIAS3](https://github.com/nsip/nias3), implemented in Go with NATS Streaming, Ledis as Hexastore graph database. Move to triples as storage unit, peer-to-peer synchronisation of data.

## Functionality implemented to date

Functionality illustrated in `test.sh`

### SIF XML server

* Endpoint for SIF XML: `http://localhost:1492/sifxml`. This is not a full implementation of the SIF infrastructure at all; it deals only with CRUD object services at a rudimentary level, with single object payloads. 
  * `POST http://localhost:1492/sifxml/:objects`: post a single SIF object. Any supplied RefID is overwritten with a local GUID. Saves the SIF/XML object as triples to the Hexastore. Returns the SIF/XML object retrieved from the Hexastore, with the local RefID. Returns error if the SIF/XML object is not of type `:object`.
  * `GET http://localhost:1492/sifxml/:objects/:refid`: retrieve SIF/XML object with RefId `:refid` from Hexastore.
  * `GET http://localhost:1492/sifxml/:object`: retrieves stream of all SIF/XML object of type `:object` from Hexastore.
  * `PUT http://localhost:1492/sixml/:object/:refid`: with HTTP header `replacement` = `FULL`: does full object update: the triples for the object in the Hexastore are deleted, and replaced with the triples in the SIF/XML object.
  * `PUT http://localhost:1492/sixml/:object/:refid`: with HTTP header `replacement` = `PARTIAL`,  or with no HTTP header `replacement`, does partial object update. The triples in the payload update the corresponding entries in the Hexastore. Nothing is deleted in the Hexastore (unless there is a nil entry in the payload): lists in the original object do not have items shrunk. (So if a list in the original object has two items, and an update is sent with one item, the first list item is updated, and the second list item is left alone.)
  * `DELETE http://localhost:1492/sifxml/:objects/:refid`: deletes SIF/XML object with RefId `:refid` from Hexastore.

### SIF XML Client

* The function `webserver.SIFGetToDataStore(url string) error` performs a GET on a SIF object service endpoint, fetches the result as a multiple object payload, parses it into individual objects, and adds them to the Hexastore. An example call is made at the start of the main executable:

````
webserver.SIFGetToDataStore("http://hits.nsip.edu.au/SIF3InfraREST/hits/requests/SchoolInfos?navigationPage=1&navigationPageSize=5&access_token=ZmZhODMzNjEtMGExOC00NDk5LTgyNjMtYjMwNjI4MGRjZDRlOmYxYzA1NjNhOWIzZTQyMGJiMDdkYTJkOTBkYjQ3OWVm&authenticationMethod=Basic")
````

This function fetches the first 5 records from a SchoolInfos endpoint on the [HITS Server](http://hits.nsip.edu.au), and ingests them into the Hexastore. 

Performing the initial handshake to obtain authentication for the SIF endpoing (the access token in this instance) is outside the scope of this module.

### File watcher

A file watcher has been implemented on folder `./in`. Any file that is created or updated in that directory (or its subdirectories), with suffix `.xml`, is parsed, and its children XML records are posted to the Hexastore.

Functionality to come:
* Update SIF/XML from file watcher


## Functionality to come

* Stream input/output into Hexastore
* Genericise predicates in SIF/XML triples (currently position-specific for arrays)
* Timestamp triples with Lamport clocks
* GraphQL interface over Hexastore
* Peer-to-peer sharing of triples between NIAS3 nodes
* Encryption
* Digital signatures
