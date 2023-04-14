describe("navigation to and mileage request creation", () => {
  beforeEach(() => {
    cy.visit("http://localhost:3000");
    cy.get(".archive-btn").click();
  });
  it("achieves this using the base level direct navigation", () => {
    cy.get('[href="/mileage/create"] > h3').click();
    cy.get("h1").first().should("have.text", "New Mileage Request");
  });
  it("achieves this using the navbar navigation", () => {
    cy.get(
      ".header_navIcons__rN6Ws > :nth-child(3) > a > .header_navIcon__I9SuT"
    ).click();
    cy.get("#new").click();
    cy.get("h1").first().should("have.text", "New Mileage Request");
  });
});

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
describe("fills out the mileage form", () => {
  before(() => {
    cy.visit("http://localhost:3000");
    cy.get(".archive-btn").click();
    cy.get('[href="/mileage/create"] > h3').click();
    cy.get("h1").first().should("have.text", "New Mileage Request");
  });
  it("creates a valid form", () => {
    // grant select
    cy.get("#mileage-form > :nth-child(2)").select(test_data.grant_id);
    cy.get("#mileage-form > :nth-child(2)").should(
      "have.value",
      test_data.grant_id
    );
    // category select
    cy.get("#mileage-form > :nth-child(4)").select(test_data.category);
    cy.get("#mileage-form > :nth-child(4)").should(
      "have.value",
      test_data.category
    );
    // date select
    cy.get('[type="datetime-local"]')
      .clear()
      .type(test_data.date.split(".")[0]);
    cy.get('[type="datetime-local"]').should(
      "have.value",
      test_data.date.split(".")[0]
    );
    // location select
    cy.get("#start").clear().type(test_data.starting_location);
    cy.get("#start").should("have.value", test_data.starting_location);
    // destination select
    cy.get("#end").clear().type(test_data.destination);
    cy.get("#end").should("have.value", test_data.destination);
    // purpose select
    cy.get("textarea").clear().type(test_data.trip_purpose);
    cy.get("textarea").should("have.value", test_data.trip_purpose);
    // start odometer select
    cy.get('[name="start_odometer"]')
      .clear()
      .type(JSON.stringify(test_data.start_odometer));
    cy.get('[name="start_odometer"]').should(
      "have.value",
      JSON.stringify(test_data.start_odometer)
    );
    // end odometer select
    cy.get('[name="end_odometer"]')
      .clear()
      .type(JSON.stringify(test_data.end_odometer));
    cy.get('[name="end_odometer"]').should(
      "have.value",
      JSON.stringify(test_data.end_odometer)
    );
    // tolls select
    cy.get('[name="tolls"]').clear().type(JSON.stringify(test_data.tolls));
    cy.get('[name="tolls"]').should(
      "have.value",
      JSON.stringify(test_data.tolls)
    );
    // parking select
    cy.get('[name="parking"]').clear().type(JSON.stringify(test_data.parking));
    cy.get('[name="parking"]').should(
      "have.value",
      JSON.stringify(test_data.parking)
    );
  });
});

