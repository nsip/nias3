# nias3
NIAS as peer-to-peer triple store

See https://github.com/nsip/nias3/wiki/Design-specification

## Functionality implemented to date

Functionality illustrated in `test.sh`

* Endpoint for SIF XML: `http://localhost:1492/sifxml`
  * `POST http://localhost:1492/sifxml/:objects`: post a single SIF object. Any supplied RefID is overwritten with a local GUID. Saves the SIF/XML object as triples to the Hexastore. Returns the SIF/XML object retrieved from the Hexastore, with the local RefID. Returns error if the SIF/XML object is not of type `:object`.
  * `GET http://localhost:1492/sifxml/:objects/:refid`: retrieve SIF/XML object with RefId `:refid` from Hexastore.
  * `GET http://localhost:1492/sifxml/:object`: retrieves stream of all SIF/XML object of type `:object` from Hexastore.

A file watcher has been implemented on folder `./in`. Any file that is created or updated in that directory (or its subdirectories), with suffix `.xml`, is parsed, and its children XML records are posted to the Hexastore.

Functionality to come:
* Update SIF/XML
* Patch SIF/XML
* Delete SIF/XML
* Update SIF/XML from file watcher
* Stream input/output into Hexastore
* Genericise predicates in SIF/XML triples (currently position-specific for arrays)
* Timestamp triples with Lamport clocks
* GraphQL interface over Hexastore
* Peer-to-peer sharing of triples between NIAS3 nodes
* Encryption
* Digital signatures
