describe("initial page load and authentication", () => {
  beforeEach(() => {
    cy.visit("http://localhost:3000");
  });
  it("loads an unauthorized page load screen", () => {
    cy.get("h1").first().should("have.text", "Access Denied");
  });
  it("signs a user in with a test account on button click and out on another click", () => {
    cy.intercept("http://localhost:3000/api/auth/login").as("login");
    cy.intercept("http://localhost:3000/api/auth/logout").as("logout");
    cy.get(".archive-btn").click();
    cy.wait("@login").then(({ response }) => {
      assert.isNotNull(response?.body, "Login API call has been made");
      expect(response?.statusCode).to.be.oneOf([200, 201]);
    });
    cy.get(
      ".header_navIcons__rN6Ws > :nth-child(8) > .header_navIcon__I9SuT"
    ).click();
    cy.wait("@logout").then(({ response }) => {
      assert.isNotNull(response?.body, "Logout API has been called");
      expect(response?.statusCode).eq(200);
    });
  });
  it("signs a user in with a test account on button click displays their name", () => {
    cy.intercept("http://localhost:3000/api/auth/login").as("login");
    cy.get(".archive-btn").click();
    cy.wait("@login").then(({ response }) => {
      assert.isNotNull(response?.body, "Login API call has been made");
      expect(response?.statusCode).to.be.oneOf([200, 201]);
    });
    cy.get(
      '.header_navHeader__vsjcE > [href="/profile"] > p > .header_signedInText__s_n47 > strong'
    ).should("have.text", "TEST FINANCE");
  });
});
