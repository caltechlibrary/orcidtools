
The public API requires maintaining secrets.

Taking a look at the public data release. Tar ball is large with millions of JSON blobs. The individual blobs look fine.
Can I easily parse these into a DB (k/v, tripple, SQL???)

Tarball is actually a gzipped tar file. Gunzip then process the tar file.

How do I want to store the JSON blobs?  Bleve, Boltdb, SQL?

