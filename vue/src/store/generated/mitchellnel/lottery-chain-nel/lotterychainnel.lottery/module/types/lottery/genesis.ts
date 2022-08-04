/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../lottery/params";
import { Owner } from "../lottery/owner";
import { EntranceFee } from "../lottery/entrance_fee";
import { LotteryState } from "../lottery/lottery_state";
import { Player } from "../lottery/player";
import { LastWinner } from "../lottery/last_winner";

export const protobufPackage = "lotterychainnel.lottery";

/** GenesisState defines the lottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  owner: Owner | undefined;
  entranceFee: EntranceFee | undefined;
  lotteryState: LotteryState | undefined;
  playerList: Player[];
  playerCount: number;
  /** this line is used by starport scaffolding # genesis/proto/state */
  lastWinner: LastWinner | undefined;
}

const baseGenesisState: object = { playerCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.owner !== undefined) {
      Owner.encode(message.owner, writer.uint32(18).fork()).ldelim();
    }
    if (message.entranceFee !== undefined) {
      EntranceFee.encode(
        message.entranceFee,
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.lotteryState !== undefined) {
      LotteryState.encode(
        message.lotteryState,
        writer.uint32(34).fork()
      ).ldelim();
    }
    for (const v of message.playerList) {
      Player.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.playerCount !== 0) {
      writer.uint32(48).uint64(message.playerCount);
    }
    if (message.lastWinner !== undefined) {
      LastWinner.encode(message.lastWinner, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.playerList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.owner = Owner.decode(reader, reader.uint32());
          break;
        case 3:
          message.entranceFee = EntranceFee.decode(reader, reader.uint32());
          break;
        case 4:
          message.lotteryState = LotteryState.decode(reader, reader.uint32());
          break;
        case 5:
          message.playerList.push(Player.decode(reader, reader.uint32()));
          break;
        case 6:
          message.playerCount = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.lastWinner = LastWinner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.playerList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = Owner.fromJSON(object.owner);
    } else {
      message.owner = undefined;
    }
    if (object.entranceFee !== undefined && object.entranceFee !== null) {
      message.entranceFee = EntranceFee.fromJSON(object.entranceFee);
    } else {
      message.entranceFee = undefined;
    }
    if (object.lotteryState !== undefined && object.lotteryState !== null) {
      message.lotteryState = LotteryState.fromJSON(object.lotteryState);
    } else {
      message.lotteryState = undefined;
    }
    if (object.playerList !== undefined && object.playerList !== null) {
      for (const e of object.playerList) {
        message.playerList.push(Player.fromJSON(e));
      }
    }
    if (object.playerCount !== undefined && object.playerCount !== null) {
      message.playerCount = Number(object.playerCount);
    } else {
      message.playerCount = 0;
    }
    if (object.lastWinner !== undefined && object.lastWinner !== null) {
      message.lastWinner = LastWinner.fromJSON(object.lastWinner);
    } else {
      message.lastWinner = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.owner !== undefined &&
      (obj.owner = message.owner ? Owner.toJSON(message.owner) : undefined);
    message.entranceFee !== undefined &&
      (obj.entranceFee = message.entranceFee
        ? EntranceFee.toJSON(message.entranceFee)
        : undefined);
    message.lotteryState !== undefined &&
      (obj.lotteryState = message.lotteryState
        ? LotteryState.toJSON(message.lotteryState)
        : undefined);
    if (message.playerList) {
      obj.playerList = message.playerList.map((e) =>
        e ? Player.toJSON(e) : undefined
      );
    } else {
      obj.playerList = [];
    }
    message.playerCount !== undefined &&
      (obj.playerCount = message.playerCount);
    message.lastWinner !== undefined &&
      (obj.lastWinner = message.lastWinner
        ? LastWinner.toJSON(message.lastWinner)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.playerList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = Owner.fromPartial(object.owner);
    } else {
      message.owner = undefined;
    }
    if (object.entranceFee !== undefined && object.entranceFee !== null) {
      message.entranceFee = EntranceFee.fromPartial(object.entranceFee);
    } else {
      message.entranceFee = undefined;
    }
    if (object.lotteryState !== undefined && object.lotteryState !== null) {
      message.lotteryState = LotteryState.fromPartial(object.lotteryState);
    } else {
      message.lotteryState = undefined;
    }
    if (object.playerList !== undefined && object.playerList !== null) {
      for (const e of object.playerList) {
        message.playerList.push(Player.fromPartial(e));
      }
    }
    if (object.playerCount !== undefined && object.playerCount !== null) {
      message.playerCount = object.playerCount;
    } else {
      message.playerCount = 0;
    }
    if (object.lastWinner !== undefined && object.lastWinner !== null) {
      message.lastWinner = LastWinner.fromPartial(object.lastWinner);
    } else {
      message.lastWinner = undefined;
    }
    return message;
  },
};

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
