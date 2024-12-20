# Adi Shamir's Secret Sharing Algorithm (SSS)

## Introduction to Shamir's Secret Sharing Algorithm

**Adi Shamir's secret sharing** is a method that allows several people to **share a secret without any one of them knowing the complete secret.** To recover the secret, a certain number of shares are needed, which enables the people to work together without having to fully trust each other.

What we mean by secret key sharing is that we want to *encrypt a message* **m** and *distribute shares* to **n** people. We then want **k** of these n people with \( k < n \) to be able to *combine their shares to recover the message m*.

### Example

Imagine a family of four people who own a Bitcoin wallet. This wallet has a single key that all members can use. If each member keeps a copy of the key, a thief could steal the key from just one of them and access all the money.

If only one member keeps the key, they might lose it or decide not to share it.

Fortunately, one of the family members is an expert in cryptography. Instead of sharing the key directly, they use *Shamir's secret sharing*. They create four shares and decide that at least three shares are needed to recover the key. Here’s how this changes things:

- The key is not stored in one place, making it harder to steal.
- Members must work together to use the Bitcoin; none can betray the others.
- Even if one member loses their share or can no longer participate, the other three can still recover the key.

## Part 2: Understanding the Algorithm Through an Example

To help with understanding, let’s imagine that we are the manager of a bank and we want to protect the **code** (let's use **5887** as an example) for our **safe**. We have **5 trusted people** to whom we want to distribute shares, and if **3 of them** combine their shares, they should be able to recover the safe code. In this example, which we will use throughout this tutorial:

- \( m \) is the **safe code** (**5887**)
- \( n \) is the **number of trusted people** (**5**)
- \( k \) is the **minimum number of people needed** to recover the code (**3**)

In this key sharing process, it is important that any group of **fewer than \( k \)** people cannot gain any information about our code. Either we have **enough shares** to recover the code, or we do not, and in that case, we have **no information** about the code.

## The Naive Approach and Its Flaws

Before diving into more complex solutions, let’s consider a **simple and naive approach** to the problem and understand why we need something more powerful.

### Simple Approach

The simplest method is to **divide our key** into as many parts as necessary. For example, using the code **5887**, we could split it into four parts (5, 8, 8, and 7) and give each digit to a different person. These individuals could then come together to retrieve the original code.

### Flaws of the Naive Approach

1. **Loss of Fragments**:
   - If any one of the four people **loses their fragment** or passes away, the original code is **lost forever**. This makes recovery impossible.

2. **Information Leakage**:
   - Receiving a fragment provides **critical information** about the original code. Combining two or three fragments increases knowledge about the code. With three fragments, there’s even a **1 in 10 chance** of guessing the code correctly. This scenario is highly undesirable, as we aim to reveal information only when all fragments are present.

## Useful Functions for the Algorithm

### Polynomial Functions

#### Definitions and Examples

- **Polynomial Function**: A polynomial function is an expression of the form \( P(x) = a_n x^n + a_{n-1} x^{n-1} + \ldots + a_1 x + a_0 \), where \( a_i \) are coefficients and \( n \) is a non-negative integer.

- **Examples**:
  - \( P(x) = 2x^2 + 3x + 1 \) (degree 2 polynomial)
  - \( P(x) = 5 \) (constant polynomial)

- **Counterexamples**:
  - \( P(x) = e^x \) (exponential function, not polynomial)
  - \( P(x) = \sqrt{x} \) (square root function, not polynomial)

#### Mathematical Notations

- **Sum**: \( \sum_{i=1}^{n} a_i \) means the sum of \( a_i \) from \( i=1 \) to \( n \).
- **Product**: \( \prod_{i=1}^{n} a_i \) means the product of \( a_i \) from \( i=1 \) to \( n \).

#### Pseudo-code of the Algorithm

```pseudo
function polynomial(x, coefficients):
    result = 0
    for i from 0 to length(coefficients) - 1:
        result += coefficients[i] * (x ^ i)
    return result
```

## Lagrange Interpolation

### Definition

Lagrange interpolation is a method for constructing a polynomial that passes through a given set of points. The polynomial is of the form:
\[ P(x) = \sum_{i=0}^{n} y_i L_i(x) \]
where \( L_i(x) \) is the Lagrange polynomial given by:
\[ L_i(x) = \prod_{j=0, j \neq i}^{n} \frac{x - x_j}{x_i - x_j} \]

### Uniqueness and Existence Property

- **Uniqueness**: For a given number of distinct points (x, y), there exists a unique polynomial of degree at most \( n-1 \) that passes through these points.
- **Existence**: If we have \( n \) distinct points, it is always possible to find a polynomial that interpolates them.

### Applied Example

Consider the secret **5887** shared between **5 people** (A, B, C, D, E) with a minimum of **3** needed for reconstruction.

1. **Points**:
   - A: (1, 5)
   - B: (2, 8)
   - C: (3, 8)
   - D: (4, 7)
   - E: (5, 7)

2. **Polynomial**:
   Use Lagrange interpolation to find a polynomial \( P(x) \) that passes through these points.

## Generation of Fragments

To generate the fragments, a polynomial of degree \( k-1 \) is created, where \( k \) is the minimum number of participants required to reconstruct the secret. The values of the polynomial are then evaluated for the \( n \) points.

### Steps for Fragment Generation

#### Construction of the Polynomial Function

We will construct a **polynomial function** of degree \( k - 1 \), where \( k \) is the minimum number of keys needed to reconstruct the secret.

#### Form of the Function

The polynomial function is written as:
\[ P(X) = \sum_{i=0}^{k-1} a_i X^i \]
where **\( a_0 \) is set as the secret code \( m \)**.

#### Choice of Coefficients

- The coefficients \( a_1, a_2, \ldots, a_{k-1} \) are chosen randomly from a set \( R \).
- For our example, with \( m = 5887 \), \( k = 3 \) (which means a degree 2 polynomial), we can choose:
  - \( a_1 = 1689 \)
  - \( a_2 = 250 \)

Thus, we have:
\[ P(X) = 250X^2 + 1689X + 5887 \]

#### Generation of Fragments

We will now calculate the fragments by evaluating the polynomial for incremental values of \( x \), starting from 1.

#### Calculation of Fragments

- **Fragment 1**:
  \[ P(1) = 250(1)^2 + 1689(1) + 5887 = 7826 \]
  - **Key 1**: \( (1, 7826) \)

- **Fragment 2**:
  \[ P(2) = 250(2)^2 + 1689(2) + 5887 = 10265 \]
  - **Key 2**: \( (2, 10265) \)

- **Fragment 3**:
  \[ P(3) = 250(3)^2 + 1689(3) + 5887 = 13204 \]
  - **Key 3**: \( (3, 13204) \)

- **Fragment 4**:
  \[ P(4) = 250(4)^2 + 1689(4) + 5887 = 16643 \]
  - **Key 4**: \( (4, 16643) \)

- **Fragment 5**:
  \[ P(5) = 250(5)^2 + 1689(5) + 5887 = 20582 \]
  - **Key 5**: \( (5, 20582) \)

#### Distribution of Fragments

The generated fragments, namely:
- \( (1, 7826) \)
- \( (2, 10265) \)
- \( (3, 13204) \)
- \( (4, 16643) \)
- \( (5, 20582) \)

can now be distributed to the 5 trusted people. None of these fragments, taken individually, reveal any information about the secret code \( m \). This mechanism ensures the security of the secret as long as fewer than \( k \) fragments are possessed by any single person or malicious group.

#### Pseudo-code for Generation

```pseudo
function generate_shares(secret, k, n):
    coefficients = random_coefficients(k - 1, secret)
    shares = []
    for i from 1 to n:
        share = (i, polynomial(i, coefficients))
        shares.append(share)
    return shares
```

## Recovery of the Code from Fragments

To recover the secret, the participants with fragments must use Lagrange interpolation to reconstruct the polynomial.

#### Pseudo-code for Recovery

```pseudo
function recover_secret(shares):
    return interpolate(shares)
```

### Why is the Secret Code or Message \( m \) in the Polynomial?

In Shamir's secret sharing algorithm, the term \( m \) (which represents the secret code) is included in the polynomial \( P(X) \) for several essential reasons:

1. **Starting Point of the Polynomial**
   - **Interception**: The constant term \( a_0 = m \) is the secret itself. By placing it in the polynomial, we ensure that when participants evaluate the polynomial at different points, the secret is intrinsically linked to these evaluations.

2. **Polynomial Function**
   - The polynomial \( P(X) \) is defined as a function that must pass through the point \( (0, m) \). This means that when the polynomial is evaluated at \( X = 0 \), it always returns the value of the secret.

3. **Security**
   - Including \( m \) in the polynomial ensures that, even if a participant receives a fragment (an evaluation of the polynomial at a certain \( x \)), they cannot determine the secret unless they possess enough fragments (at least \( k \) to reconstruct the polynomial).

4. **Reconstruction**
   - During reconstruction, participants use the fragments (evaluations of the polynomial at different points) to interpolate the entire polynomial. The term \( m \) will be necessary to retrieve the value of the secret from the interpolation.

## Method with Modular Arithmetic

The algorithm often uses modular arithmetic to avoid excessively large numbers and ensure security.

### Importance of Modular Arithmetic

Modular arithmetic is essential for ensuring the **correctness** and **security** of Shamir's algorithm, even if it slightly complicates the calculations. Although this section is more technical, it does not change the basic concepts we have discussed previously.

### Choice of Coefficients

- The coefficients \( a_1, a_2, \ldots, a_{k-1} \) cannot be chosen uniformly from an infinite set \( R \).
- To address this, we use the set of **integers modulo \( p \)**, denoted \( \mathbb{Z}_p \), where \( p \) is a prime number.

#### Modulo Operation

- **Definition**: The modulo operation returns the remainder of a division. For example, \( 17 \mod 5 = 2 \).
- In modular arithmetic, we limit ourselves to numbers between \( 0 \) and \( p-1 \).

#### Examples of Operations in \( \mathbb{Z}_5 \)

- \((17 + 11) \mod 5 = 3\)
- \((3 \times 7) \mod 5 = 1\)

### Choice of a Prime Number

For Lagrange interpolation to work correctly in modular arithmetic, it is crucial to choose a number \( p \) that is **prime** and greater than the secret \( m \). This ensures the validity of the operations.

### Impact on the Algorithm

Let's take a prime number \( p = 6301 \). We will recalculate our keys using the formula:
\[ P(X) = 250X^2 + 1689X + 5887 \mod 6301 \]

#### Calculation of Keys

We generate the keys as follows:
- **Key 1**: \( P(1) = 1525 \) ⇒ \( (1, 1525) \)
- **Key 2**: \( P(2) = 3964 \) ⇒ \( (2, 3964) \)
- **Key 3**: \( P(3) = 602 \) ⇒ \( (3, 602) \)
- **Key 4**: \( P(4) = 4041 \) ⇒ \( (4, 4041) \)
- **Key 5**: \( P(5) = 1679 \) ⇒ \( (5, 1679) \)

#### Recovery of the Secret

When we provide certain keys (e.g., key1, key2, key4), the recovery algorithm uses:
\[ P(X) = y_0 l_0(X) + y_1 l_1(X) + y_2 l_2(X) \mod 6301 \]
where \( l_i(X) \) are the Lagrange polynomials.

Using modular arithmetic ensures that calculations remain within a finite set, making the choice of random and uniform coefficients possible. With this approach, we can recover the correct polynomial and, consequently, the initial secret.

### Application

- Choose a large prime number \( p \).
- Perform all operations (sum, product) modulo \( p \).

#### Example of Pseudo-code

```pseudo
function polynomial_mod(x, coefficients, p):
    result = 0
    for i from 0 to length(coefficients) - 1:
        result = (result + coefficients[i] * (x ^ i)) % p
    return result
```

## How It Works for Data Exchange on the Internet

Shamir's Secret Sharing can effectively enhance the security of data exchanged over the Internet. Here’s how it works:

### Data Encryption

- Before sharing sensitive data (like passwords or encryption keys), the data is first encrypted. This ensures that even if the data is intercepted, it remains unreadable without the necessary keys.

### Share Distribution

- The encrypted data is divided into **shares**. For example, if we have a secret key or sensitive information, it can be split into several parts.
- Each part is distributed to different participants or servers. This means that no single entity has access to the complete data.

### Threshold Requirement

- A threshold is set (e.g., \( k \) out of \( n \) shares). This means that only a specific number of shares are required to reconstruct the original data.
- For instance, if there are 5 shares and the threshold is 3, any group of 3 shares can combine to recover the original data.

### Data Recovery

- When the authorized parties need to access the data, they must combine their shares. If they meet the threshold (at least \( k \) shares), they can reconstruct the original data.
- This ensures that even if some shares are lost or if certain parties are unavailable, the data can still be accessed as long as the threshold is met.

### Security Assurance

- The key advantage of this method is that any group of fewer than \( k \) participants cannot gain any information about the original data. This enhances security significantly.
- Even if an attacker intercepts some shares, without the required number of shares, they cannot reconstruct the original data.

### Real-World Applications

- **Secure Communication**: In secure messaging apps, users can share encryption keys using secret sharing to ensure that no single party can access the key alone.
- **Cloud Storage**: Sensitive files can be split into shares stored across multiple servers, ensuring that no single server has access to the complete file.
- **Cryptographic Protocols**: Many cryptographic protocols use secret sharing as a foundation for secure multi-party computations.

To illustrate how data exchange with Shamir's Secret Sharing would work among 5 computers connected to the Internet from different countries, let's break it down step by step:

### Scenario Overview

We have 5 computers (let's call them A, B, C, D, and E) located in different countries. They want to securely share a secret (e.g., an encryption key or sensitive data).

### Step-by-Step Process

1. **Data Preparation**
   - The owner of the secret (let's say Computer A) prepares the sensitive data (e.g., a secret key) that needs to be shared securely.

2. **Secret Sharing**
   - Computer A uses Shamir's Secret Sharing algorithm to divide the secret into **5 shares**. For example, if the secret is \( K \), it might generate shares like \( S_1, S_2, S_3, S_4, S_5 \).
   - A threshold \( k \) is set, which in this case is say **3**. This means that any group of 3 computers can reconstruct the original secret.

3. **Share Distribution**
   - Computer A securely sends each share to the other computers:
     - \( S_1 \) to Computer A (self)
     - \( S_2 \) to Computer B
     - \( S_3 \) to Computer C
     - \( S_4 \) to Computer D
     - \( S_5 \) to Computer E
   - This can be done using secure channels (e.g., **TLS/SSL**) to prevent interception.

4. **Data Recovery Process**
   - If, at any point, 3 of the computers (e.g., B, C, and D) need to access the secret, they will combine their shares:
     - Each computer uses its share to reconstruct the original secret using the secret sharing algorithm.
     - Since they meet the threshold of 3, they can successfully recover the original secret.

5. **Security Assurance**
   - If any 2 computers try to collaborate (e.g., B and C), they would not be able to reconstruct the secret since they do not have enough shares.
   - If a malicious actor intercepts one or more shares, they will gain no useful information unless they have at least 3 shares.

### Benefits of This Approach

- **Redundancy**: Even if one computer goes offline or loses its share, the secret can still be reconstructed as long as at least 3 shares are available.
- **Security**: No single computer has enough information to deduce the secret on its own, minimizing the risk of data breaches.
- **Geographic Distribution**: The computers can be in different countries, leveraging the global nature of the Internet while maintaining security.

## How TLS/SSL Works

**Transport Layer Security (TLS)** and its predecessor, **Secure Sockets Layer (SSL)**, are cryptographic protocols designed to provide secure communication over a computer network. Here’s a breakdown of how they work:

### Establishing a Secure Connection

The process of establishing a secure connection using TLS/SSL typically involves several steps:

#### Handshake

- **Client Hello**: The client (e.g., a web browser) sends a "Client Hello" message to the server. This message includes:
  - The TLS/SSL version supported by the client.
  - A randomly generated number (nonce).
  - A list of supported cipher suites (encryption algorithms).

- **Server Hello**: The server responds with a "Server Hello" message, which includes:
  - The chosen TLS/SSL version.
  - A randomly generated number from the server.
  - The cipher suite selected from the client's list.

#### Server Authentication and Pre-Master Secret

- The server sends its **digital certificate**, which contains its public key and is signed by a trusted Certificate Authority (CA). The client verifies this certificate to ensure it’s communicating with the legitimate server.

- The client generates a **pre-master secret**, encrypts it using the server's public key, and sends it to the server. Only the server can decrypt this message with its private key.

#### Session Keys Creation

- Both the client and server use the pre-master secret along with the two nonce values (from the Client Hello and Server Hello) to generate a set of **session keys**. These keys are symmetric keys used for encrypting the data transmitted during the session.

### Secure Data Transmission

Once the handshake is complete, a secure encrypted session is established:

- **Data Encryption**: All data exchanged between the client and server is encrypted using the session keys. This ensures confidentiality, meaning that even if data is intercepted, it cannot be read without the keys.

- **Data Integrity**: TLS/SSL also ensures the integrity of the data using **message authentication codes (MAC)**. This verifies that the data has not been altered in transit.

- **Sequence Number**: A sequence number is used to prevent replay attacks. Each message includes a sequence number to ensure that messages are processed in order and not duplicated.

### Connection Termination

Once the data exchange is complete, the connection can be securely terminated. This involves:

- The client and server send a "close_notify" alert to each other to signal that the connection will be closed.
- Both parties discard any session keys to ensure that the session cannot be resumed.

### Benefits of TLS/SSL

- **Secure Communication**: Encrypts data in transit, protecting sensitive information from eavesdropping.
- **Authentication**: Verifies the identity of the parties involved, ensuring that clients are communicating with legitimate servers.
- **Data Integrity**: Ensures that the data has not been tampered with during transmission.
```
