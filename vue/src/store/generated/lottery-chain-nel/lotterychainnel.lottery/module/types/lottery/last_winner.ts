/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "lotterychainnel.lottery";

export interface LastWinner {}

const baseLastWinner: object = {};

export const LastWinner = {
  encode(_: LastWinner, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LastWinner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLastWinner } as LastWinner;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): LastWinner {
    const message = { ...baseLastWinner } as LastWinner;
    return message;
  },

  toJSON(_: LastWinner): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<LastWinner>): LastWinner {
    const message = { ...baseLastWinner } as LastWinner;
    return message;
  },
};

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
