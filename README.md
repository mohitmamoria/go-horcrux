### Horcrux

```
go run main.go
```

Sample output:

```
-> given...
private key: 6a7f8201ccf6366175d67235b9915cf6676130aa6e2353181b587f51fd33ab43
passphrase: avada_kedavra

-> we will split it into 2-of-4 scheme
shard #1: 90ae6adfd80c6ad572ccc060de9b909b90034b820e78760a0186ce5e5fcc2869d1c53e43e48736fce230104a682d9a4f8a85af7dc388a48eba48c1d424359afe4d
shard #2: af0aa11f3fe900ded40bff252d2aaf8b942508564a8fe3d42fdc155f890b33d497c9fdcb1e1e48f15538474ca3e9b1103cc303e8553686d4ce46cd92b14b662c1b
shard #3: 30b4af824bed684d1a3e607f14e1302674f4c286285689af2558b667293e0d2485f2280fc44f6df20ee3c0dcad94634c5fd16682778eb750400df680db6e2c6ca4
shard #4: 5cfac793dc172f3439230c111c8b5c63aec91b08fd84b3e3b7d09bdcca23313b3d87efbed800b53cd34c866bc55f90a5396987b89b4750d85a158338e1b6875559

-> storing sever_shard (90ae6adfd80c6ad572ccc060de9b909b90034b820e78760a0186ce5e5fcc2869d1c53e43e48736fce230104a682d9a4f8a85af7dc388a48eba48c1d424359afe4d) on the server

-> storing encrypted client_shard (dc92f5d46d598e7d47ce5cde95ddabda7524807ab0ee76550543a561176a3094241475b22526d2c7bbbdf29031c208797e3f9074fb3539f2029e70969f1f8df8e2bdcb041cbc52f7b63e8660b2d0fc2d8620b020c62181db124e7056c2) on the server

-> showing backup_shards to the user
30b4af824bed684d1a3e607f14e1302674f4c286285689af2558b667293e0d2485f2280fc44f6df20ee3c0dcad94634c5fd16682778eb750400df680db6e2c6ca4
5cfac793dc172f3439230c111c8b5c63aec91b08fd84b3e3b7d09bdcca23313b3d87efbed800b53cd34c866bc55f90a5396987b89b4750d85a158338e1b6875559

-> when logging in, server will respond with encrypted client shard (dc92f5d46d598e7d47ce5cde95ddabda7524807ab0ee76550543a561176a3094241475b22526d2c7bbbdf29031c208797e3f9074fb3539f2029e70969f1f8df8e2bdcb041cbc52f7b63e8660b2d0fc2d8620b020c62181db124e7056c2)
then, client will decrypt it using passphrase (avada_kedavra)
into the client_shard (af0aa11f3fe900ded40bff252d2aaf8b942508564a8fe3d42fdc155f890b33d497c9fdcb1e1e48f15538474ca3e9b1103cc303e8553686d4ce46cd92b14b662c1b)

-> when signing a transaction, we will use the server_shard (90ae6adfd80c6ad572ccc060de9b909b90034b820e78760a0186ce5e5fcc2869d1c53e43e48736fce230104a682d9a4f8a85af7dc388a48eba48c1d424359afe4d) and client_shard (af0aa11f3fe900ded40bff252d2aaf8b942508564a8fe3d42fdc155f890b33d497c9fdcb1e1e48f15538474ca3e9b1103cc303e8553686d4ce46cd92b14b662c1b)
to get recreate the private key to sign the transaction:
6a7f8201ccf6366175d67235b9915cf6676130aa6e2353181b587f51fd33ab43
****************************👆🏻*********************************
**************************TADAA!******************************
**************************************************************
```