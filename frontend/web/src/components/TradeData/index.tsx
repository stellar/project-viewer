import React, { useContext } from "react";
import { Heading2, IconCalendar } from "@stellar/design-system";
import { getPeriodOptions } from "helpers/getPeriodOptions";
import { formatDateYYYYMMDD } from "helpers/formatDate";
import { DataContext } from "DataContext";
import { Asset, TradeDataResponseItem } from "types.d";
import "./styles.scss";

export const TradeData = () => {
  const { dataContextValue } = useContext(DataContext);

  if (!dataContextValue) {
    return null;
  }

  const { requestParams, response, isRates } = dataContextValue;
  const periodItem = requestParams.period
    ? getPeriodOptions().find((p) => p.key === requestParams.period)
    : null;

  const getAssetLabel = (asset: Asset) =>
    asset ? `${asset.code} - ${asset.alias}` : "Any";

  const getPeriodLabel = () => {
    const { startDate, endDate } = requestParams;

    if (!periodItem || !(startDate && endDate)) {
      return null;
    }

    return periodItem.key === "custom"
      ? `${formatDateYYYYMMDD(startDate, "/")} - ${formatDateYYYYMMDD(
          endDate,
          "/",
        )}`
      : periodItem?.label;
  };

  const tableLabels = {
    title: requestParams.aggregate ? "Date" : "Ledger",
    tradeCount: "# of trades",
    tradeVolume: "Trade volume",
    paymentCount: "# of payments",
    paymentVolume: "Payment volume",
    rate: "Rate",
  };

  const renderTable = () => {
    if (isRates) {
      return (
        <table className="TradeDataTable">
          <thead>
            <tr>
              <th>{tableLabels.title}</th>
              <th>{tableLabels.rate}</th>
            </tr>
          </thead>
          <tbody>
            {response.results.map((item: TradeDataResponseItem) => (
              <tr key={item.title}>
                <td data-label={tableLabels.title}>{item.title}</td>
                <td data-label={tableLabels.rate}>{item.rate}</td>
              </tr>
            ))}
          </tbody>
        </table>
      );
    }

    return (
      <table className="TradeDataTable">
        <thead>
          <tr>
            <th>{tableLabels.title}</th>
            <th>{tableLabels.tradeCount}</th>
            <th>{tableLabels.tradeVolume}</th>
            <th>{tableLabels.paymentCount}</th>
            <th>{tableLabels.paymentVolume}</th>
          </tr>
        </thead>
        <tbody>
          {response.results.map((item: TradeDataResponseItem) => (
            <tr key={item.title}>
              <td data-label={tableLabels.title}>{item.title}</td>
              <td data-label={tableLabels.tradeCount}>{item.tradeCount}</td>
              <td data-label={tableLabels.tradeVolume}>{item.tradeVolume}</td>
              <td data-label={tableLabels.paymentCount}>{item.paymentCount}</td>
              <td data-label={tableLabels.paymentVolume}>
                {item.paymentVolume}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  };

  return (
    <div className="TradeData">
      <div className="TradeDataHeader">
        <Heading2>{isRates ? "Rates" : "Trade data"}</Heading2>

        <div className="SearchInfo">
          <div className="SearchInfoAssets">
            From <strong>{getAssetLabel(requestParams.fromAsset)}</strong> to{" "}
            <strong>{getAssetLabel(requestParams.toAsset)}</strong>
          </div>
          {periodItem && (
            <div className="SearchInfoPeriod">
              <IconCalendar />
              <span>{getPeriodLabel()}</span>
            </div>
          )}
        </div>
      </div>

      {response?.results === null && (
        <p>Nothing was found matching your search criteria.</p>
      )}

      {response?.results?.length > 0 && renderTable()}
    </div>
  );
};
