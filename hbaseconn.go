package main

import (
        "flag"
        "log"

        "hbase"
        "hbase/pb"
)

func main() {
        flag.Parse()
        hclient := hbase.NewClient("hbasetable")
        rows := make([]*pb.CellSet_Row, 0)
        key := []byte("<hashkey>")
        values := make([]*pb.Cell, 0)
        values = append(values, &pb.Cell{
                Row:    key,
                Column: []byte("cf:column"),
                Data:   []byte("asdf"),
        })
        row := &pb.CellSet_Row{
                Key:    key,
                Values: values,
        }
        rows = append(rows, row)
        if err := hclient.Put(&pb.CellSet{Rows: rows}); err != nil {
                log.Printf("Failed to put: %+v", err)
        }
}

