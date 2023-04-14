import React from "react";
import ReceiptUpload from "./receiptUpload";

describe("<ReceiptUpload />", () => {
  it("renders", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<ReceiptUpload receipts={[]} setReceipts={() => {}} />);
    cy.get(".description").should("contain.text", "Upload Receipt Images");
  });
});
