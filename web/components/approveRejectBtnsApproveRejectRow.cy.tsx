import React from "react";
import ApproveRejectRow from "./approveRejectBtns";

describe("<ApproveRejectRow />", () => {
  it("renders nothing when an employee visits", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(
      <ApproveRejectRow
        execReview={false}
        setExecReview={() => {}}
        approveRequest={() => {}}
        rejectRequest={() => {}}
      />
    );
    cy.get("section").should("be.empty");
  });
});
