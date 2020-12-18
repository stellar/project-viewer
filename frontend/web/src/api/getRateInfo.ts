import { getEpochTimeFromDate } from "helpers/formatDate";
import { Aggregate } from "types.d";

type RateInfoProps = {
  baseUrl: string;
  sourceCode: string;
  sourceIssuer: string;
  destCode: string;
  destIssuer: string;
  start?: Date | null;
  end?: Date | null;
  aggregateBy?: Aggregate | string;
};

export const getRateInfo = async ({
  baseUrl,
  sourceCode,
  sourceIssuer,
  destCode,
  destIssuer,
  start,
  end,
  aggregateBy,
}: RateInfoProps) => {
  const corridorURL = `${baseUrl}/rate`;
  const params = {
    sourceCode,
    sourceIssuer,
    destCode,
    destIssuer,
    ...(aggregateBy ? { aggregateBy } : {}),
    ...(start ? { start: getEpochTimeFromDate(start) } : {}),
    ...(end ? { end: getEpochTimeFromDate(end) } : {}),
  };
  const response = await fetch(`${corridorURL}?${new URLSearchParams(params)}`);
  return response.json();
};
