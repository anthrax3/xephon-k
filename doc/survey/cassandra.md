# Cassandra

NOTE: moved to [separated folder](cassandra)

## Usage

For Fedora `sudo systemctl start docker.service` might be needed

- `docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.9`
  - `docker stop tsdb-cassandra`
  - `docker start tsdb-cassandra`
- `docker exec -it tsdb-cassandra bash`

## TODO

## DONE

- [x] keys http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra
- [x] Thrift, CQL and the underlying storage (I think I got a bit confused when trying to use
  CQL to understand KairosDB's schema design), see [Cassandra](cassandra/README.md)
