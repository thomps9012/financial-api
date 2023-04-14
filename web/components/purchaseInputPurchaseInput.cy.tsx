import React from "react";
import PurchaseInput from "./purchaseInput";
const blank_purchase = {
  amount: 0.0,
  description: "",
  grant_line_item: "",
};
const test_purchase = {
  amount: 10.0,
  description: "test",
  grant_line_item: "line item 1",
};
const update_purchase = {
  amount: 50.0,
  description: "test update",
  grant_line_item: "line item 2",
};
describe("<PurchaseInput /> Blank Input", () => {
  it("renders a blank purchase", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<PurchaseInput purchase={blank_purchase} />);
    cy.get("[name=description]").should("have.value", "");
    cy.get("[name=amount]").should("have.value", 0);
    cy.get("[name=grant_line_item]").should("have.value", "");
  });
  it("renders a blank purchase and updates accordingly", () => {
    cy.mount(<PurchaseInput purchase={blank_purchase} />);
    cy.get("[name=description]").should("have.value", "");
    cy.get("[name=amount]").should("have.value", 0);
    cy.get("[name=grant_line_item]").should("have.value", "");
    cy.get("[name=description]")
      .clear()
      .type(update_purchase.description)
      .should("have.value", update_purchase.description);
    cy.get("[name=amount]")
      .clear()
      .type(JSON.stringify(update_purchase.amount))
      .should("have.value", JSON.stringify(update_purchase.amount));
    cy.get("[name=grant_line_item]")
      .clear()
      .type(update_purchase.grant_line_item)
      .should("have.value", update_purchase.grant_line_item);
  });
});

describe("<PurchaseInput /> Prefilled Input", () => {
  it("renders a prefilled purchase", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<PurchaseInput purchase={test_purchase} />);
    cy.get("[name=description]").should(
      "have.value",
      test_purchase.description
    );
    cy.get("[name=amount]").should("have.value", test_purchase.amount);
    cy.get("[name=grant_line_item]").should(
      "have.value",
      test_purchase.grant_line_item
    );
  });
  it("renders a prefilled purchase and updates accordingly", () => {
    cy.mount(<PurchaseInput purchase={test_purchase} />);
    cy.get("[name=description]").should(
      "have.value",
      test_purchase.description
    );
    cy.get("[name=amount]").should("have.value", test_purchase.amount);
    cy.get("[name=grant_line_item]").should(
      "have.value",
      test_purchase.grant_line_item
    );
    cy.get("[name=description]")
      .clear()
      .type(update_purchase.description)
      .should("have.value", update_purchase.description);
    cy.get("[name=amount]")
      .clear()
      .type(JSON.stringify(update_purchase.amount))
      .should("have.value", JSON.stringify(update_purchase.amount));
    cy.get("[name=grant_line_item]")
      .clear()
      .type(update_purchase.grant_line_item)
      .should("have.value", update_purchase.grant_line_item);
  });
});
