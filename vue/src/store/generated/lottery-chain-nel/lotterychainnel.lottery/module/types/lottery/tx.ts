/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

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
  entranceFee: string;
}

export interface MsgSetupLotteryResponse {
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

const baseMsgSetupLottery: object = { creator: "", entranceFee: "" };

export const MsgSetupLottery = {
  encode(message: MsgSetupLottery, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.entranceFee !== "") {
      writer.uint32(18).string(message.entranceFee);
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
          message.entranceFee = reader.string();
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
      message.entranceFee = String(object.entranceFee);
    } else {
      message.entranceFee = "";
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
      message.entranceFee = "";
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

/** Msg defines the Msg service. */
export interface Msg {
  ClaimOwner(request: MsgClaimOwner): Promise<MsgClaimOwnerResponse>;
  ChangeOwner(request: MsgChangeOwner): Promise<MsgChangeOwnerResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetupLottery(request: MsgSetupLottery): Promise<MsgSetupLotteryResponse>;
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
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
