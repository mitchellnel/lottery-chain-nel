// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgChangeEntranceFee } from "./types/lottery/tx";
import { MsgSetupLottery } from "./types/lottery/tx";
import { MsgClaimOwner } from "./types/lottery/tx";
import { MsgChangeOwner } from "./types/lottery/tx";
import { MsgEnterLottery } from "./types/lottery/tx";


const types = [
  ["/lotterychainnel.lottery.MsgChangeEntranceFee", MsgChangeEntranceFee],
  ["/lotterychainnel.lottery.MsgSetupLottery", MsgSetupLottery],
  ["/lotterychainnel.lottery.MsgClaimOwner", MsgClaimOwner],
  ["/lotterychainnel.lottery.MsgChangeOwner", MsgChangeOwner],
  ["/lotterychainnel.lottery.MsgEnterLottery", MsgEnterLottery],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgChangeEntranceFee: (data: MsgChangeEntranceFee): EncodeObject => ({ typeUrl: "/lotterychainnel.lottery.MsgChangeEntranceFee", value: MsgChangeEntranceFee.fromPartial( data ) }),
    msgSetupLottery: (data: MsgSetupLottery): EncodeObject => ({ typeUrl: "/lotterychainnel.lottery.MsgSetupLottery", value: MsgSetupLottery.fromPartial( data ) }),
    msgClaimOwner: (data: MsgClaimOwner): EncodeObject => ({ typeUrl: "/lotterychainnel.lottery.MsgClaimOwner", value: MsgClaimOwner.fromPartial( data ) }),
    msgChangeOwner: (data: MsgChangeOwner): EncodeObject => ({ typeUrl: "/lotterychainnel.lottery.MsgChangeOwner", value: MsgChangeOwner.fromPartial( data ) }),
    msgEnterLottery: (data: MsgEnterLottery): EncodeObject => ({ typeUrl: "/lotterychainnel.lottery.MsgEnterLottery", value: MsgEnterLottery.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
