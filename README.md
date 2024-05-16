# GNARK Circuits
    some zero knowledge circuit implementations using Gnark framework

    Note: Signature circuit is refactored to have separate prover and verifier. It could be used as a complete example flow.

## Circuits

### Merkle Inclusion proof
- An inclusion proof verification circuit
- Data (index of data segment to be proven, IS THERE A WAY TO VERIFY IF THE LEAF DATA MATCHES DATA SEGMENT AT INDEX??), merkleProof as inputs
  
      merkle.verify(data, merkleproof)

### Mimc Hash Function
- A simple Mimc Hash function verifier with BN254 curve

      Hash(data) == Expected Hash

### Cube checker
- A simple circuit that check cube of a number
  
      x^3 == y
- X is the private input and y is the public input  

### Signature Verification
- An eddsa signature verfication that checks if a digital signature is valid or not
      pubkey.verify(data, signature)
- data is private input and pubkey, signature are public inputs
