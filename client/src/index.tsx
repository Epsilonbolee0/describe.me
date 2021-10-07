import * as React from "react";
import { render } from "react-dom";

import App from "./app/App";

import "styles/index.scss";

render(
  <App />, 
  document.getElementById("root")
);

if (module.hot) {
  module.hot.accept();
}