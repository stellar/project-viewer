import React, {
  useState,
  useEffect,
  useContext,
  useMemo,
  useCallback,
} from "react";
import { Select, Input, Button, Loader } from "@stellar/design-system";
import { getAssetInfo } from "api/getAssetInfo";
import { getCorridorInfo } from "api/getCorridorInfo";
import { getRateInfo } from "api/getRateInfo";
import { getVolumeInfo } from "api/getVolumeInfo";
import { getPeriodOptions } from "helpers/getPeriodOptions";
import { getDateFromDaysAgo, formatDateYYYYMMDD } from "helpers/formatDate";
import { DataContext } from "DataContext";
import { Asset, RequestParams, Aggregate } from "types.d";
import "./styles.scss";

export const Form = ({ baseUrl }: { baseUrl: string }) => {
  const [fromAssetCodeValue, setFromAssetCodeValue] = useState("");
  const [toAssetCodeValue, setToAssetCodeValue] = useState("");
  const [periodValue, setPeriodValue] = useState("");
  const [startDateValue, setStartDateValue] = useState<Date | null>(null);
  const [endDateValue, setEndDateValue] = useState<Date | null>(null);
  const [aggregateValue, setAggregateValue] = useState("");

  const [datePickerStart, setDatePickerStart] = useState("");
  const [datePickerEnd, setDatePickerEnd] = useState("");

  const [assets, setAssets] = useState<Asset[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const { setDataContextValue } = useContext(DataContext);

  const periodOptions = useMemo(() => getPeriodOptions(), []);

  // fullDate in format YYYY-MM-DD (2020-12-30)
  const getDate = (fullDate: string) => {
    const [year, month, day] = fullDate.split("-");
    const date = new Date();

    date.setFullYear(Number(year));
    // Month is 0 based
    date.setMonth(Number(month) - 1);
    date.setDate(Number(day));

    return date;
  };

  const setStartAndEndDateFromPeriod = useCallback(() => {
    let start: Date | null = null;
    let end: Date | null = null;

    switch (periodValue) {
      case "days7":
        start = getDateFromDaysAgo(7);
        end = new Date();
        break;
      case "days30":
        start = getDateFromDaysAgo(30);
        end = new Date();
        break;
      case "2020q1":
        start = getDate("2020-01-01");
        end = getDate("2020-03-31");
        break;
      case "2020q2":
        start = getDate("2020-04-01");
        end = getDate("2020-06-30");
        break;
      case "2020q3":
        start = getDate("2020-07-01");
        end = getDate("2020-09-30");
        break;
      case "2020q4":
        start = getDate("2020-10-01");
        end = getDate("2020-12-31");
        break;
      case "custom":
      default:
      // do nothing
    }

    setStartDateValue(start);
    setEndDateValue(end);
  }, [periodValue]);

  useEffect(() => {
    getAssetInfo(baseUrl).then((response) =>
      setAssets(response?.results || []),
    );
  }, [baseUrl]);

  useEffect(() => {
    setStartAndEndDateFromPeriod();
    setDatePickerStart("");
    setDatePickerEnd("");
  }, [setStartAndEndDateFromPeriod]);

  const findAssetByCode = (code: string) =>
    assets.find((asset) => asset.code === code);

  const handleFetchTradeData = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  ) => {
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

  const handleFetchRates = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  ) => {
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

    const fromAsset = findAssetByCode(fromAssetCodeValue);
    const toAsset = findAssetByCode(toAssetCodeValue);

    if (!fromAsset || !toAsset) {
      // TODO: handle error
      setIsLoading(false);
      return;
    }

    console.log("Fetching rates with params: ", requestParams);

    const response = await getRateInfo({
      baseUrl,
      sourceCode: fromAsset.code,
      sourceIssuer: fromAsset.issuer,
      destCode: toAsset.code,
      destIssuer: toAsset.issuer,
      aggregateBy: aggregateValue,
      start: startDateValue,
      end: endDateValue,
    });

    console.log("Rates RESPONSE: ", response);
    setDataContextValue({ requestParams, response, isRates: true });
    setIsLoading(false);
  };

  const renderAssetOptions = () =>
    assets.map((asset) => (
      <option key={`${asset.code}-${asset.alias}`} value={asset.code}>
        {asset.code}:{asset.alias}
      </option>
    ));

  return (
    <form className="Form">
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
              value={datePickerStart}
              onChange={(e) => {
                setDatePickerStart(e.currentTarget.value);
                setStartDateValue(getDate(e.currentTarget.value));
              }}
              type="date"
              max={formatDateYYYYMMDD()}
            />
            <Input
              id="end-date"
              label="End date"
              value={datePickerEnd}
              onChange={(e) => {
                setDatePickerEnd(e.currentTarget.value);
                setEndDateValue(getDate(e.currentTarget.value));
              }}
              type="date"
              min={datePickerStart}
              max={formatDateYYYYMMDD()}
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
        <Button
          onClick={handleFetchTradeData}
          disabled={isLoading || !(fromAssetCodeValue || toAssetCodeValue)}
        >
          Fetch trade data
        </Button>
        <Button
          onClick={handleFetchRates}
          disabled={isLoading || !(fromAssetCodeValue && toAssetCodeValue)}
        >
          Fetch rates
        </Button>
        {isLoading && (
          <div className="FormLoader">
            <Loader />
          </div>
        )}
      </div>
    </form>
  );
};
