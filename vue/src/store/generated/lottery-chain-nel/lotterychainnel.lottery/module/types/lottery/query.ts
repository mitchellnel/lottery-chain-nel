/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../lottery/params";
import { Owner } from "../lottery/owner";
import { EntranceFee } from "../lottery/entrance_fee";
import { LotteryState } from "../lottery/lottery_state";
import { Player } from "../lottery/player";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "lotterychainnel.lottery";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetOwnerRequest {}

export interface QueryGetOwnerResponse {
  Owner: Owner | undefined;
}

export interface QueryGetEntranceFeeRequest {}

export interface QueryGetEntranceFeeResponse {
  EntranceFee: EntranceFee | undefined;
}

export interface QueryGetLotteryStateRequest {}

export interface QueryGetLotteryStateResponse {
  LotteryState: LotteryState | undefined;
}

export interface QueryGetPlayerRequest {
  address: string;
}

export interface QueryGetPlayerResponse {
  player: Player | undefined;
}

export interface QueryAllPlayerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPlayerResponse {
  player: Player[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetOwnerRequest: object = {};

export const QueryGetOwnerRequest = {
  encode(_: QueryGetOwnerRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOwnerRequest } as QueryGetOwnerRequest;
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

  fromJSON(_: any): QueryGetOwnerRequest {
    const message = { ...baseQueryGetOwnerRequest } as QueryGetOwnerRequest;
    return message;
  },

  toJSON(_: QueryGetOwnerRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryGetOwnerRequest>): QueryGetOwnerRequest {
    const message = { ...baseQueryGetOwnerRequest } as QueryGetOwnerRequest;
    return message;
  },
};

const baseQueryGetOwnerResponse: object = {};

export const QueryGetOwnerResponse = {
  encode(
    message: QueryGetOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Owner !== undefined) {
      Owner.encode(message.Owner, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOwnerResponse } as QueryGetOwnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Owner = Owner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOwnerResponse {
    const message = { ...baseQueryGetOwnerResponse } as QueryGetOwnerResponse;
    if (object.Owner !== undefined && object.Owner !== null) {
      message.Owner = Owner.fromJSON(object.Owner);
    } else {
      message.Owner = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOwnerResponse): unknown {
    const obj: any = {};
    message.Owner !== undefined &&
      (obj.Owner = message.Owner ? Owner.toJSON(message.Owner) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOwnerResponse>
  ): QueryGetOwnerResponse {
    const message = { ...baseQueryGetOwnerResponse } as QueryGetOwnerResponse;
    if (object.Owner !== undefined && object.Owner !== null) {
      message.Owner = Owner.fromPartial(object.Owner);
    } else {
      message.Owner = undefined;
    }
    return message;
  },
};

const baseQueryGetEntranceFeeRequest: object = {};

export const QueryGetEntranceFeeRequest = {
  encode(
    _: QueryGetEntranceFeeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetEntranceFeeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEntranceFeeRequest,
    } as QueryGetEntranceFeeRequest;
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

  fromJSON(_: any): QueryGetEntranceFeeRequest {
    const message = {
      ...baseQueryGetEntranceFeeRequest,
    } as QueryGetEntranceFeeRequest;
    return message;
  },

  toJSON(_: QueryGetEntranceFeeRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetEntranceFeeRequest>
  ): QueryGetEntranceFeeRequest {
    const message = {
      ...baseQueryGetEntranceFeeRequest,
    } as QueryGetEntranceFeeRequest;
    return message;
  },
};

const baseQueryGetEntranceFeeResponse: object = {};

export const QueryGetEntranceFeeResponse = {
  encode(
    message: QueryGetEntranceFeeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.EntranceFee !== undefined) {
      EntranceFee.encode(
        message.EntranceFee,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetEntranceFeeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEntranceFeeResponse,
    } as QueryGetEntranceFeeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EntranceFee = EntranceFee.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEntranceFeeResponse {
    const message = {
      ...baseQueryGetEntranceFeeResponse,
    } as QueryGetEntranceFeeResponse;
    if (object.EntranceFee !== undefined && object.EntranceFee !== null) {
      message.EntranceFee = EntranceFee.fromJSON(object.EntranceFee);
    } else {
      message.EntranceFee = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetEntranceFeeResponse): unknown {
    const obj: any = {};
    message.EntranceFee !== undefined &&
      (obj.EntranceFee = message.EntranceFee
        ? EntranceFee.toJSON(message.EntranceFee)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEntranceFeeResponse>
  ): QueryGetEntranceFeeResponse {
    const message = {
      ...baseQueryGetEntranceFeeResponse,
    } as QueryGetEntranceFeeResponse;
    if (object.EntranceFee !== undefined && object.EntranceFee !== null) {
      message.EntranceFee = EntranceFee.fromPartial(object.EntranceFee);
    } else {
      message.EntranceFee = undefined;
    }
    return message;
  },
};

const baseQueryGetLotteryStateRequest: object = {};

export const QueryGetLotteryStateRequest = {
  encode(
    _: QueryGetLotteryStateRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLotteryStateRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLotteryStateRequest,
    } as QueryGetLotteryStateRequest;
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

  fromJSON(_: any): QueryGetLotteryStateRequest {
    const message = {
      ...baseQueryGetLotteryStateRequest,
    } as QueryGetLotteryStateRequest;
    return message;
  },

  toJSON(_: QueryGetLotteryStateRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetLotteryStateRequest>
  ): QueryGetLotteryStateRequest {
    const message = {
      ...baseQueryGetLotteryStateRequest,
    } as QueryGetLotteryStateRequest;
    return message;
  },
};

const baseQueryGetLotteryStateResponse: object = {};

export const QueryGetLotteryStateResponse = {
  encode(
    message: QueryGetLotteryStateResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.LotteryState !== undefined) {
      LotteryState.encode(
        message.LotteryState,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetLotteryStateResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetLotteryStateResponse,
    } as QueryGetLotteryStateResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.LotteryState = LotteryState.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLotteryStateResponse {
    const message = {
      ...baseQueryGetLotteryStateResponse,
    } as QueryGetLotteryStateResponse;
    if (object.LotteryState !== undefined && object.LotteryState !== null) {
      message.LotteryState = LotteryState.fromJSON(object.LotteryState);
    } else {
      message.LotteryState = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetLotteryStateResponse): unknown {
    const obj: any = {};
    message.LotteryState !== undefined &&
      (obj.LotteryState = message.LotteryState
        ? LotteryState.toJSON(message.LotteryState)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetLotteryStateResponse>
  ): QueryGetLotteryStateResponse {
    const message = {
      ...baseQueryGetLotteryStateResponse,
    } as QueryGetLotteryStateResponse;
    if (object.LotteryState !== undefined && object.LotteryState !== null) {
      message.LotteryState = LotteryState.fromPartial(object.LotteryState);
    } else {
      message.LotteryState = undefined;
    }
    return message;
  },
};

const baseQueryGetPlayerRequest: object = { address: "" };

export const QueryGetPlayerRequest = {
  encode(
    message: QueryGetPlayerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPlayerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPlayerRequest } as QueryGetPlayerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPlayerRequest {
    const message = { ...baseQueryGetPlayerRequest } as QueryGetPlayerRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetPlayerRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPlayerRequest>
  ): QueryGetPlayerRequest {
    const message = { ...baseQueryGetPlayerRequest } as QueryGetPlayerRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetPlayerResponse: object = {};

export const QueryGetPlayerResponse = {
  encode(
    message: QueryGetPlayerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.player !== undefined) {
      Player.encode(message.player, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPlayerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPlayerResponse } as QueryGetPlayerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player = Player.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPlayerResponse {
    const message = { ...baseQueryGetPlayerResponse } as QueryGetPlayerResponse;
    if (object.player !== undefined && object.player !== null) {
      message.player = Player.fromJSON(object.player);
    } else {
      message.player = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPlayerResponse): unknown {
    const obj: any = {};
    message.player !== undefined &&
      (obj.player = message.player ? Player.toJSON(message.player) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPlayerResponse>
  ): QueryGetPlayerResponse {
    const message = { ...baseQueryGetPlayerResponse } as QueryGetPlayerResponse;
    if (object.player !== undefined && object.player !== null) {
      message.player = Player.fromPartial(object.player);
    } else {
      message.player = undefined;
    }
    return message;
  },
};

const baseQueryAllPlayerRequest: object = {};

export const QueryAllPlayerRequest = {
  encode(
    message: QueryAllPlayerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPlayerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPlayerRequest } as QueryAllPlayerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPlayerRequest {
    const message = { ...baseQueryAllPlayerRequest } as QueryAllPlayerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPlayerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPlayerRequest>
  ): QueryAllPlayerRequest {
    const message = { ...baseQueryAllPlayerRequest } as QueryAllPlayerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPlayerResponse: object = {};

export const QueryAllPlayerResponse = {
  encode(
    message: QueryAllPlayerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.player) {
      Player.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPlayerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPlayerResponse } as QueryAllPlayerResponse;
    message.player = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.player.push(Player.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPlayerResponse {
    const message = { ...baseQueryAllPlayerResponse } as QueryAllPlayerResponse;
    message.player = [];
    if (object.player !== undefined && object.player !== null) {
      for (const e of object.player) {
        message.player.push(Player.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPlayerResponse): unknown {
    const obj: any = {};
    if (message.player) {
      obj.player = message.player.map((e) =>
        e ? Player.toJSON(e) : undefined
      );
    } else {
      obj.player = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPlayerResponse>
  ): QueryAllPlayerResponse {
    const message = { ...baseQueryAllPlayerResponse } as QueryAllPlayerResponse;
    message.player = [];
    if (object.player !== undefined && object.player !== null) {
      for (const e of object.player) {
        message.player.push(Player.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Owner by index. */
  Owner(request: QueryGetOwnerRequest): Promise<QueryGetOwnerResponse>;
  /** Queries a EntranceFee by index. */
  EntranceFee(
    request: QueryGetEntranceFeeRequest
  ): Promise<QueryGetEntranceFeeResponse>;
  /** Queries a LotteryState by index. */
  LotteryState(
    request: QueryGetLotteryStateRequest
  ): Promise<QueryGetLotteryStateResponse>;
  /** Queries a Player by index. */
  Player(request: QueryGetPlayerRequest): Promise<QueryGetPlayerResponse>;
  /** Queries a list of Player items. */
  PlayerAll(request: QueryAllPlayerRequest): Promise<QueryAllPlayerResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Owner(request: QueryGetOwnerRequest): Promise<QueryGetOwnerResponse> {
    const data = QueryGetOwnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Query",
      "Owner",
      data
    );
    return promise.then((data) =>
      QueryGetOwnerResponse.decode(new Reader(data))
    );
  }

  EntranceFee(
    request: QueryGetEntranceFeeRequest
  ): Promise<QueryGetEntranceFeeResponse> {
    const data = QueryGetEntranceFeeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Query",
      "EntranceFee",
      data
    );
    return promise.then((data) =>
      QueryGetEntranceFeeResponse.decode(new Reader(data))
    );
  }

  LotteryState(
    request: QueryGetLotteryStateRequest
  ): Promise<QueryGetLotteryStateResponse> {
    const data = QueryGetLotteryStateRequest.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Query",
      "LotteryState",
      data
    );
    return promise.then((data) =>
      QueryGetLotteryStateResponse.decode(new Reader(data))
    );
  }

  Player(request: QueryGetPlayerRequest): Promise<QueryGetPlayerResponse> {
    const data = QueryGetPlayerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Query",
      "Player",
      data
    );
    return promise.then((data) =>
      QueryGetPlayerResponse.decode(new Reader(data))
    );
  }

  PlayerAll(request: QueryAllPlayerRequest): Promise<QueryAllPlayerResponse> {
    const data = QueryAllPlayerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "lotterychainnel.lottery.Query",
      "PlayerAll",
      data
    );
    return promise.then((data) =>
      QueryAllPlayerResponse.decode(new Reader(data))
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
