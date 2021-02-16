/*
 * Copyright (c) 2012-2020 MIRACL UK Ltd.
 *
 * This file is part of MIRACL Core
 * (see https://github.com/miracl/core).
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* CORE x509 header file */

/**
 * @file x509.h
 * @author Mike Scott and Kealan McCusker
 * @date 19th May 2015
 * @brief X509 function Header File
 *
 * defines structures
 * declares functions
 *
 */

#ifndef X509_H
#define X509_H

// Supported Encryption Methods

#define X509_ECC 1
#define X509_RSA 2

// Supported Hash functions

#define X509_H256 2
#define X509_H384 3
#define X509_H512 4

// Supported Curves

#define USE_NIST256 0    /**< For the NIST 256-bit standard curve - WEIERSTRASS only */
#define USE_C25519 1     /**< Bernstein's Modulus 2^255-19 - EDWARDS or MONTGOMERY only */
#define USE_BRAINPOOL 2  /**< For Brainpool 256-bit curve - WEIERSTRASS only */
#define USE_ANSSI 3      /**< For French 256-bit standard curve - WEIERSTRASS only */
#define USE_NIST384 10   /**< For the NIST 384-bit standard curve - WEIERSTRASS only */
#define USE_NIST521 12   /**< For the NIST 521-bit standard curve - WEIERSTRASS only */

// Some ASN.1 tags

#define ANY 0x00
#define SEQ 0x30
#define OID 0x06
#define INT 0x02
#define NUL 0x05
#define ZER 0x00
#define UTF 0x0C
#define UTC 0x17
#define LOG 0x01
#define BIT 0x03
#define OCT 0x04
#define STR 0x13
#define SET 0x31
#define IA5 0x16
#define EXT 0xA3
#define DNS 0x82

using namespace core;

extern octet X509_CN;  // Country Name
extern octet X509_ON;  // organisation Name
extern octet X509_EN;  // email
extern octet X509_LN;  // local name
extern octet X509_UN;  // Unit name (aka Organisation Unit OU)
extern octet X509_MN;  // My Name (aka Common Name)
extern octet X509_SN;  // State Name

extern octet X509_AN;  // Alternate Name
extern octet X509_KU;  // Key Usage

/**
 * @brief Public key type
 */
typedef struct
{
    int type;  /**< signature type (ECC or RSA) */
    int hash;  /**< hash type */
    int curve; /**< elliptic curve used or RSA key length in bits  */
} pktype;


/* X.509 functions */
/** @brief Extract certificate signature
 *
	@param c an X.509 certificate
	@param s the extracted signature
	@return 0 on failure, or indicator of signature type (ECC or RSA)

*/
extern pktype X509_extract_cert_sig(octet *c, octet *s);
/** @brief
 *
	@param sc a signed certificate
	@param c the extracted certificate
	@return 0 on failure
*/
extern int X509_extract_cert(octet *sc, octet *c);
/** @brief
 *
	@param c an X.509 certificate
	@param k the extracted key
	@return 0 on failure, or indicator of public key type (ECC or RSA)
*/
extern pktype X509_extract_public_key(octet *c, octet *k);
/** @brief
 *
	@param c an X.509 certificate
	@return 0 on failure, or pointer to issuer field in cert
*/
extern int X509_find_issuer(octet *c);
/** @brief
 *
	@param c an X.509 certificate
	@return 0 on failure, or pointer to validity field in cert
*/
extern int X509_find_validity(octet *c);
/** @brief
 *
	@param c an X.509 certificate
	@return 0 on failure, or pointer to subject field in cert
*/
extern int X509_find_subject(octet *c);

/** @brief
 *
	@param c an X.509 certificate
	@return true if self-signed, else false
*/
extern int X509_self_signed(octet *c);

/** @brief
 *
	@param c an X.509 certificate
	@param S is OID of property we are looking for
	@param s is a pointer to the section of interest in the cert
	@param f is pointer to the length of the property
	@return 0 on failure, or pointer to the property
*/
extern int X509_find_entity_property(octet *c, octet *S, int s, int *f);
/** @brief
 *
	@param c an X.509 certificate
	@param s is a pointer to the start of the validity field
	@return 0 on failure, or pointer to the start date
*/
extern int X509_find_start_date(octet *c, int s);
/** @brief
 *
	@param c an X.509 certificate
	@param s is a pointer to the start of the validity field
	@return 0 on failure, or pointer to the expiry date
*/
extern int X509_find_expiry_date(octet *c, int s);


/** @brief
 *
	@param c an X.509 certificate
	@return 0 on failure (or no extensions), or pointer to extensions field in cert
*/
extern int X509_find_extensions(octet *c);
/** @brief
 *
	@param c an X.509 certificate
	@param S is OID of particular extension we are looking for
	@param s is a pointer to the section of interest in the cert
	@param f is pointer to the length of the extension
	@return 0 on failure, or pointer to the extension
*/
extern int X509_find_extension(octet *c, octet *S, int s, int *f);

/** @brief
 *
	@param c an X.509 certificate
    @param s is a pointer to certificate extension SubjectAltNames
    @param name is a URL
    @return 0 on failure, 1 if URL is in list of alt names
*/
extern int X509_find_alt_name(octet *c,int s,char *name);

#endif
