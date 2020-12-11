type VolumeInfoProps = {
  baseUrl: string;
  code: string;
  issuer: string;
  isVolumeFrom: boolean;
};

export const getVolumeInfo = async ({
  baseUrl,
  code,
  issuer,
  isVolumeFrom,
}: VolumeInfoProps) => {
  const volumeURL = `${baseUrl}/volume?code=${code}&issuer=${issuer}&isVolumeFrom=${isVolumeFrom}`;
  const response = await fetch(volumeURL);
  return response.json();
};
