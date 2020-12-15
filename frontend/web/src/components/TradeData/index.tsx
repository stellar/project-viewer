import React, { useContext } from "react";
import { Heading2 } from "@stellar/design-system";
import { getPeriodOptions } from "helpers/getPeriodOptions";
import { DataContext } from "DataContext";
import { Asset } from "types.d";
import "./styles.scss";

export const TradeData = () => {
  const { dataContextValue } = useContext(DataContext);

  console.log("dataContextValue: ", dataContextValue);

  if (!dataContextValue) {
    return null;
  }

  const { requestParams, response } = dataContextValue;
  const periodItem = requestParams.period
    ? getPeriodOptions().find((p) => p.key === requestParams.period)
    : null;

  const getAssetLabel = (asset: Asset) =>
    asset ? `${asset.code} - ${asset.alias}` : "Any";

  const tableLabels = {
    title: requestParams.aggregate ? "Date" : "Ledger",
    tradeCount: "# of trades",
    tradeVolume: "Trade volume",
    paymentCount: "# of payments",
    paymentVolume: "Payment volume",
  };

  return (
    <div className="TradeData">
      <div className="TradeDataHeader">
        <Heading2>Trade data</Heading2>

        <div className="SearchInfo">
          <div className="SearchInfoAssets">
            From <strong>{getAssetLabel(requestParams.fromAsset)}</strong> to{" "}
            <strong>{getAssetLabel(requestParams.toAsset)}</strong>
          </div>
          {/* TODO: add icon */}
          {periodItem && <div>{periodItem.label}</div>}
        </div>
      </div>

      {response?.results?.length > 0 && (
        <table className="TradeDataTable" data-media={950}>
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
            {/* TODO: update type */}
            {response.results.map((item: any) => (
              <tr key={item.title}>
                <td data-label={tableLabels.title}>{item.title}</td>
                <td data-label={tableLabels.tradeCount}>{item.tradeCount}</td>
                <td data-label={tableLabels.tradeVolume}>{item.tradeVolume}</td>
                <td data-label={tableLabels.paymentCount}>
                  {item.paymentCount}
                </td>
                <td data-label={tableLabels.paymentVolume}>
                  {item.paymentVolume}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};
