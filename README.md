# GNARK Circuits

    some zero knowledge circuit implementations using Gnark framework

## Circuits

### HTTP Circuit

- A circuit which requests an endpoint for some data and uses that data for some operations
- For this to work as expected, the api needs to DETERMINISTIC (an endpoint should return same reponse no matter how many times it's called) !
- It's like a pure function
- If the endpoint returns different response, it will lead inconsistency error when compiling the circuit

### Merkle Inclusion proof

- An inclusion proof verification circuit
- Data (index of data segment to be proven, IS THERE A WAY TO VERIFY IF THE LEAF DATA MATCHES DATA SEGMENT AT INDEX??), merkleProof as inputs

      merkle.verify(data, merkleproof)

### Mimc Hash Function

- A simple Mimc Hash function verifier with BN254 curve

      Hash(data) == Expected Hash

### Signature Verification

- An eddsa signature verfication that checks if a digital signature is valid or not
  pubkey.verify(data, signature)
- data is private input and pubkey, signature are public inputs

      Note: Signature circuit is refactored to have separate prover and verifier. It could be used as a complete example flow.

### Sha2 vs crypto/sha256

- Tried to check hashing compatibility of gnark sha2 implementation with crypto/256. Failed!

### Comarisions

- To check if we can directly compare circuit variables without using assert. Failed!

### Conditional Circuit

- A Zk circuit doesn't support if statement, but we can use api.Select{} for somewhat similar effect

```
// this is equivalent to toAdd1 := circuit.A_Support == 1 ? circuit.A : 0
toAdd1 := api.Select(circuit.A_Support, circuit.A, 0)
```

### Cube checker

- A simple circuit that check cube of a number

      x^3 == y

- X is the private input and y is the public input

### Recursive proof (Not completed)

![recursive proof design](https://github.com/Teja2045/GNARK-Circuits/assets/106052623/30482e17-57ff-41ac-bc54-a1cdc22f956d)
