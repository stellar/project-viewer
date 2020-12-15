import { Aggregate } from "types.d";

type VolumeInfoProps = {
  baseUrl: string;
  code: string;
  issuer: string;
  isVolumeFrom: boolean;
  start?: string;
  end?: string;
  aggregateBy?: Aggregate | string;
};

export const getVolumeInfo = async ({
  baseUrl,
  code,
  issuer,
  isVolumeFrom,
  start,
  end,
  aggregateBy,
}: VolumeInfoProps) => {
  const volumeURL = `${baseUrl}/volume`;
  const params = {
    code,
    issuer,
    volumeFrom: isVolumeFrom.toString(),
    ...(aggregateBy ? { aggregateBy } : {}),
    ...(start ? { start } : {}),
    ...(end ? { end } : {}),
  };
  const response = await fetch(`${volumeURL}?${new URLSearchParams(params)}`);
  return response.json();
};
