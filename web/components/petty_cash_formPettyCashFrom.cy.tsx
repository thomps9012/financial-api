import React from "react";
import PettyCashFrom from "./petty_cash_form";

describe("NEW <PettyCashFrom />", () => {
  it("renders a blank form", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<PettyCashFrom new_request={true} />);
    cy.get("[name=amount]").should("have.value", 0)
    cy.get("[name=date]").should("have.value", "")
    cy.get("[name=description]").should("have.value", "")
  });
  it("allows for non tested fields to be filled", () => {
    cy.mount(<PettyCashFrom new_request={true} />);
    cy.get("[name=amount]").should("have.value", 0.0).clear().type("15.5").should("have.value", "15.5")
    cy.get("[name=date]").should("have.value", "").clear().type("2020-01-02").should("have.value", "2020-01-02")
    cy.get("textarea").should("have.value", "").clear().type("test description").should("have.value", "test description")
  })
});
// OTHER FIELDS ARE TESTED IN RELEVANT COMPONENT TESTS