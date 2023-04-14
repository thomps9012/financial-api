import React from "react";
import AccessDenied from "./accessDenied";
describe("<AccessDenied />", () => {
  it("renders", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<AccessDenied />);
    cy.get("h1").first().should("have.text", "Access Denied");
    cy.get("a").should("have.text", "Sign In");
  });
 
});
