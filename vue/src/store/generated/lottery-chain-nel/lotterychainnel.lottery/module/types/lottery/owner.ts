/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "lotterychainnel.lottery";

export interface Owner {}

const baseOwner: object = {};

export const Owner = {
  encode(_: Owner, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Owner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOwner } as Owner;
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

  fromJSON(_: any): Owner {
    const message = { ...baseOwner } as Owner;
    return message;
  },

  toJSON(_: Owner): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<Owner>): Owner {
    const message = { ...baseOwner } as Owner;
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
