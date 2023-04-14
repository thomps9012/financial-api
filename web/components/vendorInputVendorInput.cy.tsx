import React from "react";
import VendorInput from "./vendorInput";

const blank_state = {
  name: "",
  website: "",
  address_line_one: "",
  address_line_two: "",
};
const test_state = {
  name: "Test Company",
  website: "test.com",
  address_line_one: "123 st, Test City, TA, 12345",
  address_line_two: "Building 1",
};
const update_state = {
  name: "Test Company Edit",
  website: "https://www.test.com",
  address_line_one: "12345 ST, Test Place, TQ, 54321",
  address_line_two: "Floor 1",
};
const setState = () => {};
describe("<VendorInput /> Blank Input", () => {
  it("renders with a blank state", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<VendorInput state={blank_state} setState={setState} />);
    cy.get('[name="name"]').should("have.value", "");
    cy.get('[name="website"]').should("have.value", "");
    cy.get('[name="address_line_one"]').should("have.value", "");
    cy.get('[name="address_line_two"]').should("have.value", "");
  });
  it("allows the user to update the form state", () => {
    cy.mount(<VendorInput state={blank_state} setState={setState} />);
    cy.get('[name="name"]')
      .clear()
      .type(test_state.name)
      .should("have.value", test_state.name);
    cy.get('[name="website"]')
      .clear()
      .type(test_state.website)
      .should("have.value", test_state.website);
    cy.get('[name="address_line_one"]')
      .clear()
      .type(test_state.address_line_one)
      .should("have.value", test_state.address_line_one);
    cy.get('[name="address_line_two"]')
      .clear()
      .type(test_state.address_line_two)
      .should("have.value", test_state.address_line_two);
  });
});
describe("<VendorInput /> Prefilled Input", () => {
  it("renders with a prefilled state", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<VendorInput state={test_state} setState={setState} />);
    cy.get('[name="name"]').should("have.value", test_state.name);
    cy.get('[name="website"]').should("have.value", test_state.website);
    cy.get('[name="address_line_one"]').should(
      "have.value",
      test_state.address_line_one
    );
    cy.get('[name="address_line_two"]').should(
      "have.value",
      test_state.address_line_two
    );
  });
  it("allows the user to update the form state", () => {
    cy.mount(<VendorInput state={test_state} setState={setState} />);
    cy.get('[name="name"]')
      .clear()
      .type(update_state.name)
      .should("have.value", update_state.name);
    cy.get('[name="website"]')
      .clear()
      .type(update_state.website)
      .should("have.value", update_state.website);
    cy.get('[name="address_line_one"]')
      .clear()
      .type(update_state.address_line_one)
      .should("have.value", update_state.address_line_one);
    cy.get('[name="address_line_two"]')
      .clear()
      .type(update_state.address_line_two)
      .should("have.value", update_state.address_line_two);
  });
});
