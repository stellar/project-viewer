import React from "react";
import { ProjectLogo, TextLink, TextLinkVariant } from "@stellar/design-system";
import { Form } from "components/Form";
import { TradeData } from "components/TradeData";

import "./App.scss";

const App = () => (
  <>
    <div className="Header">
      <div className="Inset">
        <ProjectLogo title="Project Viewer" />
      </div>
    </div>

    <div className="Content">
      <div className="Section">
        <div className="Inset">
          <Form baseUrl="http://localhost:8080" />
        </div>
      </div>

      <div className="Inset">
        {/* TODO: pass data through context */}
        <TradeData />
      </div>
    </div>

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

export default App;
