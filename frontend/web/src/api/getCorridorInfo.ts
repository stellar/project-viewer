import { Aggregate } from "types.d";

type CorridorInfoProps = {
  baseUrl: string;
  sourceCode: string;
  sourceIssuer: string;
  destCode: string;
  destIssuer: string;
  start?: string;
  end?: string;
  aggregateBy?: Aggregate | string;
};

export const getCorridorInfo = async ({
  baseUrl,
  sourceCode,
  sourceIssuer,
  destCode,
  destIssuer,
  start,
  end,
  aggregateBy,
}: CorridorInfoProps) => {
  const corridorURL = `${baseUrl}/corridor`;
  const params = {
    sourceCode,
    sourceIssuer,
    destCode,
    destIssuer,
    ...(aggregateBy ? { aggregateBy } : {}),
    ...(start ? { start } : {}),
    ...(end ? { end } : {}),
  };
  const response = await fetch(`${corridorURL}?${new URLSearchParams(params)}`);
  return response.json();
};
