import React from "react";
import Check_Request_Form from "./check_request_form";

describe("NEW <Check_Request_Form />", () => {
  it("renders a new blank request", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Check_Request_Form new_request={true} />);
    cy.get("[name=description]").should("have.value", "");
    cy.get("[name=date]").should("have.value", "");
    cy.get("[name=credit_card]").should("have.value", null);
  });
  it("allows for filling of form fields", () => {
    cy.mount(<Check_Request_Form new_request={true} />);
    cy.get("textarea").should("have.value", "").type("test description").should("have.value", "test description");
    cy.get("[name=date]").type("2019-02-13").should("have.value", "2019-02-13");
    cy.get("[name=credit_card]").select("1234").should("have.value", "1234");
  })
  // OTHER COMPONENTS ARE IMPLICITLY TESTED VIA RELATED COMPONENT TESTS
  it("should add a second purchase field on button click", () => {
    cy.mount(<Check_Request_Form new_request={true} />);
    cy.get("#add_purchase").click();
    cy.get(".purchase-row").should('have.length', 2)
  });
  it("should not be able to add more than five purchase fields", () => {
    cy.mount(<Check_Request_Form new_request={true} />);
    cy.get("#add_purchase")
    .click()
    .click()
    .click()
    .click()
    .click()
    .click()
    .click();
    cy.get(".purchase-row").should('have.length', 5)
  });
  it("should remove a purchase field on the remove purchase button click", () => {
    cy.mount(<Check_Request_Form new_request={true} />);
    cy.get("#add_purchase").click().click();
    cy.get(".purchase-row").should('have.length', 3)
    cy.get("#remove_purchase").click();
    cy.get(".purchase-row").should('have.length', 2)
    cy.get("#remove_purchase").click();
    cy.get(".purchase-row").should('have.length', 1)
    cy.get("#add_purchase").click();
    cy.get("#remove_purchase").click();
    cy.get(".purchase-row").should('have.length', 1)
  });
  it("should not be able to remove the initial purchase field", () => {
    cy.mount(<Check_Request_Form new_request={true} />);
    cy.get("#add_purchase").click().click();
    cy.get(".purchase-row").should('have.length', 3)
    cy.get("#remove_purchase").click();
    cy.get("#remove_purchase").click();
    cy.get("#remove_purchase").click().click().click();
    cy.get(".purchase-row").should('have.length', 1)
  });
});
