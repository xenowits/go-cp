package main

import (
	"fmt"
	"strings"
)

//Replace prose1 with any string to be encrypted

var prose1 string =
	`The load caused the primary replica of the password manager to become unrespon‐
sive, so the load balancer diverted traffic to the secondary replica, which promptly
failed in the same way. At this point, the system paged the on-call engineer. The engi‐
neer had no experience responding to failures of the service: the password manager
was supported on a best-effort basis, and had never suffered an outage in its five years
of existence. The engineer attempted to restart the service, but did not know that a
restart required a hardware security module (HSM) smart card.`

func convertProse(str string) (string) {

	str = strings.ToLower(str)
	ans := ""

	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			ans += string(ch)
		}
	}

	return ans
}

func encryptProse(prose, key string) string {
	//Follow vigenere Cipher procedure
	ans := ""
	key_length := len(key)
	prose_length := len(prose)

	i := 0
	for i < prose_length {
		cnt := 0
		ans_seg := ""
		for cnt < key_length && i < prose_length {
			ans_seg += string('a'+((prose[i]-'a') + (key[cnt]-'a'))%26)
			i += 1
			cnt += 1
		}
		fmt.Printf("%s\n----------------------------------\n", ans_seg)
		ans += ans_seg
	}

	return ans
}

func decryptProse(prose, key string) string {
	ans := ""

	key_length := len(key)
        prose_length := len(prose)

        i := 0
        for i < prose_length {
                cnt := 0
		ans_seg := ""
                for cnt < key_length && i < prose_length {
                        ans_seg += string('a'+((prose[i]-'a') - (key[cnt]-'a') + 26)%26)
                        i += 1
                        cnt += 1
                }
                // fmt.Println("%s\n----------------------------------\n", ans_seg)
                ans += ans_seg
        }

	return ans
}

func main() {
	cipher_key := "vigenerecipher"

	prose_to_be_encrypted := convertProse(prose1)

	encrypted_prose := encryptProse(prose_to_be_encrypted, cipher_key)

	fmt.Print("Prose to be Encrypted:\n", prose_to_be_encrypted, "\n")

	fmt.Print("Encrypted Prose:\n", encrypted_prose, "\n")

	decrypted_prose := decryptProse(encrypted_prose, cipher_key)

	fmt.Print("Decrypted Prose:\n", decrypted_prose, "\n")
}
