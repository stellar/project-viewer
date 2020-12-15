export enum Aggregate {
  day = "day",
  week = "week",
  month = "month",
  quarter = "quarter",
  year = "year",
}

export type Asset = {
  code: string;
  issuer: string;
  alias: string;
};

export type RequestParams = {
  fromAsset?: Asset;
  toAsset?: Asset;
  period?: string;
  startDate?: string;
  endDate?: string;
  // TODO: handle without string
  aggregate?: Aggregate | string;
};

export type TradeDataResponseItem = {
  title: string;
  tradeCount: number;
  tradeVolume: number;
  paymentCount: number;
  paymentVolume: number;
};
