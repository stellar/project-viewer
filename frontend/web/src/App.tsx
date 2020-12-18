import React, { useState, useMemo } from "react";
import { ProjectLogo, TextLink, TextLinkVariant } from "@stellar/design-system";
import { Form } from "components/Form";
import { TradeData } from "components/TradeData";
import { DataContext } from "./DataContext";

import "./App.scss";

const App = () => {
  const [dataContextValue, setDataContextValue] = useState(null);
  const tradeDataValue = useMemo(
    () => ({ dataContextValue, setDataContextValue }),
    [dataContextValue, setDataContextValue],
  );

  return (
    <>
      <div className="Header">
        <div className="Inset">
          <ProjectLogo title="Project Viewer" />
        </div>
      </div>

      <DataContext.Provider value={tradeDataValue}>
        <div className="Content">
          <div className="Section">
            <div className="Inset">
              {/* TODO: update before launch */}
              <Form baseUrl="http://localhost:8080" />
            </div>
          </div>

          <div className="Inset">
            <TradeData />
          </div>
        </div>
      </DataContext.Provider>

      <div className="Footer">
        <div className="Inset">
          <TextLink
            href="https://github.com/stellar/project-viewer"
            rel="noreferrer"
            target="_blank"
            variant={TextLinkVariant.secondary}
          >
            GitHub
          </TextLink>
        </div>
      </div>
    </>
  );
};

export default App;
