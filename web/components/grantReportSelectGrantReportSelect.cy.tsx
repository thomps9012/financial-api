import React from "react";
import GrantReportSelect from "./grantReportSelect";

describe("<GrantReportSelect />", () => {
  it("renders the correct title", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<GrantReportSelect reportType="Mileage" setReport={() => {}} />);
    cy.get("h3").should("have.text", "Grant Mileage Report");
    cy.get(".archive-btn").should("have.text", "Generate Report");
    cy.mount(
      <GrantReportSelect reportType="Check Request" setReport={() => {}} />
    );
    cy.get("h3").should("have.text", "Grant Check Request Report");
    cy.mount(
      <GrantReportSelect reportType="Check Petty Cash" setReport={() => {}} />
    );
    cy.get("h3").should("have.text", "Grant Check Petty Cash Report");
  });
});
