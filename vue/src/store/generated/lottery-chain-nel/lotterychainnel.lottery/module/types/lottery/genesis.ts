/* eslint-disable */
import { Params } from "../lottery/params";
import { Owner } from "../lottery/owner";
import { EntranceFee } from "../lottery/entrance_fee";
import { LotteryState } from "../lottery/lottery_state";
import { Player } from "../lottery/player";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "lotterychainnel.lottery";

/** GenesisState defines the lottery module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  owner: Owner | undefined;
  entranceFee: EntranceFee | undefined;
  lotteryState: LotteryState | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  playerList: Player[];
}

const baseGenesisState: object = {};

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
