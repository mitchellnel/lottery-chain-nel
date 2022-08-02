/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "lotterychainnel.lottery";

export interface LotteryState {
  lottery_state: LotteryState_LOTTERY_STATE;
}

export enum LotteryState_LOTTERY_STATE {
  CLOSED = 0,
  CALCULATING_WINNER = 1,
  OPEN = 2,
  UNRECOGNIZED = -1,
}

export function lotteryState_LOTTERY_STATEFromJSON(
  object: any
): LotteryState_LOTTERY_STATE {
  switch (object) {
    case 0:
    case "CLOSED":
      return LotteryState_LOTTERY_STATE.CLOSED;
    case 1:
    case "CALCULATING_WINNER":
      return LotteryState_LOTTERY_STATE.CALCULATING_WINNER;
    case 2:
    case "OPEN":
      return LotteryState_LOTTERY_STATE.OPEN;
    case -1:
    case "UNRECOGNIZED":
    default:
      return LotteryState_LOTTERY_STATE.UNRECOGNIZED;
  }
}

export function lotteryState_LOTTERY_STATEToJSON(
  object: LotteryState_LOTTERY_STATE
): string {
  switch (object) {
    case LotteryState_LOTTERY_STATE.CLOSED:
      return "CLOSED";
    case LotteryState_LOTTERY_STATE.CALCULATING_WINNER:
      return "CALCULATING_WINNER";
    case LotteryState_LOTTERY_STATE.OPEN:
      return "OPEN";
    default:
      return "UNKNOWN";
  }
}

const baseLotteryState: object = { lottery_state: 0 };

export const LotteryState = {
  encode(message: LotteryState, writer: Writer = Writer.create()): Writer {
    if (message.lottery_state !== 0) {
      writer.uint32(8).int32(message.lottery_state);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LotteryState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLotteryState } as LotteryState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lottery_state = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LotteryState {
    const message = { ...baseLotteryState } as LotteryState;
    if (object.lottery_state !== undefined && object.lottery_state !== null) {
      message.lottery_state = lotteryState_LOTTERY_STATEFromJSON(
        object.lottery_state
      );
    } else {
      message.lottery_state = 0;
    }
    return message;
  },

  toJSON(message: LotteryState): unknown {
    const obj: any = {};
    message.lottery_state !== undefined &&
      (obj.lottery_state = lotteryState_LOTTERY_STATEToJSON(
        message.lottery_state
      ));
    return obj;
  },

  fromPartial(object: DeepPartial<LotteryState>): LotteryState {
    const message = { ...baseLotteryState } as LotteryState;
    if (object.lottery_state !== undefined && object.lottery_state !== null) {
      message.lottery_state = object.lottery_state;
    } else {
      message.lottery_state = 0;
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
