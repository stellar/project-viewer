type CorridorInfoProps = {
  baseUrl: string;
  fromCode: string;
  fromIssuer: string;
  toCode: string;
  toIssuer: string;
};

export const getCorridorInfo = async ({
  baseUrl,
  fromCode,
  fromIssuer,
  toCode,
  toIssuer,
}: CorridorInfoProps) => {
  const corridorURL = `${baseUrl}/corridor?sourceCode=${fromCode}&sourceIssuer=${fromIssuer}&destCode=${toCode}&destIssuer=${toIssuer}`;
  const response = await fetch(corridorURL);
  return response.json();
};
