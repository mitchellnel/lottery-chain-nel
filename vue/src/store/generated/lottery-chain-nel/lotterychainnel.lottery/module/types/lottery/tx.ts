/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "lotterychainnel.lottery";

export interface MsgClaimOwner {
  creator: string;
}

export interface MsgClaimOwnerResponse {
  success: boolean;
}

export interface MsgChangeOwner {
  creator: string;
  newOwner: string;
}

export interface MsgChangeOwnerResponse {
  success: boolean;
}

export interface MsgSetupLottery {
  creator: string;
  entranceFee: number;
}

export interface MsgSetupLotteryResponse {
  success: boolean;
}

export interface MsgChangeEntranceFee {
  creator: string;
  newEntranceFee: number;
}

export interface MsgChangeEntranceFeeResponse {
  success: boolean;
}

export interface MsgEnterLottery {
  creator: string;
}

export interface MsgEnterLotteryResponse {
  success: boolean;
}

export interface MsgStartLottery {
  creator: string;
}

export interface MsgStartLotteryResponse {
  success: boolean;
}

export interface MsgEndLottery {
  creator: string;
}

export interface MsgEndLotteryResponse {
  success: boolean;
}

const baseMsgClaimOwner: object = { creator: "" };

export const MsgClaimOwner = {
  encode(message: MsgClaimOwner, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimOwner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimOwner } as MsgClaimOwner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimOwner {
    const message = { ...baseMsgClaimOwner } as MsgClaimOwner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgClaimOwner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgClaimOwner>): MsgClaimOwner {
    const message = { ...baseMsgClaimOwner } as MsgClaimOwner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgClaimOwnerResponse: object = { success: false };

export const MsgClaimOwnerResponse = {
  encode(
    message: MsgClaimOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimOwnerResponse } as MsgClaimOwnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimOwnerResponse {
    const message = { ...baseMsgClaimOwnerResponse } as MsgClaimOwnerResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgClaimOwnerResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgClaimOwnerResponse>
  ): MsgClaimOwnerResponse {
    const message = { ...baseMsgClaimOwnerResponse } as MsgClaimOwnerResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseMsgChangeOwner: object = { creator: "", newOwner: "" };

export const MsgChangeOwner = {
  encode(message: MsgChangeOwner, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newOwner !== "") {
      writer.uint32(18).string(message.newOwner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChangeOwner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChangeOwner } as MsgChangeOwner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newOwner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChangeOwner {
    const message = { ...baseMsgChangeOwner } as MsgChangeOwner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newOwner !== undefined && object.newOwner !== null) {
      message.newOwner = String(object.newOwner);
    } else {
      message.newOwner = "";
    }
    return message;
  },

  toJSON(message: MsgChangeOwner): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newOwner !== undefined && (obj.newOwner = message.newOwner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgChangeOwner>): MsgChangeOwner {
    const message = { ...baseMsgChangeOwner } as MsgChangeOwner;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newOwner !== undefined && object.newOwner !== null) {
      message.newOwner = object.newOwner;
    } else {
      message.newOwner = "";
    }
    return message;
  },
};

const baseMsgChangeOwnerResponse: object = { success: false };

export const MsgChangeOwnerResponse = {
  encode(
    message: MsgChangeOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChangeOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChangeOwnerResponse } as MsgChangeOwnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChangeOwnerResponse {
    const message = { ...baseMsgChangeOwnerResponse } as MsgChangeOwnerResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgChangeOwnerResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgChangeOwnerResponse>
  ): MsgChangeOwnerResponse {
    const message = { ...baseMsgChangeOwnerResponse } as MsgChangeOwnerResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseMsgSetupLottery: object = { creator: "", entranceFee: 0 };

export const MsgSetupLottery = {
  encode(message: MsgSetupLottery, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.entranceFee !== 0) {
      writer.uint32(16).uint64(message.entranceFee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetupLottery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSetupLottery } as MsgSetupLottery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.entranceFee = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetupLottery {
    const message = { ...baseMsgSetupLottery } as MsgSetupLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.entranceFee !== undefined && object.entranceFee !== null) {
      message.entranceFee = Number(object.entranceFee);
    } else {
      message.entranceFee = 0;
    }
    return message;
  },

  toJSON(message: MsgSetupLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.entranceFee !== undefined &&
      (obj.entranceFee = message.entranceFee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSetupLottery>): MsgSetupLottery {
    const message = { ...baseMsgSetupLottery } as MsgSetupLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.entranceFee !== undefined && object.entranceFee !== null) {
      message.entranceFee = object.entranceFee;
    } else {
      message.entranceFee = 0;
    }
    return message;
  },
};

const baseMsgSetupLotteryResponse: object = { success: false };

export const MsgSetupLotteryResponse = {
  encode(
    message: MsgSetupLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSetupLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSetupLotteryResponse,
    } as MsgSetupLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSetupLotteryResponse {
    const message = {
      ...baseMsgSetupLotteryResponse,
    } as MsgSetupLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgSetupLotteryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgSetupLotteryResponse>
  ): MsgSetupLotteryResponse {
    const message = {
      ...baseMsgSetupLotteryResponse,
    } as MsgSetupLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseMsgChangeEntranceFee: object = { creator: "", newEntranceFee: 0 };

export const MsgChangeEntranceFee = {
  encode(
    message: MsgChangeEntranceFee,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.newEntranceFee !== 0) {
      writer.uint32(16).uint64(message.newEntranceFee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgChangeEntranceFee {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgChangeEntranceFee } as MsgChangeEntranceFee;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.newEntranceFee = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChangeEntranceFee {
    const message = { ...baseMsgChangeEntranceFee } as MsgChangeEntranceFee;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.newEntranceFee !== undefined && object.newEntranceFee !== null) {
      message.newEntranceFee = Number(object.newEntranceFee);
    } else {
      message.newEntranceFee = 0;
    }
    return message;
  },

  toJSON(message: MsgChangeEntranceFee): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.newEntranceFee !== undefined &&
      (obj.newEntranceFee = message.newEntranceFee);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgChangeEntranceFee>): MsgChangeEntranceFee {
    const message = { ...baseMsgChangeEntranceFee } as MsgChangeEntranceFee;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.newEntranceFee !== undefined && object.newEntranceFee !== null) {
      message.newEntranceFee = object.newEntranceFee;
    } else {
      message.newEntranceFee = 0;
    }
    return message;
  },
};

const baseMsgChangeEntranceFeeResponse: object = { success: false };

export const MsgChangeEntranceFeeResponse = {
  encode(
    message: MsgChangeEntranceFeeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgChangeEntranceFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgChangeEntranceFeeResponse,
    } as MsgChangeEntranceFeeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgChangeEntranceFeeResponse {
    const message = {
      ...baseMsgChangeEntranceFeeResponse,
    } as MsgChangeEntranceFeeResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgChangeEntranceFeeResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgChangeEntranceFeeResponse>
  ): MsgChangeEntranceFeeResponse {
    const message = {
      ...baseMsgChangeEntranceFeeResponse,
    } as MsgChangeEntranceFeeResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseMsgEnterLottery: object = { creator: "" };

export const MsgEnterLottery = {
  encode(message: MsgEnterLottery, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnterLottery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEnterLottery } as MsgEnterLottery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnterLottery {
    const message = { ...baseMsgEnterLottery } as MsgEnterLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgEnterLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEnterLottery>): MsgEnterLottery {
    const message = { ...baseMsgEnterLottery } as MsgEnterLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgEnterLotteryResponse: object = { success: false };

export const MsgEnterLotteryResponse = {
  encode(
    message: MsgEnterLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEnterLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgEnterLotteryResponse,
    } as MsgEnterLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnterLotteryResponse {
    const message = {
      ...baseMsgEnterLotteryResponse,
    } as MsgEnterLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgEnterLotteryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgEnterLotteryResponse>
  ): MsgEnterLotteryResponse {
    const message = {
      ...baseMsgEnterLotteryResponse,
    } as MsgEnterLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseMsgStartLottery: object = { creator: "" };

export const MsgStartLottery = {
  encode(message: MsgStartLottery, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStartLottery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStartLottery } as MsgStartLottery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStartLottery {
    const message = { ...baseMsgStartLottery } as MsgStartLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgStartLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgStartLottery>): MsgStartLottery {
    const message = { ...baseMsgStartLottery } as MsgStartLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgStartLotteryResponse: object = { success: false };

export const MsgStartLotteryResponse = {
  encode(
    message: MsgStartLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStartLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgStartLotteryResponse,
    } as MsgStartLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStartLotteryResponse {
    const message = {
      ...baseMsgStartLotteryResponse,
    } as MsgStartLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgStartLotteryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgStartLotteryResponse>
  ): MsgStartLotteryResponse {
    const message = {
      ...baseMsgStartLotteryResponse,
    } as MsgStartLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

const baseMsgEndLottery: object = { creator: "" };

export const MsgEndLottery = {
  encode(message: MsgEndLottery, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEndLottery {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEndLottery } as MsgEndLottery;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEndLottery {
    const message = { ...baseMsgEndLottery } as MsgEndLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: MsgEndLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEndLottery>): MsgEndLottery {
    const message = { ...baseMsgEndLottery } as MsgEndLottery;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseMsgEndLotteryResponse: object = { success: false };

export const MsgEndLotteryResponse = {
  encode(
    message: MsgEndLotteryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEndLotteryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEndLotteryResponse } as MsgEndLotteryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEndLotteryResponse {
    const message = { ...baseMsgEndLotteryResponse } as MsgEndLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = Boolean(object.success);
    } else {
      message.success = false;
    }
    return message;
  },

  toJSON(message: MsgEndLotteryResponse): unknown {
    const obj: any = {};
    message.success !== undefined && (obj.success = message.success);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgEndLotteryResponse>
  ): MsgEndLotteryResponse {
    const message = { ...baseMsgEndLotteryResponse } as MsgEndLotteryResponse;
    if (object.success !== undefined && object.success !== null) {
      message.success = object.success;
    } else {
      message.success = false;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ClaimOwner(request: MsgClaimOwner): Promise<MsgClaimOwnerResponse>;
  ChangeOwner(request: MsgChangeOwner): Promise<MsgChangeOwnerResponse>;
  SetupLottery(request: MsgSetupLottery): Promise<MsgSetupLotteryResponse>;
  ChangeEntranceFee(
    request: MsgChangeEntranceFee
  ): Promise<MsgChangeEntranceFeeResponse>;
  EnterLottery(request: MsgEnterLottery): Promise<MsgEnterLotteryResponse>;
  StartLottery(request: MsgStartLottery): Promise<MsgStartLotteryResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EndLottery(request: MsgEndLottery): Promise<MsgEndLotteryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ClaimOwner(request: MsgClaimOwner): Promise<MsgClaimOwnerResponse> {
    const data = MsgClaimOwner.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "ClaimOwner",
      data
    );
    return promise.then((data) =>
      MsgClaimOwnerResponse.decode(new Reader(data))
    );
  }

  ChangeOwner(request: MsgChangeOwner): Promise<MsgChangeOwnerResponse> {
    const data = MsgChangeOwner.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "ChangeOwner",
      data
    );
    return promise.then((data) =>
      MsgChangeOwnerResponse.decode(new Reader(data))
    );
  }

  SetupLottery(request: MsgSetupLottery): Promise<MsgSetupLotteryResponse> {
    const data = MsgSetupLottery.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "SetupLottery",
      data
    );
    return promise.then((data) =>
      MsgSetupLotteryResponse.decode(new Reader(data))
    );
  }

  ChangeEntranceFee(
    request: MsgChangeEntranceFee
  ): Promise<MsgChangeEntranceFeeResponse> {
    const data = MsgChangeEntranceFee.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "ChangeEntranceFee",
      data
    );
    return promise.then((data) =>
      MsgChangeEntranceFeeResponse.decode(new Reader(data))
    );
  }

  EnterLottery(request: MsgEnterLottery): Promise<MsgEnterLotteryResponse> {
    const data = MsgEnterLottery.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "EnterLottery",
      data
    );
    return promise.then((data) =>
      MsgEnterLotteryResponse.decode(new Reader(data))
    );
  }

  StartLottery(request: MsgStartLottery): Promise<MsgStartLotteryResponse> {
    const data = MsgStartLottery.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "StartLottery",
      data
    );
    return promise.then((data) =>
      MsgStartLotteryResponse.decode(new Reader(data))
    );
  }

  EndLottery(request: MsgEndLottery): Promise<MsgEndLotteryResponse> {
    const data = MsgEndLottery.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Msg",
      "EndLottery",
      data
    );
    return promise.then((data) =>
      MsgEndLotteryResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
