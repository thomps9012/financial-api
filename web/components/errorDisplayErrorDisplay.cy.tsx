import MockRouter from "@/cypress/mockRouter";
import ErrorDisplay from "./errorDisplay";

describe("<ErrorDisplay />", () => {
  it("renders different error messages", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(
      <MockRouter>
        <ErrorDisplay path={"/test_error"} message="test error" />
      </MockRouter>
    );
    cy.get("#error_message").should("have.text", "test error");
    cy.get("#error_path").should("have.text", "Occurred on /test_error");
    cy.mount(
      <MockRouter>
        <ErrorDisplay path={"/second_error"} message="test two" />
      </MockRouter>
    );
    cy.get("#error_message").should("have.text", "test two");
    cy.get("#error_path").should("have.text", "Occurred on /second_error");
  });
});
