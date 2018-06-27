spanner2csv
================
SELECT query results on Cloud Spanner to CSV

## Install

```
go get -u cloud.google.com/go/spanner github.com/castaneai/spanner2csv
```

## Usage

```
./spanner2csv projects/<projectId>/instances/<instaneId>/databases/<databaseId> "SELECT * FROM ..."
```
