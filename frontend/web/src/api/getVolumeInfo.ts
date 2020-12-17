import { getEpochTimeFromDate } from "helpers/formatDate";
import { Aggregate } from "types.d";

type VolumeInfoProps = {
  baseUrl: string;
  code: string;
  issuer: string;
  isVolumeFrom: boolean;
  start?: Date | null;
  end?: Date | null;
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
    ...(start ? { start: getEpochTimeFromDate(start) } : {}),
    ...(end ? { end: getEpochTimeFromDate(end) } : {}),
  };
  const response = await fetch(`${volumeURL}?${new URLSearchParams(params)}`);
  return response.json();
};
