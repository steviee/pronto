# ProtonDB
Distributed Storage Service for Protobuf-Defined Data, written in Golang

# Why?
Storing data defined by .proto files someplace save and without wasting lots of disk-space is currently no easy feat.
Using something like Apacha Kafka to stream data to HBase or HDFS seems cumbersome at best and needs a lot of setup. 
And by the way I don't like (read: despise) the Hadoop infrastructure and Kafka for they are glorified jack-of-all-trades
services (eierlegende Wollmichsäue) and therefore involve too much hassle to be configured for simple use cases.

# How?
ProtonDB is meant to be a simple storage service where you create buckets of types (in form of .proto files) and are then able
to POST protocol buffers into those buckets. Each buffer will be assosiated with a cluster-unique Id-hash that you can use for 
fast access.

The buffers will be stored within the filesystem of the node running the ProntonDB service so chosing a filesystem for the 
desired type and expected buffer sizes is important (while Ext4 is almost always a good choice).

# Motivation
I have to store lots of records (currently mainly JSON documents) in a storage that's easily accessible and Elasticsearch is not
exactly disk-space friendly, a relational database is pure misuse (and not what I really need) and - as mentioned above - I don't
want to use Apache Kafka or anything related.

# When?
It's done when it's done...
