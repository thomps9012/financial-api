import React from "react";
import MileageForm from "./mileage_form";

describe("<MileageForm />", () => {
  it("renders a blank form for a new request", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<MileageForm new_request={true} />);
  });
  it("fills out a complete form to match test data specified", () => {
    cy.mount(<MileageForm new_request={true} />);
    const test_data = {
      grant_id: "H79TI082369",
      date: "2019-02-13T07:20:50.52Z",
      category: "FINANCE",
      destination: "b",
      starting_location: "a",
      trip_purpose: "test",
      start_odometer: 1,
      end_odometer: 10,
      tolls: 0.5,
      parking: 0.0,
    };
    // grant select
    cy.get("#mileage-form > :nth-child(2)").select(test_data.grant_id);
    // category select
    cy.get("#mileage-form > :nth-child(4)").select(test_data.category);
    // date select
    cy.get('[type="datetime-local"]').type(test_data.date.split(".")[0]);
    // location select
    cy.get("#start").type(test_data.starting_location);
    // destination select
    cy.get("#end").type(test_data.destination);
    // purpose select
    cy.get("textarea").type(test_data.trip_purpose);
    // start odometer select
    cy.get('[name="start_odometer"]').type(
      JSON.stringify(test_data.start_odometer)
    );
    // end odometer select
    cy.get('[name="end_odometer"]').type(
      JSON.stringify(test_data.end_odometer)
    );
    // tolls select
    cy.get('[name="tolls"]').type(JSON.stringify(test_data.tolls));
    // parking select
    cy.get('[name="parking"]').type(JSON.stringify(test_data.parking));
  });
});
