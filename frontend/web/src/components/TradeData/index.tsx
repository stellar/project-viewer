import React from "react";
import { Heading2 } from "@stellar/design-system";
import "./styles.scss";

export const TradeData = () => (
  // TODO: use real data
  <div className="TradeData">
    <div className="TradeDataHeader">
      <Heading2>Trade data</Heading2>
      <p>
        From <strong>[asset.code]</strong> to <strong>[asset.code]</strong>
      </p>
    </div>

    <table className="TradeDataTable">
      <thead>
        <tr>
          <th>Date</th>
          <th># of trades</th>
          <th>Trade volume</th>
          <th>Payment volume</th>
          <th>Change</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>Period total</td>
          <td>4,457,549.98</td>
          <td>$2,879.58</td>
          <td>$2,879.58</td>
          <td> - </td>
        </tr>
        <tr>
          <td>Period total</td>
          <td>4,457,549.98</td>
          <td>$2,879.58</td>
          <td>$2,879.58</td>
          <td>+10%</td>
        </tr>
        <tr>
          <td>Period total</td>
          <td>4,457,549.98</td>
          <td>$2,879.58</td>
          <td>$2,879.58</td>
          <td>+1%</td>
        </tr>
      </tbody>
    </table>
  </div>
);
