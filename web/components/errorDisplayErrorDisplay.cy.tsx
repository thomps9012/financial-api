import MockRouter from "@/cypress/mockRouter";
import ErrorDisplay from "./errorDisplay";

describe("<ErrorDisplay />", () => {
  it("renders different error messages", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(
      <MockRouter>
        <ErrorDisplay
          error={new Error("test error")}
          path={"/test_error"}
          message="test error"
        />
      </MockRouter>
    );
    cy.get("#error_message").should("contain.text", "test error");
    cy.mount(
      <MockRouter>
        <ErrorDisplay
          error={new Error("second test error")}
          path={"/second_error"}
          message="test two"
        />
      </MockRouter>
    );
    cy.get("#error_message").should("contain.text", "test two");
  });
});
