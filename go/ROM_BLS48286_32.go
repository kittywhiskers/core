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

/* Fixed Data in ROM - Field and Curve parameters */

package BLS48286

// Base Bits= 29
var Modulus= [...]Chunk {0x9C345B,0x13A815C9,0x38D4B67,0xC8388A5,0x4809AAF,0x63F0D60,0x516A1B2,0x16040390,0x11EC7BD7,0x14B9629}
var R2modp= [...]Chunk {0x8484275,0x16B4A09C,0x13973525,0x19DBD350,0x188DE1C0,0x111BD51D,0x113A0F92,0x6489CAD,0xF188E5E,0xCED0C0}
var ROI= [...]Chunk {0x9C345A,0x13A815C9,0x38D4B67,0xC8388A5,0x4809AAF,0x63F0D60,0x516A1B2,0x16040390,0x11EC7BD7,0x14B9629}
var SQRTm3= [...]Chunk {0x6524A9C,0x6BD2397,0x14CA2663,0x6B70283,0x84EB335,0xC4C4D38,0x61EE65D,0xE146445,0x1478F5D8,0x14B8E12}
var CRu= [...]Chunk {0x3773F7B,0xD329CB0,0xC2BB8E5,0x99D4594,0x667A6F2,0x1945AD4C,0x159AC407,0x20C33EA,0x332B8D8,0x14B921E}
const MConst Chunk=0x16EA242D
var Fra= [...]Chunk {0x1FC90183,0x15434AF9,0xF4FA4DB,0xAAD7DA3,0x8F09A7C,0x81F7813,0x1F8010F6,0x9D15D85,0xC522C11,0x6CA0BF}
var Frb= [...]Chunk {0xD332D8,0x1E64CACF,0x143DA68B,0x1D60B01,0x1B900033,0x1E1F954C,0x59690BB,0xC32A60A,0x59A4FC6,0xDEF56A}

//*** rom curve parameters *****
// Ate Bits= 17
// G2 Table size= 20
const CURVE_Cof_I int= 62958
var CURVE_Cof= [...]Chunk {0xF5EE,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
const CURVE_B_I int= 10
var CURVE_B= [...]Chunk {0xA,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
var CURVE_Order= [...]Chunk {0x1D612C81,0x1E208D97,0x1CA5C07F,0x11E51045,0x1E4FE229,0x13154A7D,0x127B79AC,0x1602A6BF,0x86BC9E,0x0}
var CURVE_Gx= [...]Chunk {0x95D59E0,0x44BF518,0x1BEE8577,0x1B2E8EB9,0xE84D19D,0xA23A0D2,0x103C1301,0x111F68ED,0x13D120DB,0x34B921}
var CURVE_Gy= [...]Chunk {0x10ABEB43,0x177B4FCB,0x38F8BB,0x6F9CD6B,0x1B0AC1F2,0x1846F729,0x1E5E6A8C,0x5AF112F,0xA1E656E,0x101C720}
var CURVE_HTPC= [...]Chunk {0x1,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}

var CURVE_Bnx= [...]Chunk {0xF5EF,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0,0x0}
var CURVE_Pxaaa= [...]Chunk {0x56106B5,0x20B1A4C,0x113303FA,0x187639D9,0x1D9F51AC,0xA96D657,0x132BADEA,0x1A544F60,0x531C71,0x11567EA}
var CURVE_Pxaab= [...]Chunk {0x96D1A55,0xF668DEC,0xE2476A7,0x4EE30AA,0x1EF56B94,0x5192C97,0x121CF932,0x14B1A4A1,0x196B5C85,0xC6DDCC}
var CURVE_Pxaba= [...]Chunk {0x7D7A35E,0x408D1D3,0x1162980E,0x1F5E9D19,0x12AF41C9,0xCD2835A,0x17875370,0x62F5C57,0x15DE0A7C,0x541496}
var CURVE_Pxabb= [...]Chunk {0x2F023E8,0x14742906,0xCF4B428,0x3236B9E,0x1B62D89F,0x192E7AEF,0x1E4BA079,0x19B17763,0x155FAC94,0x1015F06}
var CURVE_Pxbaa= [...]Chunk {0xF14CB9B,0x5CC6ADC,0x18955362,0x142FAC5F,0x11B96A21,0x18F7B6AA,0x96591F1,0xC5096A4,0x9A1E3F7,0xBA939B}
var CURVE_Pxbab= [...]Chunk {0xB60DD54,0x1FB59436,0x1C68910E,0x10A65726,0x6BE9C0B,0x1C6AA6BF,0x14DD8358,0x1E68D885,0x85E9D59,0x880F52}
var CURVE_Pxbba= [...]Chunk {0x19D776DA,0x1052D855,0x844D7E8,0xCAE4EC2,0x4ADABAA,0x120A2C5F,0x192AB537,0x3F9DEC1,0xE261FE4,0x7B037E}
var CURVE_Pxbbb= [...]Chunk {0x8B50A8E,0x19FF11C1,0x18273AD7,0xEAA9DE4,0x1EC9A57,0x1670C91D,0x300009C,0xACD8A2D,0x6920D19,0xF85987}
var CURVE_Pyaaa= [...]Chunk {0x7E988BE,0x1366AFEB,0x5E2F3B8,0x1FCBDF1D,0x114E8B31,0x1CD92A27,0x1A53FC0,0xAAF8FE7,0xDA962CD,0x10F87C6}
var CURVE_Pyaab= [...]Chunk {0xDF11B92,0x1BC2B368,0x1AF17821,0x6663003,0xDB40CB8,0x20094F8,0x1856186E,0x1528744B,0x1EC6FED7,0x700E0F}
var CURVE_Pyaba= [...]Chunk {0xBFDD06A,0x3E7D33D,0x13F96495,0x179B5611,0x12EB8E64,0xD7054C6,0xD7BFC95,0x145B3D76,0x18C3D6BE,0x9607FD}
var CURVE_Pyabb= [...]Chunk {0x188A47F4,0x19EEFE9B,0x41ACCA2,0x12BF346C,0x11F157D1,0x19429C28,0x14D4ACD2,0x743BB7B,0x55BDFFF,0xAB418F}
var CURVE_Pybaa= [...]Chunk {0x5E39E77,0x1B76DB9A,0x54E3132,0x190321FF,0x5119324,0x123E6F79,0x1064FB7D,0x1D69EBDE,0xAA01884,0x39C720}
var CURVE_Pybab= [...]Chunk {0x164975C6,0x194140A1,0xA2302C0,0x1C495B39,0x613723B,0xB4D86D7,0x1A6F8686,0x7238A06,0x6896E24,0x9B2C3B}
var CURVE_Pybba= [...]Chunk {0x1BF2D7C,0x1E99CD5E,0xEA89102,0xE27DE36,0x13837F1B,0xA2D04CE,0x23E06D8,0x18BDC695,0x175143C9,0xFD8541}
var CURVE_Pybbb= [...]Chunk {0x51130A6,0x154D068F,0x18954F6A,0x173CE106,0x7C687EE,0xD23B3B5,0xB58CC3B,0x1342EEF6,0xDC6AC9E,0xC9F9BC}
