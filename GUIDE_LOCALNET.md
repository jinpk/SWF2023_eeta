
```bash
eetad init node -o --chain-id eeta

eetad keys add u1

# minter addres 설정
# ~/.eeta/config/genesis.json
    "deposit": {
      "params": {
        "minter_address": <u1's address>
      }
    }
```

```bash
eetad add-genesis-account {address} 1000000000000stake 

eetad gentx u1 500000000000stake \
--chain-id eeta \
--moniker u1-vali \
--commission-rate="0.1" \
--commission-max-rate="0.20" \
--commission-max-change-rate="0.01" \
--min-self-delegation="1" \
--pubkey=$(eetad tendermint show-validator)

eetad collect-gentxs

eetad start

#  eetad tx deposit mint [receipient-address] [amount] [flags]
./eetad tx deposit mint {address} 10000krw
```