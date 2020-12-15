import React, {
  useState,
  useEffect,
  useContext,
  useMemo,
  useCallback,
} from "react";
import { Select, Input, Button } from "@stellar/design-system";
import { getAssetInfo } from "api/getAssetInfo";
import { getCorridorInfo } from "api/getCorridorInfo";
import { getVolumeInfo } from "api/getVolumeInfo";
import { getPeriodOptions } from "helpers/getPeriodOptions";
import { DataContext } from "DataContext";
import { Asset, RequestParams, Aggregate } from "types.d";
import "./styles.scss";

export const Form = ({ baseUrl }: { baseUrl: string }) => {
  const [fromAssetCodeValue, setFromAssetCodeValue] = useState("");
  const [toAssetCodeValue, setToAssetCodeValue] = useState("");
  const [periodValue, setPeriodValue] = useState("");
  const [startDateValue, setStartDateValue] = useState("");
  const [endDateValue, setendDateValue] = useState("");
  const [aggregateValue, setAggregateValue] = useState("");

  const [assets, setAssets] = useState<Asset[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const { setDataContextValue } = useContext(DataContext);

  const periodOptions = useMemo(() => getPeriodOptions(), []);

  const getEpochTimeFromDate = (date?: Date) => {
    const dt = date || new Date();
    return Math.floor(dt.getTime() / 1000).toString();
  };

  const getEpochTimeFromDaysAgo = (daysAgo: number) => {
    const startDate = new Date();
    return Math.floor(
      startDate.setDate(startDate.getDate() - daysAgo) / 1000,
    ).toString();
  };

  // fullDate in format YYYY/MM/DD (2020/12/30)
  const getDate = (fullDate: string) => {
    const [year, month, day] = fullDate.split("/");
    const date = new Date();

    date.setFullYear(Number(year));
    date.setMonth(Number(month));
    date.setDate(Number(day));

    return date;
  };

  const setStartAndEndDateFromPeriod = useCallback(() => {
    let start = "";
    let end = "";

    switch (periodValue) {
      case "days7":
        start = getEpochTimeFromDaysAgo(7);
        end = getEpochTimeFromDate();
        break;
      case "days30":
        start = getEpochTimeFromDaysAgo(30);
        end = getEpochTimeFromDate();
        break;
      case "2020q1":
        start = getEpochTimeFromDate(getDate("2020/01/01"));
        end = getEpochTimeFromDate(getDate("2020/03/31"));
        break;
      case "2020q2":
        start = getEpochTimeFromDate(getDate("2020/04/01"));
        end = getEpochTimeFromDate(getDate("2020/06/30"));
        break;
      case "2020q3":
        start = getEpochTimeFromDate(getDate("2020/07/01"));
        end = getEpochTimeFromDate(getDate("2020/09/30"));
        break;
      case "2020q4":
        start = getEpochTimeFromDate(getDate("2020/10/01"));
        end = getEpochTimeFromDate(getDate("2020/12/31"));
        break;
      case "custom":
      default:
      // do nothing
    }

    setStartDateValue(start);
    setendDateValue(end);
  }, [periodValue]);

  useEffect(() => {
    getAssetInfo(baseUrl).then((response) =>
      setAssets(response?.results || []),
    );
  }, [baseUrl]);

  useEffect(() => {
    setStartAndEndDateFromPeriod();
  }, [setStartAndEndDateFromPeriod]);

  const findAssetByCode = (code: string) =>
    assets.find((asset) => asset.code === code);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);

    const requestParams: RequestParams = {
      fromAsset: findAssetByCode(fromAssetCodeValue),
      toAsset: findAssetByCode(toAssetCodeValue),
      period: periodValue,
      startDate: startDateValue,
      endDate: endDateValue,
      aggregate: aggregateValue,
    };

    setDataContextValue(null);

    console.log("Submitting with params: ", requestParams);

    // Corridor info
    if (fromAssetCodeValue && toAssetCodeValue) {
      console.log(
        `Finding volume between ${fromAssetCodeValue}->${toAssetCodeValue}`,
      );

      const fromAsset = findAssetByCode(fromAssetCodeValue);
      const toAsset = findAssetByCode(toAssetCodeValue);

      if (!fromAsset || !toAsset) {
        // TODO: handle error
        setIsLoading(false);
        return;
      }

      const response = await getCorridorInfo({
        baseUrl,
        sourceCode: fromAsset.code,
        sourceIssuer: fromAsset.issuer,
        destCode: toAsset.code,
        destIssuer: toAsset.issuer,
        aggregateBy: aggregateValue,
        start: startDateValue,
        end: endDateValue,
      });

      console.log("Corridor RESPONSE: ", response);
      setDataContextValue({ requestParams, response });
      setIsLoading(false);

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
        aggregateBy: aggregateValue,
        start: startDateValue,
        end: endDateValue,
      });

      console.log("Volume RESPONSE: ", response);
      setDataContextValue({ requestParams, response });
      setIsLoading(false);

      return;
    }

    // TODO: handle no results
    console.log("No results");
    setIsLoading(false);
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
          {periodOptions.map((p) => (
            <option value={p.key} key={p.key}>
              {p.label}
            </option>
          ))}
        </Select>

        {periodValue === "custom" && (
          <div className="FormInputsCustom">
            <Input
              id="start-date"
              label="Start date"
              value={startDateValue}
              // TODO: update with date picker
              onChange={(e) => {
                // TODO: getEpochTimeFromDate
                setStartDateValue(e.currentTarget.value);
              }}
            />
            <Input
              id="end-date"
              label="End date"
              value={endDateValue}
              // TODO: update with date picker
              // TODO: getEpochTimeFromDate
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
          <option value={Aggregate.day}>Day</option>
          <option value={Aggregate.week}>Week</option>
          <option value={Aggregate.month}>Month</option>
          <option value={Aggregate.quarter}>Quarter</option>
          <option value={Aggregate.year}>Year</option>
        </Select>
      </div>

      <div className="FormButtonWrapper">
        <Button disabled={isLoading}>Fetch trade data</Button>
        {/* TODO: add loader */}
        {isLoading && <span className="FormLoader">Fetching…</span>}
      </div>
    </form>
  );
};
