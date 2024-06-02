package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	substrate "github.com/threefoldtech/tfchain/clients/tfchain-client-go"

	"github.com/threefoldtech/tfgrid-sdk-go/rmb-sdk-go/peer"
	// "rmbClient/peer"
)

const (
	chainUrl     = "wss://tfchain.grid.tf/"
	relayUrl     = "wss://relay.grid.tf"
	gridproxyUrl = "https://gridproxy.grid.tf"
)

func main() {

	var chainURLArg string
	var relayURLArg string
	var gridProxyArg string
	var mnemonicArg string
	var cmdArg string
	var nodeTwinIDArg uint64
	var nodeIDArg uint64

	flag.StringVar(&chainURLArg, "chainUrl", chainUrl, "chain url")
	flag.StringVar(&relayURLArg, "relayUrl", relayUrl, "relay url")
	flag.StringVar(&gridProxyArg, "gridProxyUrl", gridproxyUrl, "gridproxy url")
	flag.StringVar(&mnemonicArg, "mnemonic", "", "mnemonic")
	flag.StringVar(&cmdArg, "cmd", "", "rmb cmd")
	flag.Uint64Var(&nodeTwinIDArg, "twinID", 0, "node twin id")
	flag.Uint64Var(&nodeIDArg, "nodeID", 0, "node id")

	flag.Parse()
	if mnemonicArg == "" {
		mnemonicArg = os.Getenv("MNEMONIC")
	}
	if mnemonicArg == "" {
		log.Fatal().Msg("need to pass mnemonic either with -mnemonic or MNEMONIC env var")
	}
	if cmdArg == "" {
		log.Fatal().Msg("need to cmd with -cmd")
	}

	if nodeTwinIDArg == 0 && nodeIDArg == 0 {
		log.Fatal().Msg("need to pass twinID or nodeID")
	}

	if nodeIDArg > 0 && nodeTwinIDArg > 0 {
		log.Fatal().Msg("need to pass only twinID or nodeID")
	}
	if nodeIDArg > 0 {

		twinID, err := nodeIDToTwinID(gridProxyArg, nodeIDArg)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get twinID")
		}
		nodeTwinIDArg = uint64(twinID)
	}

	subMan := substrate.NewManager(chainUrl)

	// session := "elsession-xmon"
	client, err := peer.NewRpcClient(
		context.Background(),
		mnemonicArg,
		subMan,
		peer.WithKeyType(peer.KeyTypeSr25519),
		peer.WithRelay(relayURLArg),
		peer.WithSession("test-client"),
	)
	if err != nil {
		fmt.Println("failed to create peer client: %w", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result map[string]interface{}
	err = client.CallWithSession(ctx, uint32(nodeTwinIDArg), nil, cmdArg, nil, &result)
	if err != nil {
		fmt.Println("failed to call peer client: %w", err)
		os.Exit(1)
	}
	if data, err := json.MarshalIndent(result, "", "  "); err == nil {
		fmt.Println("========")
		fmt.Println(string(data))
		fmt.Println("========")
	}

}

func nodeIDToTwinID(gridProxyUrl string, nodeID uint64) (uint32, error) {
	endpoint := fmt.Sprintf("%s/nodes/%d", gridProxyUrl, nodeID)
	fmt.Println(endpoint)
	res, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("received non-200 response: %d", res.StatusCode)
	}
	type Node struct {
		TwinId uint32 `json:"twinId"`
	}
	var n Node
	err = json.NewDecoder(res.Body).Decode(&n)

	if err != nil {
		return 0, fmt.Errorf("failed to get nodes: %w", err)
	}

	return n.TwinId, err
}
