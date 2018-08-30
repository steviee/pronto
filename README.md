# Pronto

Distributed Storage Service for Protobuf-Defined Data, written in Golang called Pronto (because it's meant to be fast...)

## Why?

Storing data defined by .proto files someplace save and without wasting lots of disk-space is currently no easy feat.
Using something like Apacha Kafka to stream data to HBase or HDFS seems cumbersome at best and needs a lot of setup. 
And by the way I don't like (read: despise) the Hadoop infrastructure and Kafka for they are glorified jack-of-all-trades
services (eierlegende Wollmichsäue) and therefore involve too much hassle to be configured for simple use cases.

## How?

Pronto is meant to be a simple storage service where you create buckets of types (in form of .proto files) and are then able
to POST protocol buffers into those buckets. Each buffer will be assosiated with a cluster-unique Id-hash that you can use for 
fast access.

The buffers will be stored within the filesystem of the node running the ProntonDB service so chosing a filesystem for the 
desired type and expected buffer sizes is important (while Ext4 is almost always a good choice).

## Motivation

I have to store lots of records (currently mainly JSON documents) in a storage that's easily accessible and Elasticsearch is not
exactly disk-space friendly, a relational database is pure misuse (and not what I really need) and - as mentioned above - I don't want to use Apache Kafka or anything related.

I also think it's a good learning-experience regarding cluster software and high-load systems.

## When?

It's done when it's done...

# Initial Concept 

## Storing (binary) data

Data is stored in its original form (protobuf byte-streams) and directly written to disk. Incoming data gets a key determined by the node that accepted the request, the bucket the request belongs to and the time the request arrived. (for this I would use something like hashid).

This returns a relatively short key (12 chars) that is reversible and enables a quick lookup of a stored record.  
Also this solution scales quickly as a growing cluster would just broaden the number of possible keys but old keys stay valid.

It's unclear at this stage if each record will be saved as a single file or within a compound to save on diskspace (fragmentation with lots of small files).  
This depends vastly on the actual size of data records and the used indexing library (if any).

## Clustering

Each server instance should be perfectly fine alone. I suppose making use of hashicorps memberlist package makes sense for joining multiple instances into a cluster.

The cluster setup addresses three issues:

* Disk Space: Hard disk space may be added by adding instances to the cluster (new instance with new filesystem joins)
* System Load: Load may be split bei adding instances to the cluster (new instance with same filesystem joins)
* Requests may be forwarded to other nodes if
  * out of diskspace
  * disk utilization is not in parity with other nodes
  * system load is too high
  * as a general load balancing method (round-robin)

The nodes of the cluster would ping all other nodes to ascertain latency and get current disk utilisation stats.

## Accessing the service

The service is accessible via gRPC calls to create, read, update and delete a record within a bucket.

There could also be a JSON REST endpoint for CRUD operations against a bucket that would use on-demand converting (marshalling and unmarshalling) of JSON to ProtoBuf and vice-versa.

## Retention

Records get certain metadata making sure there's a TTL configurable. 

## GDPR

Records have .proto files so are generally parsable. Addig a query engine to find records with a certain property is not part of the initial implementation.

This might be added later wich added indexing libraries (bolt? lucene?bleve?)