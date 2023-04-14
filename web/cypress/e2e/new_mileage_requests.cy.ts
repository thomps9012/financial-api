const request_data = {
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
describe("creates a valid mileage request", () => {
  it("submits a valid form", () => {
    cy.visit("http://localhost:3000");
    cy.get(".archive-btn").click();
    cy.get('[href="/mileage/create"] > h3').click();
    // grant select
    cy.get("#mileage-form > :nth-child(2)").select(request_data.grant_id);
    // category select
    cy.get("#mileage-form > :nth-child(4)").select(request_data.category);
    // date select
    cy.get('[type="datetime-local"]')
      .clear()
      .type(request_data.date.split(".")[0]);
    // location select
    cy.get("#start").clear().type(request_data.starting_location);
    // destination select
    cy.get("#end").clear().type(request_data.destination);
    // purpose select
    cy.get("textarea").clear().type(request_data.trip_purpose);
    // start odometer select
    cy.get('[name="start_odometer"]')
      .clear()
      .type(JSON.stringify(request_data.start_odometer));
    // end odometer select
    cy.get('[name="end_odometer"]')
      .clear()
      .type(JSON.stringify(request_data.end_odometer));
    // tolls select
    cy.get('[name="tolls"]').clear().type(JSON.stringify(request_data.tolls));
    // parking select
    cy.get('[name="parking"]')
      .clear()
      .type(JSON.stringify(request_data.parking));
    cy.intercept("http://localhost:3000/api/mileage").as("new_mileage");
    cy.get(".archive-btn").click();
    cy.wait("@new_mileage").then(({ response, request }) => {
      assert.isNotNull(request.body, "Successful request creation");
      assert.isNotNull(response?.body, "API response");
      expect(response?.statusCode).eq(201);
    });
  });
});
describe("creates an valid mileage request", () => {
  it("submits an invalid form", () => {
    cy.visit("http://localhost:3000");
    cy.get(".archive-btn").click();
    cy.get('[href="/mileage/create"] > h3').click();
    // grant select
    cy.get("#mileage-form > :nth-child(2)").select(request_data.grant_id);
    // category select
    cy.get("#mileage-form > :nth-child(4)").select(request_data.category);
    // date select
    cy.get('[type="datetime-local"]')
      .clear()
      .type(request_data.date.split(".")[0]);
    // location select
    cy.get("#start").clear().type(request_data.starting_location);
    // destination select
    cy.get("#end").clear().type(request_data.destination);
    // purpose select
    // start odometer select
    cy.get('[name="start_odometer"]')
      .clear()
      .type(JSON.stringify(request_data.start_odometer));
    // end odometer select
    cy.get('[name="end_odometer"]')
      .clear()
      .type(JSON.stringify(request_data.end_odometer));
    // tolls select
    cy.get('[name="tolls"]').clear().type(JSON.stringify(request_data.tolls));
    // parking select
    cy.get('[name="parking"]')
      .clear()
      .type(JSON.stringify(request_data.parking));
    cy.intercept("http://localhost:3000/api/mileage").as("new_mileage");
    cy.get(".archive-btn").click();
    cy.wait("@new_mileage").then(({ response, request }) => {
      assert.isNotNull(request.body, "Successful request creation");
      assert.isNotNull(response?.body, "API response");
      expect(response?.statusCode).eq(400);
    });
  });
});
