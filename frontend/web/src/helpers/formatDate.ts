export const getDateFromDaysAgo = (daysAgo: number) => {
  const date = new Date();
  date.setDate(date.getDate() - daysAgo);
  return date;
};

export const formatDateYYYYMMDD = (date = new Date(), separator = "-") => {
  const [month, day, year] = date.toLocaleDateString("en-US").split("/");
  return [year, month, day].join(separator);
};

export const getEpochTimeFromDate = (date: Date = new Date()) =>
  Math.floor(date.getTime() / 1000).toString();
