import React from "react";
import MonthlyReportSelect from "./monthlyReportSelect";

describe("NEW <MonthlyReportSelect />", () => {
  it("renders the correct title", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<MonthlyReportSelect reportType="Mileage" setReport={() => {}} />);
    cy.get('label').should("contain.text", "Mileage Report");
    cy.mount(<MonthlyReportSelect reportType="Check Request" setReport={() => {}} />);
    cy.get('label').should("contain.text", "Check Request Report");
    cy.mount(<MonthlyReportSelect reportType="Petty Cash" setReport={() => {}} />);
    cy.get('label').should("contain.text", "Petty Cash Report");
  });
});
