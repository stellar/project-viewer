export const getAssetInfo = async (baseUrl: string) => {
  const assetURL = `${baseUrl}/assets`;
  const response = await fetch(assetURL);
  return response.json();
};
