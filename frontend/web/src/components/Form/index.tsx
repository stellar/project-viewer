import React, { useState, useEffect } from "react";
import { Select, Input, Button } from "@stellar/design-system";
import { getAssetInfo } from "api/getAssetInfo";
import { getCorridorInfo } from "api/getCorridorInfo";
import { getVolumeInfo } from "api/getVolumeInfo";
import "./styles.scss";

type Asset = {
  code: string;
  issuer: string;
  alias: string;
};

export const Form = ({ baseUrl }: { baseUrl: string }) => {
  const [fromAssetCodeValue, setFromAssetCodeValue] = useState("");
  const [toAssetCodeValue, setToAssetCodeValue] = useState("");
  const [periodValue, setPeriodValue] = useState("");
  const [startDateValue, setStartDateValue] = useState("");
  const [endDateValue, setendDateValue] = useState("");
  const [aggregateValue, setAggregateValue] = useState("");

  const [assets, setAssets] = useState<Asset[]>([]);

  useEffect(() => {
    // TODO: handle no results case
    getAssetInfo(baseUrl).then((response) =>
      setAssets(response?.results || []),
    );
  }, [baseUrl]);

  const findAssetByCode = (code: string) =>
    assets.find((asset) => asset.code === code);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    // TODO: pass period and aggregate values
    e.preventDefault();

    const params = {
      fromAssetCodeValue,
      toAssetCodeValue,
      periodValue,
      startDateValue,
      endDateValue,
      aggregateValue,
    };

    console.log("Submitting with params: ", params);

    // Corridor info
    if (fromAssetCodeValue && toAssetCodeValue) {
      console.log(
        `Finding volume between ${fromAssetCodeValue}->${toAssetCodeValue}`,
      );

      const fromAsset = findAssetByCode(fromAssetCodeValue);
      const toAsset = findAssetByCode(toAssetCodeValue);

      if (!fromAsset || !toAsset) {
        // TODO: handle error
        return;
      }

      const response = await getCorridorInfo({
        baseUrl,
        fromCode: fromAsset.code,
        fromIssuer: fromAsset.issuer,
        toCode: toAsset.code,
        toIssuer: toAsset.issuer,
      });

      console.log("Corridor RESPONSE: ", response);

      return;
    }

    // Volume info
    if (fromAssetCodeValue || toAssetCodeValue) {
      const asset = findAssetByCode(fromAssetCodeValue || toAssetCodeValue);

      if (!asset) {
        // TODO: handle error
        return;
      }

      console.log(`Finding volume from ${asset.code}`);

      const response = await getVolumeInfo({
        baseUrl,
        code: asset.code,
        issuer: asset.issuer,
        isVolumeFrom: Boolean(fromAssetCodeValue),
      });

      console.log("Volume RESPONSE: ", response);

      return;
    }

    // TODO: handle no results
    console.log("No results");
  };

  const renderAssetOptions = () =>
    assets.map((asset) => (
      <option key={`${asset.code}-${asset.alias}`} value={asset.code}>
        {asset.code}:{asset.alias}
      </option>
    ));

  return (
    <form className="Form" onSubmit={handleSubmit}>
      <div className="FormInputs">
        <Select
          id="fromAsset"
          label="From"
          value={fromAssetCodeValue}
          onChange={(e) => {
            setFromAssetCodeValue(e.currentTarget.value);
          }}
        >
          <option value="">Select…</option>
          {renderAssetOptions()}
        </Select>

        <Select
          id="toAsset"
          label="To"
          value={toAssetCodeValue}
          onChange={(e) => setToAssetCodeValue(e.currentTarget.value)}
        >
          <option value="">Select…</option>
          {renderAssetOptions()}
        </Select>

        <Select
          id="period"
          label="Period"
          value={periodValue}
          onChange={(e) => setPeriodValue(e.currentTarget.value)}
        >
          <option value="">Select…</option>
          <option value="">Last 7 days</option>
          <option value="">Last 30 days</option>
          <option value="">Q1 2020</option>
          <option value="">Q2 2020</option>
          <option value="">Q3 2020</option>
          <option value="custom">Custom</option>
        </Select>

        {periodValue === "custom" && (
          <div className="FormInputsCustom">
            <Input
              id="start-date"
              label="Start date"
              value={startDateValue}
              // TODO: update with date picker
              onChange={(e) => setStartDateValue(e.currentTarget.value)}
            />
            <Input
              id="end-date"
              label="End date"
              value={endDateValue}
              // TODO: update with date picker
              onChange={(e) => setendDateValue(e.currentTarget.value)}
            />
          </div>
        )}

        <Select
          id="aggregate"
          label="Aggregate"
          value={aggregateValue}
          onChange={(e) => setAggregateValue(e.currentTarget.value)}
        >
          <option value="">Select…</option>
          <option value="day">Day</option>
          <option value="week">Week</option>
          <option value="month">Month</option>
          <option value="quarter">Quarter</option>
          <option value="year">Year</option>
        </Select>
      </div>

      <Button>Fetch trade data</Button>
    </form>
  );
};
