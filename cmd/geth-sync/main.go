package main

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
  client, err := ethclient.Dial("/var/lib/goethereum/geth.ipc")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    sync_progress, err := client.SyncProgress(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    syncJSON, err := json.MarshalIndent(sync_progress, "", "  ")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("The client sync info: %v\n", string(syncJSON))

    block, err := client.BlockByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The top block is: %v\n", block.Number().Text(10))
    fmt.Printf("The hash is: %s\n", block.Hash().Hex())

    //Unix Timestamp to time.Time
    block_time := time.Unix(int64(block.Time()), 0)
    fmt.Printf("The block time is: %s\n", block_time)
}
