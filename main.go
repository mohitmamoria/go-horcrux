package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"horcrux/shamir"
	"io"
)

type horcrux struct {
	serverShard  []byte
	clientShard  []byte
	backup1Shard []byte
	backup2Shard []byte
}

type server struct {
	serverShard    []byte
	encClientShard []byte
}

type session struct {
	clientShard []byte
}

func main() {
	privKey := "6a7f8201ccf6366175d67235b9915cf6676130aa6e2353181b587f51fd33ab43"
	passphrase := "avada_kedavra"
	fmt.Printf("-> given... \nprivate key: %s\npassphrase: %s", privKey, passphrase)

	fmt.Printf("\n\n-> we will split it into 2-of-4 scheme")
	shards, err := shamir.Split([]byte(privKey), 4, 2)
	if err != nil {
		panic(err)
	}
	for i, shard := range shards {
		fmt.Printf("\nshard #%d: %s", i+1, hex.EncodeToString(shard))
	}
	voldemort := horcrux{
		serverShard:  shards[0],
		clientShard:  shards[1],
		backup1Shard: shards[2],
		backup2Shard: shards[3],
	}
	encClientShard := encrypt(voldemort.clientShard, passphrase)

	voldemortServer := server{}
	fmt.Printf("\n\n-> storing sever_shard (%s) on the server", hex.EncodeToString(voldemort.serverShard))
	voldemortServer.serverShard = voldemort.serverShard
	fmt.Printf("\n\n-> storing encrypted client_shard (%s) on the server", hex.EncodeToString(encClientShard))
	voldemortServer.encClientShard = encClientShard

	fmt.Printf("\n\n-> showing backup_shards to the user")
	fmt.Printf("\n%s", hex.EncodeToString(voldemort.backup1Shard))
	fmt.Printf("\n%s", hex.EncodeToString(voldemort.backup2Shard))

	voldemort = horcrux{} // forgetting all previous values, to clean the slate

	voldemortSession := session{}
	fmt.Printf("\n\n-> when logging in, server will respond with encrypted client shard (%s)", hex.EncodeToString(voldemortServer.encClientShard))
	fmt.Printf("\nthen, client will decrypt it using passphrase (%s)", passphrase)
	decClientShard := decrypt(voldemortServer.encClientShard, passphrase)
	fmt.Printf("\ninto the client_shard (%s)", hex.EncodeToString(decClientShard))
	voldemortSession.clientShard = decClientShard

	fmt.Printf(
		"\n\n-> when signing a transaction, we will use the server_shard (%s) and client_shard (%s)",
		hex.EncodeToString(voldemortServer.serverShard),
		hex.EncodeToString(voldemortSession.clientShard),
	)
	recreatedPrivateKey, err := shamir.Combine([][]byte{voldemortServer.serverShard, voldemortSession.clientShard})
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nto get recreate the private key to sign the transaction:\n%s", string(recreatedPrivateKey))

	fmt.Printf("\n****************************👆🏻*********************************")
	fmt.Printf("\n**************************TADAA!******************************")
	fmt.Printf("\n**************************************************************")
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
